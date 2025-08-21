package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetResourcesRequest represents the GetResourcesRequest schema from the OpenAPI specification
type GetResourcesRequest struct {
}

// ClientCertificates represents the ClientCertificates schema from the OpenAPI specification
type ClientCertificates struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// DeleteApiKeyRequest represents the DeleteApiKeyRequest schema from the OpenAPI specification
type DeleteApiKeyRequest struct {
}

// GetApiKeysRequest represents the GetApiKeysRequest schema from the OpenAPI specification
type GetApiKeysRequest struct {
}

// Authorizers represents the Authorizers schema from the OpenAPI specification
type Authorizers struct {
	Position string `json:"position,omitempty"`
	Items interface{} `json:"items,omitempty"`
}

// UpdateIntegrationResponseRequest represents the UpdateIntegrationResponseRequest schema from the OpenAPI specification
type UpdateIntegrationResponseRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// UpdateDocumentationVersionRequest represents the UpdateDocumentationVersionRequest schema from the OpenAPI specification
type UpdateDocumentationVersionRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// GetSdkTypeRequest represents the GetSdkTypeRequest schema from the OpenAPI specification
type GetSdkTypeRequest struct {
}

// DocumentationVersions represents the DocumentationVersions schema from the OpenAPI specification
type DocumentationVersions struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// DeleteStageRequest represents the DeleteStageRequest schema from the OpenAPI specification
type DeleteStageRequest struct {
}

// Account represents the Account schema from the OpenAPI specification
type Account struct {
	Features interface{} `json:"features,omitempty"`
	Throttlesettings interface{} `json:"throttleSettings,omitempty"`
	Apikeyversion interface{} `json:"apiKeyVersion,omitempty"`
	Cloudwatchrolearn interface{} `json:"cloudwatchRoleArn,omitempty"`
}

// MethodSetting represents the MethodSetting schema from the OpenAPI specification
type MethodSetting struct {
	Throttlingburstlimit interface{} `json:"throttlingBurstLimit,omitempty"`
	Throttlingratelimit interface{} `json:"throttlingRateLimit,omitempty"`
	Cachettlinseconds interface{} `json:"cacheTtlInSeconds,omitempty"`
	Requireauthorizationforcachecontrol interface{} `json:"requireAuthorizationForCacheControl,omitempty"`
	Unauthorizedcachecontrolheaderstrategy interface{} `json:"unauthorizedCacheControlHeaderStrategy,omitempty"`
	Cachedataencrypted interface{} `json:"cacheDataEncrypted,omitempty"`
	Metricsenabled interface{} `json:"metricsEnabled,omitempty"`
	Cachingenabled interface{} `json:"cachingEnabled,omitempty"`
	Datatraceenabled interface{} `json:"dataTraceEnabled,omitempty"`
	Logginglevel interface{} `json:"loggingLevel,omitempty"`
}

// Model represents the Model schema from the OpenAPI specification
type Model struct {
	Name interface{} `json:"name,omitempty"`
	Schema interface{} `json:"schema,omitempty"`
	Contenttype interface{} `json:"contentType,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id,omitempty"`
}

// DeleteResourceRequest represents the DeleteResourceRequest schema from the OpenAPI specification
type DeleteResourceRequest struct {
}

// DeleteVpcLinkRequest represents the DeleteVpcLinkRequest schema from the OpenAPI specification
type DeleteVpcLinkRequest struct {
}

// GetRestApiRequest represents the GetRestApiRequest schema from the OpenAPI specification
type GetRestApiRequest struct {
}

// GenerateClientCertificateRequest represents the GenerateClientCertificateRequest schema from the OpenAPI specification
type GenerateClientCertificateRequest struct {
	Description interface{} `json:"description,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// MapOfStringToString represents the MapOfStringToString schema from the OpenAPI specification
type MapOfStringToString struct {
}

// GetExportRequest represents the GetExportRequest schema from the OpenAPI specification
type GetExportRequest struct {
}

// UsagePlans represents the UsagePlans schema from the OpenAPI specification
type UsagePlans struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// DeleteModelRequest represents the DeleteModelRequest schema from the OpenAPI specification
type DeleteModelRequest struct {
}

// GetDomainNamesRequest represents the GetDomainNamesRequest schema from the OpenAPI specification
type GetDomainNamesRequest struct {
}

// CreateDeploymentRequest represents the CreateDeploymentRequest schema from the OpenAPI specification
type CreateDeploymentRequest struct {
	Variables interface{} `json:"variables,omitempty"`
	Cacheclusterenabled interface{} `json:"cacheClusterEnabled,omitempty"`
	Cacheclustersize interface{} `json:"cacheClusterSize,omitempty"`
	Canarysettings interface{} `json:"canarySettings,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Stagedescription interface{} `json:"stageDescription,omitempty"`
	Stagename interface{} `json:"stageName,omitempty"`
	Tracingenabled interface{} `json:"tracingEnabled,omitempty"`
}

// TestInvokeMethodRequest represents the TestInvokeMethodRequest schema from the OpenAPI specification
type TestInvokeMethodRequest struct {
	Clientcertificateid interface{} `json:"clientCertificateId,omitempty"`
	Headers interface{} `json:"headers,omitempty"`
	Multivalueheaders interface{} `json:"multiValueHeaders,omitempty"`
	Pathwithquerystring interface{} `json:"pathWithQueryString,omitempty"`
	Stagevariables interface{} `json:"stageVariables,omitempty"`
	Body interface{} `json:"body,omitempty"`
}

// PutGatewayResponseRequest represents the PutGatewayResponseRequest schema from the OpenAPI specification
type PutGatewayResponseRequest struct {
	Statuscode interface{} `json:"statusCode,omitempty"`
	Responseparameters interface{} `json:"responseParameters,omitempty"`
	Responsetemplates interface{} `json:"responseTemplates,omitempty"`
}

