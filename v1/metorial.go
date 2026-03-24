// Package metorial provides a Go client for the Metorial API.
//
// The MetorialSdk allows you to interact with the Metorial platform
// for managing MCP server integrations, sessions, and providers.
//
// Create a client with your API key:
//
//	sdk, err := metorial.New(metorial.WithAPIKey("metorial_sk_..."))
//	if err != nil {
//		log.Fatal(err)
//	}
//
// Then use the endpoint fields to interact with the API:
//
//	// List providers
//	providers, err := sdk.Providers.List(nil)
//
//	// Get a specific provider
//	provider, err := sdk.Providers.Get("provider_id")
//
//	// Management endpoints require an instance ID
//	providers, err := sdk.Management.Providers.List("inst_id", nil)
package metorial

import (
	"fmt"

	"github.com/metorial/metorial-go/v1/endpoints"
	"github.com/metorial/metorial-go/v1/endpoints/management"
	"github.com/metorial/metorial-go/v1/internal/endpoint"
)

const apiVersion = "2026-01-01-magnetar"

// ManagementEndpoints provides access to management API endpoints.
// Management endpoints operate across instances and require an instance ID parameter.
type ManagementEndpoints struct {
	Instance                              *management.InstanceEndpoint
	Providers                             *management.ProvidersEndpoint
	ProvidersVersions                     *management.ProvidersVersionsEndpoint
	ProvidersTools                        *management.ProvidersToolsEndpoint
	ProvidersAuthMethods                  *management.ProvidersAuthMethodsEndpoint
	ProvidersSpecifications               *management.ProvidersSpecificationsEndpoint
	Publishers                            *management.PublishersEndpoint
	ProviderCategories                    *management.ProviderCategoriesEndpoint
	ProviderCollections                   *management.ProviderCollectionsEndpoint
	ProviderGroups                        *management.ProviderGroupsEndpoint
	ProviderListings                      *management.ProviderListingsEndpoint
	ProviderDeployments                   *management.ProviderDeploymentsEndpoint
	ProviderDeploymentsConfigs            *management.ProviderDeploymentsConfigsEndpoint
	ProviderDeploymentsConfigVaults       *management.ProviderDeploymentsConfigVaultsEndpoint
	ProviderDeploymentsAuthConfigs        *management.ProviderDeploymentsAuthConfigsEndpoint
	ProviderDeploymentsAuthConfigsImports *management.ProviderDeploymentsAuthConfigsImportsEndpoint
	ProviderDeploymentsAuthConfigsExports *management.ProviderDeploymentsAuthConfigsExportsEndpoint
	ProviderDeploymentsAuthCredentials    *management.ProviderDeploymentsAuthCredentialsEndpoint
	ProviderDeploymentsSetupSessions      *management.ProviderDeploymentsSetupSessionsEndpoint
	ProviderRuns                          *management.ProviderRunsEndpoint
	ProviderTemplates                     *management.ProviderTemplatesEndpoint
	Sessions                              *management.SessionsEndpoint
	SessionsMessages                      *management.SessionsMessagesEndpoint
	SessionsConnections                   *management.SessionsConnectionsEndpoint
	SessionsEvents                        *management.SessionsEventsEndpoint
	SessionsProviders                     *management.SessionsProvidersEndpoint
	SessionsParticipants                  *management.SessionsParticipantsEndpoint
	SessionsErrors                        *management.SessionsErrorsEndpoint
	SessionsErrorGroups                   *management.SessionsErrorGroupsEndpoint
	SessionTemplates                      *management.SessionTemplatesEndpoint
	SessionTemplatesProviders             *management.SessionTemplatesProvidersEndpoint
	ToolCalls                             *management.ToolCallsEndpoint
	CustomProviders                       *management.CustomProvidersEndpoint
	CustomProvidersVersions               *management.CustomProvidersVersionsEndpoint
	CustomProvidersDeployments            *management.CustomProvidersDeploymentsEndpoint
	CustomProvidersEnvironments           *management.CustomProvidersEnvironmentsEndpoint
	CustomProvidersCommits                *management.CustomProvidersCommitsEndpoint
	Identities                            *management.IdentitiesEndpoint
	IdentitiesCredentials                 *management.IdentitiesCredentialsEndpoint
	IdentitiesDelegations                 *management.IdentitiesDelegationsEndpoint
	IdentitiesDelegationConfigs           *management.IdentitiesDelegationConfigsEndpoint
	IdentitiesDelegationRequests          *management.IdentitiesDelegationRequestsEndpoint
	IdentityActors                        *management.IdentityActorsEndpoint
	Files                                 *management.FilesEndpoint
	FileLinks                             *management.FileLinksEndpoint
	Portals                               *management.PortalsEndpoint
	MagicMcpGroups                        *management.MagicMcpGroupsEndpoint
	MagicMcpServers                       *management.MagicMcpServersEndpoint
	MagicMcpSessions                      *management.MagicMcpSessionsEndpoint
	MagicMcpTokens                        *management.MagicMcpTokensEndpoint
}

