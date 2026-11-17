package metorial

import (
	"encoding/json"
	"math/big"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var apiKeyHostPattern = regexp.MustCompile(`^metorial_[a-z]+_[A-Za-z0-9]{60}([A-Za-z0-9]+)v[12]$`)

var apiKeyBase62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func apiHostFromToken(token string) string {
	if len(token) > 8192 {
		return ""
	}

	match := apiKeyHostPattern.FindStringSubmatch(token)
	if match == nil {
		return ""
	}

	value := new(big.Int)
	base := big.NewInt(62)
	for _, character := range match[1] {
		value.Mul(value, base)
		value.Add(value, big.NewInt(int64(strings.IndexRune(apiKeyBase62Alphabet, character))))
	}

	data := value.Bytes()
	var config []json.RawMessage
	if !utf8.Valid(data) || json.Unmarshal(data, &config) != nil || len(config) != 2 {
		return ""
	}

	var host, instance string
	if json.Unmarshal(config[0], &host) != nil || json.Unmarshal(config[1], &instance) != nil || instance == "" {
		return ""
	}
	if strings.ContainsAny(host, `\?#`) || strings.IndexFunc(host, unicode.IsSpace) >= 0 {
		return ""
	}

	parsed, err := url.Parse(host)
	if err != nil || parsed.Hostname() == "" || parsed.User != nil ||
		(!strings.EqualFold(parsed.Scheme, "http") && !strings.EqualFold(parsed.Scheme, "https")) {
		return ""
	}
	if port := parsed.Port(); port != "" {
		number, err := strconv.Atoi(port)
		if err != nil || number > 65535 {
			return ""
		}
	}

	return host
}
