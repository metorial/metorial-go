package metorial

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/files"
)

// Must stay in sync with maxDirectUploadSize in the files API.
var directUploadLimit int64 = 100 * 1024 * 1024

// UploadFileOptions configures UploadFile and UploadFileReader.
type UploadFileOptions struct {
	// Purpose defaults to "generic".
	Purpose string
	Title   string
	// FileName overrides the stored file name. UploadFile defaults to the local file name.
	FileName string
	// ContentType defaults to a guess based on the file name.
	ContentType string
	// StoreId and StorePath also place the file in a file store and must be set together.
	StoreId      string
	StorePath    string
	StoreReplace bool
	InstanceId   string
}

type uploadSource struct {
	reader      io.Reader
	name        string
	contentType string
	size        int64
}

// UploadFile uploads a local file and returns the created file. Files over
// 100 MB are sent to a presigned storage URL instead of through the API.
func (sdk *MetorialSdk) UploadFile(ctx context.Context, path string, opts *UploadFileOptions) (*files.FilesGetOutput, error) {
	options := resolveUploadOptions(opts)

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("metorial: %w", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("metorial: %w", err)
	}
	if !info.Mode().IsRegular() {
		return nil, fmt.Errorf("metorial: %s is not a file", path)
	}

	if options.FileName == "" {
		options.FileName = filepath.Base(path)
	}
	return sdk.upload(ctx, uploadSource{reader: file, name: options.FileName, size: info.Size()}, options)
}

// UploadFileReader uploads size bytes read from reader as a file called name
// and returns the created file. FileName in opts is ignored.
func (sdk *MetorialSdk) UploadFileReader(ctx context.Context, name string, reader io.Reader, size int64, opts *UploadFileOptions) (*files.FilesGetOutput, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("metorial: a file name is required")
	}
	if size < 0 {
		return nil, errors.New("metorial: size must not be negative")
	}
	return sdk.upload(ctx, uploadSource{reader: reader, name: name, size: size}, resolveUploadOptions(opts))
}

func resolveUploadOptions(opts *UploadFileOptions) UploadFileOptions {
	options := UploadFileOptions{}
	if opts != nil {
		options = *opts
	}
	if options.Purpose == "" {
		options.Purpose = "generic"
	}
	return options
}

func (sdk *MetorialSdk) upload(ctx context.Context, source uploadSource, options UploadFileOptions) (*files.FilesGetOutput, error) {
	if (options.StoreId == "") != (options.StorePath == "") {
		return nil, errors.New("metorial: StoreId and StorePath must be set together")
	}
	if options.StoreReplace && options.StoreId == "" {
		return nil, errors.New("metorial: StoreReplace requires StoreId and StorePath")
	}

	source.contentType = options.ContentType
	if source.contentType == "" {
		source.contentType = mime.TypeByExtension(filepath.Ext(source.name))
	}
	if source.contentType == "" {
		source.contentType = "application/octet-stream"
	}

	if source.size <= directUploadLimit {
		return sdk.uploadDirect(ctx, source, options)
	}
	return sdk.uploadWithURL(ctx, source, options)
}

func (sdk *MetorialSdk) uploadDirect(ctx context.Context, source uploadSource, options UploadFileOptions) (*files.FilesGetOutput, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fields := [][2]string{{"purpose", options.Purpose}}
	if options.Title != "" {
		fields = append(fields, [2]string{"title", options.Title})
	}
	if options.InstanceId != "" {
		fields = append(fields, [2]string{"instance_id", options.InstanceId})
	}
	if options.StoreId != "" {
		fields = append(fields,
			[2]string{"store_id", options.StoreId},
			[2]string{"path", options.StorePath},
			[2]string{"store_replace", strconv.FormatBool(options.StoreReplace)},
		)
	}
	for _, field := range fields {
		if err := writer.WriteField(field[0], field[1]); err != nil {
			return nil, fmt.Errorf("metorial: %w", err)
		}
	}

	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, quoteEscaper.Replace(source.name)))
	header.Set("Content-Type", source.contentType)
	part, err := writer.CreatePart(header)
	if err != nil {
		return nil, fmt.Errorf("metorial: %w", err)
	}
	if _, err := io.Copy(part, source.reader); err != nil {
		return nil, fmt.Errorf("metorial: failed to read %s: %w", source.name, err)
	}
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("metorial: %w", err)
	}

	response, err := sdk.client.DoStream(ctx, http.MethodPost, sdk.client.URL("files"), body, int64(body.Len()), map[string]string{
		"Content-Type": writer.FormDataContentType(),
	})
	if err != nil {
		return nil, err
	}

	var result files.FilesGetOutput
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, fmt.Errorf("metorial: failed to unmarshal response: %w", err)
	}
	return &result, nil
}

var quoteEscaper = strings.NewReplacer(`\`, `\\`, `"`, `\"`)

func (sdk *MetorialSdk) uploadWithURL(ctx context.Context, source uploadSource, options UploadFileOptions) (*files.FilesGetOutput, error) {
	request := map[string]any{
		"mode":      "get_upload_url",
		"purpose":   options.Purpose,
		"file_name": source.name,
		"file_size": source.size,
		"file_type": source.contentType,
	}
	if options.Title != "" {
		request["title"] = options.Title
	}
	if options.InstanceId != "" {
		request["instance_id"] = options.InstanceId
	}
	if options.StoreId != "" {
		request["store_id"] = options.StoreId
		request["path"] = options.StorePath
		request["store_replace"] = options.StoreReplace
	}

	var session struct {
		Id     string `json:"id"`
		Upload struct {
			Url    string `json:"url"`
			Method string `json:"method"`
		} `json:"upload"`
	}
	if err := sdk.client.Post(&endpoint.Request{Path: []string{"files"}, Body: request}, &session); err != nil {
		return nil, err
	}
	if session.Id == "" || session.Upload.Url == "" {
		return nil, errors.New("metorial: the API did not return an upload URL")
	}

	if err := putToStorage(ctx, session.Upload.Method, session.Upload.Url, source); err != nil {
		return nil, err
	}

	complete := map[string]any{
		"mode":           "complete",
		"file_upload_id": session.Id,
	}
	if options.InstanceId != "" {
		complete["instance_id"] = options.InstanceId
	}

	var result files.FilesGetOutput
	if err := sdk.client.Post(&endpoint.Request{Path: []string{"files"}, Body: complete}, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// putToStorage sends the file to a presigned URL. The signature covers the
// content type, so it must match the declared file_type, and Metorial
// credentials must not be forwarded to the storage backend.
func putToStorage(ctx context.Context, method, url string, source uploadSource) error {
	if method == "" {
		method = http.MethodPut
	}

	request, err := http.NewRequestWithContext(ctx, method, url, source.reader)
	if err != nil {
		return fmt.Errorf("metorial: invalid upload URL: %w", err)
	}
	request.ContentLength = source.size
	request.Header.Set("Content-Type", source.contentType)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("metorial: failed to upload %s: %w", source.name, err)
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		detail, _ := io.ReadAll(io.LimitReader(response.Body, 2048))
		message := strings.TrimSpace(string(detail))
		if message == "" {
			return fmt.Errorf("metorial: failed to upload %s: storage responded with %s", source.name, response.Status)
		}
		return fmt.Errorf("metorial: failed to upload %s: storage responded with %s: %s", source.name, response.Status, message)
	}

	_, _ = io.Copy(io.Discard, response.Body)
	return nil
}