// MetorialSdk is the main entry point for the Metorial API.
type MetorialSdk struct {
	client *endpoint.Client

	// Public API endpoints (scoped to the current instance via API key).
	Instance                              *endpoints.InstanceEndpoint
	Instances                             *endpoints.InstancesEndpoint
	Providers                             *endpoints.ProvidersEndpoint
	ProvidersVersions                     *endpoints.ProvidersVersionsEndpoint
	ProvidersTools                        *endpoints.ProvidersToolsEndpoint
	ProvidersAuthMethods                  *endpoints.ProvidersAuthMethodsEndpoint
	ProvidersSpecifications               *endpoints.ProvidersSpecificationsEndpoint
	Publishers                            *endpoints.PublishersEndpoint
	ProviderCategories                    *endpoints.ProviderCategoriesEndpoint
	ProviderCollections                   *endpoints.ProviderCollectionsEndpoint
	ProviderGroups                        *endpoints.ProviderGroupsEndpoint
	ProviderListings                      *endpoints.ProviderListingsEndpoint
	ProviderDeployments                   *endpoints.ProviderDeploymentsEndpoint
	ProviderDeploymentsConfigs            *endpoints.ProviderDeploymentsConfigsEndpoint
	ProviderDeploymentsConfigVaults       *endpoints.ProviderDeploymentsConfigVaultsEndpoint
	ProviderDeploymentsAuthConfigs        *endpoints.ProviderDeploymentsAuthConfigsEndpoint
	ProviderDeploymentsAuthConfigsImports *endpoints.ProviderDeploymentsAuthConfigsImportsEndpoint
	ProviderDeploymentsAuthConfigsExports *endpoints.ProviderDeploymentsAuthConfigsExportsEndpoint
	ProviderDeploymentsAuthCredentials    *endpoints.ProviderDeploymentsAuthCredentialsEndpoint
	ProviderDeploymentsSetupSessions      *endpoints.ProviderDeploymentsSetupSessionsEndpoint
	ProviderRuns                          *endpoints.ProviderRunsEndpoint
	ProviderTemplates                     *endpoints.ProviderTemplatesEndpoint
	Sessions                              *endpoints.SessionsEndpoint
	SessionsMessages                      *endpoints.SessionsMessagesEndpoint
	SessionsConnections                   *endpoints.SessionsConnectionsEndpoint
	SessionsEvents                        *endpoints.SessionsEventsEndpoint
	SessionsProviders                     *endpoints.SessionsProvidersEndpoint
	SessionsParticipants                  *endpoints.SessionsParticipantsEndpoint
	SessionsErrors                        *endpoints.SessionsErrorsEndpoint
	SessionsErrorGroups                   *endpoints.SessionsErrorGroupsEndpoint
	SessionTemplates                      *endpoints.SessionTemplatesEndpoint
	SessionTemplatesProviders             *endpoints.SessionTemplatesProvidersEndpoint
	ToolCalls                             *endpoints.ToolCallsEndpoint
	CustomProviders                       *endpoints.CustomProvidersEndpoint
	CustomProvidersVersions               *endpoints.CustomProvidersVersionsEndpoint
	CustomProvidersDeployments            *endpoints.CustomProvidersDeploymentsEndpoint
	CustomProvidersEnvironments           *endpoints.CustomProvidersEnvironmentsEndpoint
	CustomProvidersCommits                *endpoints.CustomProvidersCommitsEndpoint
	Identities                            *endpoints.IdentitiesEndpoint
	IdentitiesCredentials                 *endpoints.IdentitiesCredentialsEndpoint
	IdentitiesDelegations                 *endpoints.IdentitiesDelegationsEndpoint
	IdentitiesDelegationConfigs           *endpoints.IdentitiesDelegationConfigsEndpoint
	IdentitiesDelegationRequests          *endpoints.IdentitiesDelegationRequestsEndpoint
	IdentityActors                        *endpoints.IdentityActorsEndpoint
	Files                                 *endpoints.FilesEndpoint
	FileLinks                             *endpoints.FileLinksEndpoint
	Portals                               *endpoints.PortalsEndpoint
	MagicMcpGroups                        *endpoints.MagicMcpGroupsEndpoint
	MagicMcpServers                       *endpoints.MagicMcpServersEndpoint
	MagicMcpSessions                      *endpoints.MagicMcpSessionsEndpoint
	MagicMcpTokens                        *endpoints.MagicMcpTokensEndpoint

	// Management provides access to management API endpoints.
	// These endpoints operate across instances and require an instance ID parameter.
	Management *ManagementEndpoints
}