// BasePathMappings represents the BasePathMappings schema from the OpenAPI specification
type BasePathMappings struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// MutualTlsAuthenticationInput represents the MutualTlsAuthenticationInput schema from the OpenAPI specification
type MutualTlsAuthenticationInput struct {
	Truststoreuri interface{} `json:"truststoreUri,omitempty"`
	Truststoreversion interface{} `json:"truststoreVersion,omitempty"`
}

// Stages represents the Stages schema from the OpenAPI specification
type Stages struct {
	Item interface{} `json:"item,omitempty"`
}

// GetDeploymentsRequest represents the GetDeploymentsRequest schema from the OpenAPI specification
type GetDeploymentsRequest struct {
}

// GetApiKeyRequest represents the GetApiKeyRequest schema from the OpenAPI specification
type GetApiKeyRequest struct {
}

// UpdateIntegrationRequest represents the UpdateIntegrationRequest schema from the OpenAPI specification
type UpdateIntegrationRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// Resource represents the Resource schema from the OpenAPI specification
type Resource struct {
	Parentid interface{} `json:"parentId,omitempty"`
	Path interface{} `json:"path,omitempty"`
	Pathpart interface{} `json:"pathPart,omitempty"`
	Resourcemethods interface{} `json:"resourceMethods,omitempty"`
	Id interface{} `json:"id,omitempty"`
}

// CreateRestApiRequest represents the CreateRestApiRequest schema from the OpenAPI specification
type CreateRestApiRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	Apikeysource interface{} `json:"apiKeySource,omitempty"`
	Clonefrom interface{} `json:"cloneFrom,omitempty"`
	Endpointconfiguration interface{} `json:"endpointConfiguration,omitempty"`
	Policy interface{} `json:"policy,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Minimumcompressionsize interface{} `json:"minimumCompressionSize,omitempty"`
	Name interface{} `json:"name"`
	Binarymediatypes interface{} `json:"binaryMediaTypes,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Disableexecuteapiendpoint interface{} `json:"disableExecuteApiEndpoint,omitempty"`
}

// GetSdkRequest represents the GetSdkRequest schema from the OpenAPI specification
type GetSdkRequest struct {
}

// CreateModelRequest represents the CreateModelRequest schema from the OpenAPI specification
type CreateModelRequest struct {
	Contenttype interface{} `json:"contentType"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
	Schema interface{} `json:"schema,omitempty"`
}

// DocumentationVersion represents the DocumentationVersion schema from the OpenAPI specification
type DocumentationVersion struct {
	Description interface{} `json:"description,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
}

// StageKey represents the StageKey schema from the OpenAPI specification
type StageKey struct {
	Stagename interface{} `json:"stageName,omitempty"`
	Restapiid interface{} `json:"restApiId,omitempty"`
}

// UpdateApiKeyRequest represents the UpdateApiKeyRequest schema from the OpenAPI specification
type UpdateApiKeyRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// Template represents the Template schema from the OpenAPI specification
type Template struct {
	Value interface{} `json:"value,omitempty"`
}

// DeleteDocumentationVersionRequest represents the DeleteDocumentationVersionRequest schema from the OpenAPI specification
type DeleteDocumentationVersionRequest struct {
}

// GetBasePathMappingsRequest represents the GetBasePathMappingsRequest schema from the OpenAPI specification
type GetBasePathMappingsRequest struct {
}

// QuotaSettings represents the QuotaSettings schema from the OpenAPI specification
type QuotaSettings struct {
	Offset interface{} `json:"offset,omitempty"`
	Period interface{} `json:"period,omitempty"`
	Limit interface{} `json:"limit,omitempty"`
}

// CreateRequestValidatorRequest represents the CreateRequestValidatorRequest schema from the OpenAPI specification
type CreateRequestValidatorRequest struct {
	Name interface{} `json:"name,omitempty"`
	Validaterequestbody interface{} `json:"validateRequestBody,omitempty"`
	Validaterequestparameters interface{} `json:"validateRequestParameters,omitempty"`
}

// ImportApiKeysRequest represents the ImportApiKeysRequest schema from the OpenAPI specification
type ImportApiKeysRequest struct {
	Body interface{} `json:"body"`
}

// ApiKeyIds represents the ApiKeyIds schema from the OpenAPI specification
type ApiKeyIds struct {
	Ids interface{} `json:"ids,omitempty"`
	Warnings interface{} `json:"warnings,omitempty"`
}

// GetDocumentationPartsRequest represents the GetDocumentationPartsRequest schema from the OpenAPI specification
type GetDocumentationPartsRequest struct {
}

// TlsConfig represents the TlsConfig schema from the OpenAPI specification
type TlsConfig struct {
	Insecureskipverification interface{} `json:"insecureSkipVerification,omitempty"`
}

// GetBasePathMappingRequest represents the GetBasePathMappingRequest schema from the OpenAPI specification
type GetBasePathMappingRequest struct {
}