// New creates a new MetorialSdk with the given options.
// At minimum, an API key must be provided via WithAPIKey.
func New(opts ...Option) (*MetorialSdk, error) {
	o := &options{
		apiHost: defaultAPIHost,
	}
	for _, opt := range opts {
		opt(o)
	}

	if o.apiKey == "" {
		return nil, fmt.Errorf("metorial: API key is required, use WithAPIKey option")
	}

	c := endpoint.NewClient(o.apiHost, o.apiKey, apiVersion, o.headers)

	return &MetorialSdk{
		client:                                c,
		Instance:                              endpoints.NewInstanceEndpoint(c),
		Instances:                             endpoints.NewInstancesEndpoint(c),
		Providers:                             endpoints.NewProvidersEndpoint(c),
		ProvidersVersions:                     endpoints.NewProvidersVersionsEndpoint(c),
		ProvidersTools:                        endpoints.NewProvidersToolsEndpoint(c),
		ProvidersAuthMethods:                  endpoints.NewProvidersAuthMethodsEndpoint(c),
		ProvidersSpecifications:               endpoints.NewProvidersSpecificationsEndpoint(c),
		Publishers:                            endpoints.NewPublishersEndpoint(c),
		ProviderCategories:                    endpoints.NewProviderCategoriesEndpoint(c),
		ProviderCollections:                   endpoints.NewProviderCollectionsEndpoint(c),
		ProviderGroups:                        endpoints.NewProviderGroupsEndpoint(c),
		ProviderListings:                      endpoints.NewProviderListingsEndpoint(c),
		ProviderDeployments:                   endpoints.NewProviderDeploymentsEndpoint(c),
		ProviderDeploymentsConfigs:            endpoints.NewProviderDeploymentsConfigsEndpoint(c),
		ProviderDeploymentsConfigVaults:       endpoints.NewProviderDeploymentsConfigVaultsEndpoint(c),
		ProviderDeploymentsAuthConfigs:        endpoints.NewProviderDeploymentsAuthConfigsEndpoint(c),
		ProviderDeploymentsAuthConfigsImports: endpoints.NewProviderDeploymentsAuthConfigsImportsEndpoint(c),
		ProviderDeploymentsAuthConfigsExports: endpoints.NewProviderDeploymentsAuthConfigsExportsEndpoint(c),
		ProviderDeploymentsAuthCredentials:    endpoints.NewProviderDeploymentsAuthCredentialsEndpoint(c),
		ProviderDeploymentsSetupSessions:      endpoints.NewProviderDeploymentsSetupSessionsEndpoint(c),
		ProviderRuns:                          endpoints.NewProviderRunsEndpoint(c),
		ProviderTemplates:                     endpoints.NewProviderTemplatesEndpoint(c),
		Sessions:                              endpoints.NewSessionsEndpoint(c),
		SessionsMessages:                      endpoints.NewSessionsMessagesEndpoint(c),
		SessionsConnections:                   endpoints.NewSessionsConnectionsEndpoint(c),
		SessionsEvents:                        endpoints.NewSessionsEventsEndpoint(c),
		SessionsProviders:                     endpoints.NewSessionsProvidersEndpoint(c),
		SessionsParticipants:                  endpoints.NewSessionsParticipantsEndpoint(c),
		SessionsErrors:                        endpoints.NewSessionsErrorsEndpoint(c),
		SessionsErrorGroups:                   endpoints.NewSessionsErrorGroupsEndpoint(c),
		SessionTemplates:                      endpoints.NewSessionTemplatesEndpoint(c),
		SessionTemplatesProviders:             endpoints.NewSessionTemplatesProvidersEndpoint(c),
		ToolCalls:                             endpoints.NewToolCallsEndpoint(c),
		CustomProviders:                       endpoints.NewCustomProvidersEndpoint(c),
		CustomProvidersVersions:               endpoints.NewCustomProvidersVersionsEndpoint(c),
		CustomProvidersDeployments:            endpoints.NewCustomProvidersDeploymentsEndpoint(c),
		CustomProvidersEnvironments:           endpoints.NewCustomProvidersEnvironmentsEndpoint(c),
		CustomProvidersCommits:                endpoints.NewCustomProvidersCommitsEndpoint(c),
		Identities:                            endpoints.NewIdentitiesEndpoint(c),
		IdentitiesCredentials:                 endpoints.NewIdentitiesCredentialsEndpoint(c),
		IdentitiesDelegations:                 endpoints.NewIdentitiesDelegationsEndpoint(c),
		IdentitiesDelegationConfigs:           endpoints.NewIdentitiesDelegationConfigsEndpoint(c),
		IdentitiesDelegationRequests:          endpoints.NewIdentitiesDelegationRequestsEndpoint(c),
		IdentityActors:                        endpoints.NewIdentityActorsEndpoint(c),
		Files:                                 endpoints.NewFilesEndpoint(c),
		FileLinks:                             endpoints.NewFileLinksEndpoint(c),
		Portals:                               endpoints.NewPortalsEndpoint(c),
		MagicMcpGroups:                        endpoints.NewMagicMcpGroupsEndpoint(c),
		MagicMcpServers:                       endpoints.NewMagicMcpServersEndpoint(c),
		MagicMcpSessions:                      endpoints.NewMagicMcpSessionsEndpoint(c),
		MagicMcpTokens:                        endpoints.NewMagicMcpTokensEndpoint(c),
		Management: &ManagementEndpoints{
			Providers:                             management.NewProvidersEndpoint(c),
			ProvidersVersions:                     management.NewProvidersVersionsEndpoint(c),
			ProvidersTools:                        management.NewProvidersToolsEndpoint(c),
			ProvidersAuthMethods:                  management.NewProvidersAuthMethodsEndpoint(c),
			ProvidersSpecifications:               management.NewProvidersSpecificationsEndpoint(c),
			Publishers:                            management.NewPublishersEndpoint(c),
			ProviderCategories:                    management.NewProviderCategoriesEndpoint(c),
			ProviderCollections:                   management.NewProviderCollectionsEndpoint(c),
			ProviderGroups:                        management.NewProviderGroupsEndpoint(c),
			ProviderListings:                      management.NewProviderListingsEndpoint(c),
			ProviderDeployments:                   management.NewProviderDeploymentsEndpoint(c),
			ProviderDeploymentsConfigs:            management.NewProviderDeploymentsConfigsEndpoint(c),
			ProviderDeploymentsConfigVaults:       management.NewProviderDeploymentsConfigVaultsEndpoint(c),
			ProviderDeploymentsAuthConfigs:        management.NewProviderDeploymentsAuthConfigsEndpoint(c),
			ProviderDeploymentsAuthConfigsImports: management.NewProviderDeploymentsAuthConfigsImportsEndpoint(c),
			ProviderDeploymentsAuthConfigsExports: management.NewProviderDeploymentsAuthConfigsExportsEndpoint(c),
			ProviderDeploymentsAuthCredentials:    management.NewProviderDeploymentsAuthCredentialsEndpoint(c),
			ProviderDeploymentsSetupSessions:      management.NewProviderDeploymentsSetupSessionsEndpoint(c),
			ProviderRuns:                          management.NewProviderRunsEndpoint(c),
			ProviderTemplates:                     management.NewProviderTemplatesEndpoint(c),
			Sessions:                              management.NewSessionsEndpoint(c),
			SessionsMessages:                      management.NewSessionsMessagesEndpoint(c),
			SessionsConnections:                   management.NewSessionsConnectionsEndpoint(c),
			SessionsEvents:                        management.NewSessionsEventsEndpoint(c),
			SessionsProviders:                     management.NewSessionsProvidersEndpoint(c),
			SessionsParticipants:                  management.NewSessionsParticipantsEndpoint(c),
			SessionsErrors:                        management.NewSessionsErrorsEndpoint(c),
			SessionsErrorGroups:                   management.NewSessionsErrorGroupsEndpoint(c),
			SessionTemplates:                      management.NewSessionTemplatesEndpoint(c),
			SessionTemplatesProviders:             management.NewSessionTemplatesProvidersEndpoint(c),
			ToolCalls:                             management.NewToolCallsEndpoint(c),
			CustomProviders:                       management.NewCustomProvidersEndpoint(c),
			CustomProvidersVersions:               management.NewCustomProvidersVersionsEndpoint(c),
			CustomProvidersDeployments:            management.NewCustomProvidersDeploymentsEndpoint(c),
			CustomProvidersEnvironments:           management.NewCustomProvidersEnvironmentsEndpoint(c),
			CustomProvidersCommits:                management.NewCustomProvidersCommitsEndpoint(c),
			Identities:                            management.NewIdentitiesEndpoint(c),
			IdentitiesCredentials:                 management.NewIdentitiesCredentialsEndpoint(c),
			IdentitiesDelegations:                 management.NewIdentitiesDelegationsEndpoint(c),
			IdentitiesDelegationConfigs:           management.NewIdentitiesDelegationConfigsEndpoint(c),
			IdentitiesDelegationRequests:          management.NewIdentitiesDelegationRequestsEndpoint(c),
			IdentityActors:                        management.NewIdentityActorsEndpoint(c),
			Files:                                 management.NewFilesEndpoint(c),
			FileLinks:                             management.NewFileLinksEndpoint(c),
			Portals:                               management.NewPortalsEndpoint(c),
			MagicMcpGroups:                        management.NewMagicMcpGroupsEndpoint(c),
			MagicMcpServers:                       management.NewMagicMcpServersEndpoint(c),
			MagicMcpSessions:                      management.NewMagicMcpSessionsEndpoint(c),
			MagicMcpTokens:                        management.NewMagicMcpTokensEndpoint(c),
		},
	}, nil
}