// UpdateRequestValidatorRequest represents the UpdateRequestValidatorRequest schema from the OpenAPI specification
type UpdateRequestValidatorRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// ApiKey represents the ApiKey schema from the OpenAPI specification
type ApiKey struct {
	Customerid interface{} `json:"customerId,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Lastupdateddate interface{} `json:"lastUpdatedDate,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Stagekeys interface{} `json:"stageKeys,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// FlushStageCacheRequest represents the FlushStageCacheRequest schema from the OpenAPI specification
type FlushStageCacheRequest struct {
}

// GetModelTemplateRequest represents the GetModelTemplateRequest schema from the OpenAPI specification
type GetModelTemplateRequest struct {
}

// DeleteIntegrationResponseRequest represents the DeleteIntegrationResponseRequest schema from the OpenAPI specification
type DeleteIntegrationResponseRequest struct {
}

// UpdateRestApiRequest represents the UpdateRestApiRequest schema from the OpenAPI specification
type UpdateRestApiRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// DeleteDocumentationPartRequest represents the DeleteDocumentationPartRequest schema from the OpenAPI specification
type DeleteDocumentationPartRequest struct {
}

// MapOfMethod represents the MapOfMethod schema from the OpenAPI specification
type MapOfMethod struct {
}

// UpdateAccountRequest represents the UpdateAccountRequest schema from the OpenAPI specification
type UpdateAccountRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// GetUsagePlanKeysRequest represents the GetUsagePlanKeysRequest schema from the OpenAPI specification
type GetUsagePlanKeysRequest struct {
}

// GetUsageRequest represents the GetUsageRequest schema from the OpenAPI specification
type GetUsageRequest struct {
}

// UpdateDomainNameRequest represents the UpdateDomainNameRequest schema from the OpenAPI specification
type UpdateDomainNameRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// DomainNames represents the DomainNames schema from the OpenAPI specification
type DomainNames struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// DocumentationPartIds represents the DocumentationPartIds schema from the OpenAPI specification
type DocumentationPartIds struct {
	Ids interface{} `json:"ids,omitempty"`
	Warnings interface{} `json:"warnings,omitempty"`
}

// GetDocumentationPartRequest represents the GetDocumentationPartRequest schema from the OpenAPI specification
type GetDocumentationPartRequest struct {
}

// Models represents the Models schema from the OpenAPI specification
type Models struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// GetSdkTypesRequest represents the GetSdkTypesRequest schema from the OpenAPI specification
type GetSdkTypesRequest struct {
}

// GetVpcLinksRequest represents the GetVpcLinksRequest schema from the OpenAPI specification
type GetVpcLinksRequest struct {
}

// VpcLinks represents the VpcLinks schema from the OpenAPI specification
type VpcLinks struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// SdkType represents the SdkType schema from the OpenAPI specification
type SdkType struct {
	Description interface{} `json:"description,omitempty"`
	Friendlyname interface{} `json:"friendlyName,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Configurationproperties interface{} `json:"configurationProperties,omitempty"`
}

// CreateUsagePlanRequest represents the CreateUsagePlanRequest schema from the OpenAPI specification
type CreateUsagePlanRequest struct {
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
	Quota interface{} `json:"quota,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Throttle interface{} `json:"throttle,omitempty"`
	Apistages interface{} `json:"apiStages,omitempty"`
}

// CreateVpcLinkRequest represents the CreateVpcLinkRequest schema from the OpenAPI specification
type CreateVpcLinkRequest struct {
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
	Targetarns interface{} `json:"targetArns"`
}

// UpdateModelRequest represents the UpdateModelRequest schema from the OpenAPI specification
type UpdateModelRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// GetStagesRequest represents the GetStagesRequest schema from the OpenAPI specification
type GetStagesRequest struct {
}

// Authorizer represents the Authorizer schema from the OpenAPI specification
type Authorizer struct {
	TypeField interface{} `json:"type,omitempty"`
	Authtype interface{} `json:"authType,omitempty"`
	Authorizercredentials interface{} `json:"authorizerCredentials,omitempty"`
	Identitysource interface{} `json:"identitySource,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Authorizeruri interface{} `json:"authorizerUri,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Identityvalidationexpression interface{} `json:"identityValidationExpression,omitempty"`
	Authorizerresultttlinseconds interface{} `json:"authorizerResultTtlInSeconds,omitempty"`
	Providerarns interface{} `json:"providerARNs,omitempty"`
}

// ClientCertificate represents the ClientCertificate schema from the OpenAPI specification
type ClientCertificate struct {
	Clientcertificateid interface{} `json:"clientCertificateId,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Expirationdate interface{} `json:"expirationDate,omitempty"`
	Pemencodedcertificate interface{} `json:"pemEncodedCertificate,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// UpdateClientCertificateRequest represents the UpdateClientCertificateRequest schema from the OpenAPI specification
type UpdateClientCertificateRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// CreateDomainNameRequest represents the CreateDomainNameRequest schema from the OpenAPI specification
type CreateDomainNameRequest struct {
	Securitypolicy interface{} `json:"securityPolicy,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Certificatechain interface{} `json:"certificateChain,omitempty"`
	Endpointconfiguration interface{} `json:"endpointConfiguration,omitempty"`
	Mutualtlsauthentication MutualTlsAuthenticationInput `json:"mutualTlsAuthentication,omitempty"` // The mutual TLS authentication configuration for a custom domain name. If specified, API Gateway performs two-way authentication between the client and the server. Clients must present a trusted certificate to access your API.
	Regionalcertificatearn interface{} `json:"regionalCertificateArn,omitempty"`
	Ownershipverificationcertificatearn interface{} `json:"ownershipVerificationCertificateArn,omitempty"`
	Regionalcertificatename interface{} `json:"regionalCertificateName,omitempty"`
	Certificatearn interface{} `json:"certificateArn,omitempty"`
	Certificatebody interface{} `json:"certificateBody,omitempty"`
	Certificateprivatekey interface{} `json:"certificatePrivateKey,omitempty"`
	Domainname interface{} `json:"domainName"`
	Certificatename interface{} `json:"certificateName,omitempty"`
}

// GetModelRequest represents the GetModelRequest schema from the OpenAPI specification
type GetModelRequest struct {
}

// DeleteBasePathMappingRequest represents the DeleteBasePathMappingRequest schema from the OpenAPI specification
type DeleteBasePathMappingRequest struct {
}

// PutIntegrationResponseRequest represents the PutIntegrationResponseRequest schema from the OpenAPI specification
type PutIntegrationResponseRequest struct {
	Contenthandling interface{} `json:"contentHandling,omitempty"`
	Responseparameters interface{} `json:"responseParameters,omitempty"`
	Responsetemplates interface{} `json:"responseTemplates,omitempty"`
	Selectionpattern interface{} `json:"selectionPattern,omitempty"`
}

// GetResourceRequest represents the GetResourceRequest schema from the OpenAPI specification
type GetResourceRequest struct {
}

// CreateApiKeyRequest represents the CreateApiKeyRequest schema from the OpenAPI specification
type CreateApiKeyRequest struct {
	Stagekeys interface{} `json:"stageKeys,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Customerid interface{} `json:"customerId,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
	Generatedistinctid interface{} `json:"generateDistinctId,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// PathToMapOfMethodSnapshot represents the PathToMapOfMethodSnapshot schema from the OpenAPI specification
type PathToMapOfMethodSnapshot struct {
}

// GetRequestValidatorRequest represents the GetRequestValidatorRequest schema from the OpenAPI specification
type GetRequestValidatorRequest struct {
}

// MapOfMethodSettings represents the MapOfMethodSettings schema from the OpenAPI specification
type MapOfMethodSettings struct {
}

// MutualTlsAuthentication represents the MutualTlsAuthentication schema from the OpenAPI specification
type MutualTlsAuthentication struct {
	Truststorewarnings interface{} `json:"truststoreWarnings,omitempty"`
	Truststoreuri interface{} `json:"truststoreUri,omitempty"`
	Truststoreversion interface{} `json:"truststoreVersion,omitempty"`
}

// GetGatewayResponseRequest represents the GetGatewayResponseRequest schema from the OpenAPI specification
type GetGatewayResponseRequest struct {
}

// PutRestApiRequest represents the PutRestApiRequest schema from the OpenAPI specification
type PutRestApiRequest struct {
	Body interface{} `json:"body"`
}

// GetAuthorizerRequest represents the GetAuthorizerRequest schema from the OpenAPI specification
type GetAuthorizerRequest struct {
}

// SdkTypes represents the SdkTypes schema from the OpenAPI specification
type SdkTypes struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// GetIntegrationResponseRequest represents the GetIntegrationResponseRequest schema from the OpenAPI specification
type GetIntegrationResponseRequest struct {
}

// UpdateBasePathMappingRequest represents the UpdateBasePathMappingRequest schema from the OpenAPI specification
type UpdateBasePathMappingRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// GetUsagePlanRequest represents the GetUsagePlanRequest schema from the OpenAPI specification
type GetUsagePlanRequest struct {
}

// CreateDocumentationPartRequest represents the CreateDocumentationPartRequest schema from the OpenAPI specification
type CreateDocumentationPartRequest struct {
	Location interface{} `json:"location"`
	Properties interface{} `json:"properties"` // The new documentation content map of the targeted API entity. Enclosed key-value pairs are API-specific, but only OpenAPI-compliant key-value pairs can be exported and, hence, published.
}

// UpdateStageRequest represents the UpdateStageRequest schema from the OpenAPI specification
type UpdateStageRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// ApiStage represents the ApiStage schema from the OpenAPI specification
type ApiStage struct {
	Throttle interface{} `json:"throttle,omitempty"`
	Apiid interface{} `json:"apiId,omitempty"`
	Stage interface{} `json:"stage,omitempty"`
}

// ImportDocumentationPartsRequest represents the ImportDocumentationPartsRequest schema from the OpenAPI specification
type ImportDocumentationPartsRequest struct {
	Body interface{} `json:"body"`
}

// MapOfMethodSnapshot represents the MapOfMethodSnapshot schema from the OpenAPI specification
type MapOfMethodSnapshot struct {
}

// UpdateUsageRequest represents the UpdateUsageRequest schema from the OpenAPI specification
type UpdateUsageRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// GetClientCertificatesRequest represents the GetClientCertificatesRequest schema from the OpenAPI specification
type GetClientCertificatesRequest struct {
}

// GetUsagePlanKeyRequest represents the GetUsagePlanKeyRequest schema from the OpenAPI specification
type GetUsagePlanKeyRequest struct {
}

// GetRestApisRequest represents the GetRestApisRequest schema from the OpenAPI specification
type GetRestApisRequest struct {
}

// GetIntegrationRequest represents the GetIntegrationRequest schema from the OpenAPI specification
type GetIntegrationRequest struct {
}

// EndpointConfiguration represents the EndpointConfiguration schema from the OpenAPI specification
type EndpointConfiguration struct {
	Types interface{} `json:"types,omitempty"`
	Vpcendpointids interface{} `json:"vpcEndpointIds,omitempty"`
}

// GetTagsRequest represents the GetTagsRequest schema from the OpenAPI specification
type GetTagsRequest struct {
}

// GetGatewayResponsesRequest represents the GetGatewayResponsesRequest schema from the OpenAPI specification
type GetGatewayResponsesRequest struct {
}

// UpdateUsagePlanRequest represents the UpdateUsagePlanRequest schema from the OpenAPI specification
type UpdateUsagePlanRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// RestApi represents the RestApi schema from the OpenAPI specification
type RestApi struct {
	Name interface{} `json:"name,omitempty"`
	Apikeysource interface{} `json:"apiKeySource,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Disableexecuteapiendpoint interface{} `json:"disableExecuteApiEndpoint,omitempty"`
	Policy interface{} `json:"policy,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Version interface{} `json:"version,omitempty"`
	Warnings interface{} `json:"warnings,omitempty"`
	Binarymediatypes interface{} `json:"binaryMediaTypes,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Endpointconfiguration interface{} `json:"endpointConfiguration,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Minimumcompressionsize interface{} `json:"minimumCompressionSize,omitempty"`
}

// GetClientCertificateRequest represents the GetClientCertificateRequest schema from the OpenAPI specification
type GetClientCertificateRequest struct {
}

// UpdateResourceRequest represents the UpdateResourceRequest schema from the OpenAPI specification
type UpdateResourceRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// TestInvokeAuthorizerRequest represents the TestInvokeAuthorizerRequest schema from the OpenAPI specification
type TestInvokeAuthorizerRequest struct {
	Multivalueheaders interface{} `json:"multiValueHeaders,omitempty"`
	Pathwithquerystring interface{} `json:"pathWithQueryString,omitempty"`
	Stagevariables interface{} `json:"stageVariables,omitempty"`
	Additionalcontext interface{} `json:"additionalContext,omitempty"`
	Body interface{} `json:"body,omitempty"`
	Headers interface{} `json:"headers,omitempty"`
}

// DeleteIntegrationRequest represents the DeleteIntegrationRequest schema from the OpenAPI specification
type DeleteIntegrationRequest struct {
}

// ImportRestApiRequest represents the ImportRestApiRequest schema from the OpenAPI specification
type ImportRestApiRequest struct {
	Body interface{} `json:"body"`
}

// DomainName represents the DomainName schema from the OpenAPI specification
type DomainName struct {
	Distributiondomainname interface{} `json:"distributionDomainName,omitempty"`
	Domainnamestatusmessage interface{} `json:"domainNameStatusMessage,omitempty"`
	Certificatename interface{} `json:"certificateName,omitempty"`
	Regionalcertificatename interface{} `json:"regionalCertificateName,omitempty"`
	Securitypolicy interface{} `json:"securityPolicy,omitempty"`
	Certificateuploaddate interface{} `json:"certificateUploadDate,omitempty"`
	Mutualtlsauthentication interface{} `json:"mutualTlsAuthentication,omitempty"`
	Ownershipverificationcertificatearn interface{} `json:"ownershipVerificationCertificateArn,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Domainnamestatus interface{} `json:"domainNameStatus,omitempty"`
	Regionaldomainname interface{} `json:"regionalDomainName,omitempty"`
	Distributionhostedzoneid interface{} `json:"distributionHostedZoneId,omitempty"`
	Endpointconfiguration interface{} `json:"endpointConfiguration,omitempty"`
	Regionalcertificatearn interface{} `json:"regionalCertificateArn,omitempty"`
	Certificatearn interface{} `json:"certificateArn,omitempty"`
	Domainname interface{} `json:"domainName,omitempty"`
	Regionalhostedzoneid interface{} `json:"regionalHostedZoneId,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
}

// PutMethodResponseRequest represents the PutMethodResponseRequest schema from the OpenAPI specification
type PutMethodResponseRequest struct {
	Responseparameters interface{} `json:"responseParameters,omitempty"`
	Responsemodels interface{} `json:"responseModels,omitempty"`
}

// GetDocumentationVersionsRequest represents the GetDocumentationVersionsRequest schema from the OpenAPI specification
type GetDocumentationVersionsRequest struct {
}

// MethodResponse represents the MethodResponse schema from the OpenAPI specification
type MethodResponse struct {
	Responseparameters interface{} `json:"responseParameters,omitempty"`
	Statuscode interface{} `json:"statusCode,omitempty"`
	Responsemodels interface{} `json:"responseModels,omitempty"`
}

// UsagePlan represents the UsagePlan schema from the OpenAPI specification
type UsagePlan struct {
	Throttle interface{} `json:"throttle,omitempty"`
	Apistages interface{} `json:"apiStages,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Productcode interface{} `json:"productCode,omitempty"`
	Quota interface{} `json:"quota,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// MapOfIntegrationResponse represents the MapOfIntegrationResponse schema from the OpenAPI specification
type MapOfIntegrationResponse struct {
}

// UpdateGatewayResponseRequest represents the UpdateGatewayResponseRequest schema from the OpenAPI specification
type UpdateGatewayResponseRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// CanarySettings represents the CanarySettings schema from the OpenAPI specification
type CanarySettings struct {
	Deploymentid interface{} `json:"deploymentId,omitempty"`
	Percenttraffic interface{} `json:"percentTraffic,omitempty"`
	Stagevariableoverrides interface{} `json:"stageVariableOverrides,omitempty"`
	Usestagecache interface{} `json:"useStageCache,omitempty"`
}

// GetDomainNameRequest represents the GetDomainNameRequest schema from the OpenAPI specification
type GetDomainNameRequest struct {
}

// GetVpcLinkRequest represents the GetVpcLinkRequest schema from the OpenAPI specification
type GetVpcLinkRequest struct {
}

// CreateResourceRequest represents the CreateResourceRequest schema from the OpenAPI specification
type CreateResourceRequest struct {
	Pathpart interface{} `json:"pathPart"`
}

// MapOfKeyUsages represents the MapOfKeyUsages schema from the OpenAPI specification
type MapOfKeyUsages struct {
}

// VpcLink represents the VpcLink schema from the OpenAPI specification
type VpcLink struct {
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Targetarns interface{} `json:"targetArns,omitempty"`
}

// GetModelsRequest represents the GetModelsRequest schema from the OpenAPI specification
type GetModelsRequest struct {
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	Timeoutinmillis interface{} `json:"timeoutInMillis,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Cachekeyparameters interface{} `json:"cacheKeyParameters,omitempty"`
	Connectionid interface{} `json:"connectionId,omitempty"`
	Httpmethod interface{} `json:"httpMethod,omitempty"`
	Passthroughbehavior interface{} `json:"passthroughBehavior,omitempty"`
	Requesttemplates interface{} `json:"requestTemplates,omitempty"`
	Tlsconfig interface{} `json:"tlsConfig,omitempty"`
	Uri interface{} `json:"uri,omitempty"`
	Credentials interface{} `json:"credentials,omitempty"`
	Integrationresponses interface{} `json:"integrationResponses,omitempty"`
	Connectiontype interface{} `json:"connectionType,omitempty"`
	Contenthandling interface{} `json:"contentHandling,omitempty"`
	Cachenamespace interface{} `json:"cacheNamespace,omitempty"`
	Requestparameters interface{} `json:"requestParameters,omitempty"`
}

// ThrottleSettings represents the ThrottleSettings schema from the OpenAPI specification
type ThrottleSettings struct {
	Ratelimit interface{} `json:"rateLimit,omitempty"`
	Burstlimit interface{} `json:"burstLimit,omitempty"`
}

// PutIntegrationRequest represents the PutIntegrationRequest schema from the OpenAPI specification
type PutIntegrationRequest struct {
	Cachenamespace interface{} `json:"cacheNamespace,omitempty"`
	Integrationhttpmethod interface{} `json:"integrationHttpMethod,omitempty"`
	Requestparameters interface{} `json:"requestParameters,omitempty"`
	TypeField interface{} `json:"type"`
	Cachekeyparameters interface{} `json:"cacheKeyParameters,omitempty"`
	Contenthandling interface{} `json:"contentHandling,omitempty"`
	Credentials interface{} `json:"credentials,omitempty"`
	Requesttemplates interface{} `json:"requestTemplates,omitempty"`
	Tlsconfig TlsConfig `json:"tlsConfig,omitempty"` // Specifies the TLS configuration for an integration.
	Connectiontype interface{} `json:"connectionType,omitempty"`
	Timeoutinmillis interface{} `json:"timeoutInMillis,omitempty"`
	Connectionid interface{} `json:"connectionId,omitempty"`
	Passthroughbehavior interface{} `json:"passthroughBehavior,omitempty"`
	Uri interface{} `json:"uri,omitempty"`
}

// UpdateMethodRequest represents the UpdateMethodRequest schema from the OpenAPI specification
type UpdateMethodRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// DeleteMethodRequest represents the DeleteMethodRequest schema from the OpenAPI specification
type DeleteMethodRequest struct {
}

// AccessLogSettings represents the AccessLogSettings schema from the OpenAPI specification
type AccessLogSettings struct {
	Format interface{} `json:"format,omitempty"`
	Destinationarn interface{} `json:"destinationArn,omitempty"`
}

// GetDeploymentRequest represents the GetDeploymentRequest schema from the OpenAPI specification
type GetDeploymentRequest struct {
}

// UpdateAuthorizerRequest represents the UpdateAuthorizerRequest schema from the OpenAPI specification
type UpdateAuthorizerRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// MapOfStringToBoolean represents the MapOfStringToBoolean schema from the OpenAPI specification
type MapOfStringToBoolean struct {
}

// DeploymentCanarySettings represents the DeploymentCanarySettings schema from the OpenAPI specification
type DeploymentCanarySettings struct {
	Percenttraffic interface{} `json:"percentTraffic,omitempty"`
	Stagevariableoverrides interface{} `json:"stageVariableOverrides,omitempty"`
	Usestagecache interface{} `json:"useStageCache,omitempty"`
}

// DeleteAuthorizerRequest represents the DeleteAuthorizerRequest schema from the OpenAPI specification
type DeleteAuthorizerRequest struct {
}

// DeleteDomainNameRequest represents the DeleteDomainNameRequest schema from the OpenAPI specification
type DeleteDomainNameRequest struct {
}

// CreateAuthorizerRequest represents the CreateAuthorizerRequest schema from the OpenAPI specification
type CreateAuthorizerRequest struct {
	Authtype interface{} `json:"authType,omitempty"`
	Identitysource interface{} `json:"identitySource,omitempty"`
	Identityvalidationexpression interface{} `json:"identityValidationExpression,omitempty"`
	Providerarns interface{} `json:"providerARNs,omitempty"`
	Authorizercredentials interface{} `json:"authorizerCredentials,omitempty"`
	TypeField interface{} `json:"type"`
	Authorizerresultttlinseconds interface{} `json:"authorizerResultTtlInSeconds,omitempty"`
	Name interface{} `json:"name"`
	Authorizeruri interface{} `json:"authorizerUri,omitempty"`
}

// DeleteRequestValidatorRequest represents the DeleteRequestValidatorRequest schema from the OpenAPI specification
type DeleteRequestValidatorRequest struct {
}

// DocumentationPart represents the DocumentationPart schema from the OpenAPI specification
type DocumentationPart struct {
	Location interface{} `json:"location,omitempty"`
	Properties interface{} `json:"properties,omitempty"` // A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., <code>"{ \"description\": \"The API does ...\" }"</code>. Only OpenAPI-compliant documentation-related fields from the properties map are exported and, hence, published as part of the API entity definitions, while the original documentation parts are exported in a OpenAPI extension of <code>x-amazon-apigateway-documentation</code>.
	Id interface{} `json:"id,omitempty"`
}

// UpdateDeploymentRequest represents the UpdateDeploymentRequest schema from the OpenAPI specification
type UpdateDeploymentRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// DeleteRestApiRequest represents the DeleteRestApiRequest schema from the OpenAPI specification
type DeleteRestApiRequest struct {
}

// UsagePlanKey represents the UsagePlanKey schema from the OpenAPI specification
type UsagePlanKey struct {
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// MethodSnapshot represents the MethodSnapshot schema from the OpenAPI specification
type MethodSnapshot struct {
	Apikeyrequired interface{} `json:"apiKeyRequired,omitempty"`
	Authorizationtype interface{} `json:"authorizationType,omitempty"`
}

// RestApis represents the RestApis schema from the OpenAPI specification
type RestApis struct {
	Position string `json:"position,omitempty"`
	Items interface{} `json:"items,omitempty"`
}

// Resources represents the Resources schema from the OpenAPI specification
type Resources struct {
	Position string `json:"position,omitempty"`
	Items interface{} `json:"items,omitempty"`
}

// ExportResponse represents the ExportResponse schema from the OpenAPI specification
type ExportResponse struct {
	Body interface{} `json:"body,omitempty"`
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Apisummary interface{} `json:"apiSummary,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
}

// DeleteUsagePlanRequest represents the DeleteUsagePlanRequest schema from the OpenAPI specification
type DeleteUsagePlanRequest struct {
}

// GetDocumentationVersionRequest represents the GetDocumentationVersionRequest schema from the OpenAPI specification
type GetDocumentationVersionRequest struct {
}

// Method represents the Method schema from the OpenAPI specification
type Method struct {
	Authorizerid interface{} `json:"authorizerId,omitempty"`
	Httpmethod interface{} `json:"httpMethod,omitempty"`
	Operationname interface{} `json:"operationName,omitempty"`
	Requestparameters interface{} `json:"requestParameters,omitempty"`
	Authorizationtype interface{} `json:"authorizationType,omitempty"`
	Authorizationscopes interface{} `json:"authorizationScopes,omitempty"`
	Requestmodels interface{} `json:"requestModels,omitempty"`
	Requestvalidatorid interface{} `json:"requestValidatorId,omitempty"`
	Apikeyrequired interface{} `json:"apiKeyRequired,omitempty"`
	Methodintegration interface{} `json:"methodIntegration,omitempty"`
	Methodresponses interface{} `json:"methodResponses,omitempty"`
}

// GatewayResponses represents the GatewayResponses schema from the OpenAPI specification
type GatewayResponses struct {
	Position string `json:"position,omitempty"`
	Items interface{} `json:"items,omitempty"`
}

// ApiKeys represents the ApiKeys schema from the OpenAPI specification
type ApiKeys struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
	Warnings interface{} `json:"warnings,omitempty"`
}

// GetAccountRequest represents the GetAccountRequest schema from the OpenAPI specification
type GetAccountRequest struct {
}

// CreateBasePathMappingRequest represents the CreateBasePathMappingRequest schema from the OpenAPI specification
type CreateBasePathMappingRequest struct {
	Basepath interface{} `json:"basePath,omitempty"`
	Restapiid interface{} `json:"restApiId"`
	Stage interface{} `json:"stage,omitempty"`
}

// CreateUsagePlanKeyRequest represents the CreateUsagePlanKeyRequest schema from the OpenAPI specification
type CreateUsagePlanKeyRequest struct {
	Keytype interface{} `json:"keyType"`
	Keyid interface{} `json:"keyId"`
}

// MapOfApiStageThrottleSettings represents the MapOfApiStageThrottleSettings schema from the OpenAPI specification
type MapOfApiStageThrottleSettings struct {
}

// MapOfStringToList represents the MapOfStringToList schema from the OpenAPI specification
type MapOfStringToList struct {
}

// Stage represents the Stage schema from the OpenAPI specification
type Stage struct {
	Cacheclustersize interface{} `json:"cacheClusterSize,omitempty"`
	Cacheclusterenabled interface{} `json:"cacheClusterEnabled,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Accesslogsettings interface{} `json:"accessLogSettings,omitempty"`
	Tracingenabled interface{} `json:"tracingEnabled,omitempty"`
	Variables interface{} `json:"variables,omitempty"`
	Webaclarn interface{} `json:"webAclArn,omitempty"`
	Createddate interface{} `json:"createdDate,omitempty"`
	Cacheclusterstatus interface{} `json:"cacheClusterStatus,omitempty"`
	Canarysettings interface{} `json:"canarySettings,omitempty"`
	Documentationversion interface{} `json:"documentationVersion,omitempty"`
	Lastupdateddate interface{} `json:"lastUpdatedDate,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Methodsettings interface{} `json:"methodSettings,omitempty"`
	Stagename interface{} `json:"stageName,omitempty"`
	Clientcertificateid interface{} `json:"clientCertificateId,omitempty"`
	Deploymentid interface{} `json:"deploymentId,omitempty"`
}

// DeleteGatewayResponseRequest represents the DeleteGatewayResponseRequest schema from the OpenAPI specification
type DeleteGatewayResponseRequest struct {
}

// GetStageRequest represents the GetStageRequest schema from the OpenAPI specification
type GetStageRequest struct {
}

// GetUsagePlansRequest represents the GetUsagePlansRequest schema from the OpenAPI specification
type GetUsagePlansRequest struct {
}

// PatchOperation represents the PatchOperation schema from the OpenAPI specification
type PatchOperation struct {
	From interface{} `json:"from,omitempty"`
	Op interface{} `json:"op,omitempty"`
	Path interface{} `json:"path,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// CreateDocumentationVersionRequest represents the CreateDocumentationVersionRequest schema from the OpenAPI specification
type CreateDocumentationVersionRequest struct {
	Description interface{} `json:"description,omitempty"`
	Documentationversion interface{} `json:"documentationVersion"`
	Stagename interface{} `json:"stageName,omitempty"`
}

// GetMethodResponseRequest represents the GetMethodResponseRequest schema from the OpenAPI specification
type GetMethodResponseRequest struct {
}

// Deployments represents the Deployments schema from the OpenAPI specification
type Deployments struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// SdkResponse represents the SdkResponse schema from the OpenAPI specification
type SdkResponse struct {
	Body interface{} `json:"body,omitempty"`
}

// FlushStageAuthorizersCacheRequest represents the FlushStageAuthorizersCacheRequest schema from the OpenAPI specification
type FlushStageAuthorizersCacheRequest struct {
}

// IntegrationResponse represents the IntegrationResponse schema from the OpenAPI specification
type IntegrationResponse struct {
	Responsetemplates interface{} `json:"responseTemplates,omitempty"`
	Selectionpattern interface{} `json:"selectionPattern,omitempty"`
	Statuscode interface{} `json:"statusCode,omitempty"`
	Contenthandling interface{} `json:"contentHandling,omitempty"`
	Responseparameters interface{} `json:"responseParameters,omitempty"`
}

// GetMethodRequest represents the GetMethodRequest schema from the OpenAPI specification
type GetMethodRequest struct {
}

// DeleteUsagePlanKeyRequest represents the DeleteUsagePlanKeyRequest schema from the OpenAPI specification
type DeleteUsagePlanKeyRequest struct {
}

// BasePathMapping represents the BasePathMapping schema from the OpenAPI specification
type BasePathMapping struct {
	Basepath interface{} `json:"basePath,omitempty"`
	Restapiid interface{} `json:"restApiId,omitempty"`
	Stage interface{} `json:"stage,omitempty"`
}

// SdkConfigurationProperty represents the SdkConfigurationProperty schema from the OpenAPI specification
type SdkConfigurationProperty struct {
	Description interface{} `json:"description,omitempty"`
	Friendlyname interface{} `json:"friendlyName,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Required interface{} `json:"required,omitempty"`
	Defaultvalue interface{} `json:"defaultValue,omitempty"`
}

// DeleteDeploymentRequest represents the DeleteDeploymentRequest schema from the OpenAPI specification
type DeleteDeploymentRequest struct {
}

// Usage represents the Usage schema from the OpenAPI specification
type Usage struct {
	Enddate interface{} `json:"endDate,omitempty"`
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
	Startdate interface{} `json:"startDate,omitempty"`
	Usageplanid interface{} `json:"usagePlanId,omitempty"`
}

// GatewayResponse represents the GatewayResponse schema from the OpenAPI specification
type GatewayResponse struct {
	Responsetemplates interface{} `json:"responseTemplates,omitempty"`
	Responsetype interface{} `json:"responseType,omitempty"`
	Statuscode interface{} `json:"statusCode,omitempty"`
	Defaultresponse interface{} `json:"defaultResponse,omitempty"`
	Responseparameters interface{} `json:"responseParameters,omitempty"`
}

// UpdateMethodResponseRequest represents the UpdateMethodResponseRequest schema from the OpenAPI specification
type UpdateMethodResponseRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// DeleteMethodResponseRequest represents the DeleteMethodResponseRequest schema from the OpenAPI specification
type DeleteMethodResponseRequest struct {
}

// PutMethodRequest represents the PutMethodRequest schema from the OpenAPI specification
type PutMethodRequest struct {
	Requestvalidatorid interface{} `json:"requestValidatorId,omitempty"`
	Apikeyrequired interface{} `json:"apiKeyRequired,omitempty"`
	Authorizationscopes interface{} `json:"authorizationScopes,omitempty"`
	Authorizationtype interface{} `json:"authorizationType"`
	Authorizerid interface{} `json:"authorizerId,omitempty"`
	Operationname interface{} `json:"operationName,omitempty"`
	Requestmodels interface{} `json:"requestModels,omitempty"`
	Requestparameters interface{} `json:"requestParameters,omitempty"`
}

// TestInvokeAuthorizerResponse represents the TestInvokeAuthorizerResponse schema from the OpenAPI specification
type TestInvokeAuthorizerResponse struct {
	Clientstatus interface{} `json:"clientStatus,omitempty"`
	Latency interface{} `json:"latency,omitempty"`
	Log interface{} `json:"log,omitempty"`
	Policy interface{} `json:"policy,omitempty"`
	Principalid interface{} `json:"principalId,omitempty"`
	Authorization interface{} `json:"authorization,omitempty"`
	Claims interface{} `json:"claims,omitempty"`
}

// Tags represents the Tags schema from the OpenAPI specification
type Tags struct {
	Tags interface{} `json:"tags,omitempty"`
}

// GetAuthorizersRequest represents the GetAuthorizersRequest schema from the OpenAPI specification
type GetAuthorizersRequest struct {
}

// MapOfMethodResponse represents the MapOfMethodResponse schema from the OpenAPI specification
type MapOfMethodResponse struct {
}

// DeleteClientCertificateRequest represents the DeleteClientCertificateRequest schema from the OpenAPI specification
type DeleteClientCertificateRequest struct {
}

// TestInvokeMethodResponse represents the TestInvokeMethodResponse schema from the OpenAPI specification
type TestInvokeMethodResponse struct {
	Multivalueheaders interface{} `json:"multiValueHeaders,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Body interface{} `json:"body,omitempty"`
	Headers interface{} `json:"headers,omitempty"`
	Latency interface{} `json:"latency,omitempty"`
	Log interface{} `json:"log,omitempty"`
}

// DocumentationPartLocation represents the DocumentationPartLocation schema from the OpenAPI specification
type DocumentationPartLocation struct {
	Method interface{} `json:"method,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Path interface{} `json:"path,omitempty"`
	Statuscode interface{} `json:"statusCode,omitempty"`
	TypeField interface{} `json:"type"`
}

// RequestValidator represents the RequestValidator schema from the OpenAPI specification
type RequestValidator struct {
	Validaterequestparameters interface{} `json:"validateRequestParameters,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Validaterequestbody interface{} `json:"validateRequestBody,omitempty"`
}

// UsagePlanKeys represents the UsagePlanKeys schema from the OpenAPI specification
type UsagePlanKeys struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// CreateStageRequest represents the CreateStageRequest schema from the OpenAPI specification
type CreateStageRequest struct {
	Stagename interface{} `json:"stageName"`
	Deploymentid interface{} `json:"deploymentId"`
	Description interface{} `json:"description,omitempty"`
	Tracingenabled interface{} `json:"tracingEnabled,omitempty"`
	Variables interface{} `json:"variables,omitempty"`
	Documentationversion interface{} `json:"documentationVersion,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Cacheclustersize interface{} `json:"cacheClusterSize,omitempty"`
	Canarysettings interface{} `json:"canarySettings,omitempty"`
	Cacheclusterenabled interface{} `json:"cacheClusterEnabled,omitempty"`
}

// DocumentationParts represents the DocumentationParts schema from the OpenAPI specification
type DocumentationParts struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}

// UpdateDocumentationPartRequest represents the UpdateDocumentationPartRequest schema from the OpenAPI specification
type UpdateDocumentationPartRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// UpdateVpcLinkRequest represents the UpdateVpcLinkRequest schema from the OpenAPI specification
type UpdateVpcLinkRequest struct {
	Patchoperations interface{} `json:"patchOperations,omitempty"`
}

// GetRequestValidatorsRequest represents the GetRequestValidatorsRequest schema from the OpenAPI specification
type GetRequestValidatorsRequest struct {
}

// RequestValidators represents the RequestValidators schema from the OpenAPI specification
type RequestValidators struct {
	Items interface{} `json:"items,omitempty"`
	Position string `json:"position,omitempty"`
}
