package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-api-gateway/mcp-server/config"
	"github.com/amazon-api-gateway/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func PutintegrationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		restapi_idVal, ok := args["restapi_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: restapi_id"), nil
		}
		restapi_id, ok := restapi_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: restapi_id"), nil
		}
		resource_idVal, ok := args["resource_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resource_id"), nil
		}
		resource_id, ok := resource_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resource_id"), nil
		}
		http_methodVal, ok := args["http_method"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: http_method"), nil
		}
		http_method, ok := http_methodVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: http_method"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/restapis/%s/resources/%s/methods/%s/integration", cfg.BaseURL, restapi_id, resource_id, http_method)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Integration
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreatePutintegrationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_restapis_restapi_id_resources_resource_id_methods_http_method_integration",
		mcp.WithDescription("Sets up a method's integration."),
		mcp.WithString("restapi_id", mcp.Required(), mcp.Description("The string identifier of the associated RestApi.")),
		mcp.WithString("resource_id", mcp.Required(), mcp.Description("Specifies a put integration request's resource ID.")),
		mcp.WithString("http_method", mcp.Required(), mcp.Description("Specifies the HTTP method for the integration.")),
		mcp.WithNumber("timeoutInMillis", mcp.Description("Input parameter: Custom timeout between 50 and 29,000 milliseconds. The default value is 29,000 milliseconds or 29 seconds.")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Input parameter: The integration type. The valid value is <code>HTTP</code> for integrating an API method with an HTTP backend; <code>AWS</code> with any AWS service endpoints; <code>MOCK</code> for testing without actually invoking the backend; <code>HTTP_PROXY</code> for integrating with the HTTP proxy integration; <code>AWS_PROXY</code> for integrating with the Lambda proxy integration. ")),
		mcp.WithString("connectionId", mcp.Description("Input parameter: The ID of the VpcLink used for the integration. Specify this value only if you specify <code>VPC_LINK</code> as the connection type.")),
		mcp.WithObject("tlsConfig", mcp.Description("Input parameter: Specifies the TLS configuration for an integration.")),
		mcp.WithString("uri", mcp.Description("Input parameter: Specifies Uniform Resource Identifier (URI) of the integration endpoint. For HTTP or <code>HTTP_PROXY</code> integrations, the URI must be a fully formed, encoded HTTP(S) URL according to the RFC-3986 specification, for either standard integration, where <code>connectionType</code> is not <code>VPC_LINK</code>, or private integration, where <code>connectionType</code> is <code>VPC_LINK</code>. For a private HTTP integration, the URI is not used for routing. For <code>AWS</code> or <code>AWS_PROXY</code> integrations, the URI is of the form <code>arn:aws:apigateway:{region}:{subdomain.service|service}:path|action/{service_api</code>}. Here, {Region} is the API Gateway region (e.g., us-east-1); {service} is the name of the integrated Amazon Web Services service (e.g., s3); and {subdomain} is a designated subdomain supported by certain Amazon Web Services service for fast host-name lookup. action can be used for an Amazon Web Services service action-based API, using an Action={name}&amp;{p1}={v1}&amp;p2={v2}... query string. The ensuing {service_api} refers to a supported action {name} plus any required input parameters. Alternatively, path can be used for an Amazon Web Services service path-based API. The ensuing service_api refers to the path to an Amazon Web Services service resource, including the region of the integrated Amazon Web Services service, if applicable. For example, for integration with the S3 API of <code>GetObject</code>, the <code>uri</code> can be either <code>arn:aws:apigateway:us-west-2:s3:action/GetObject&amp;Bucket={bucket}&amp;Key={key}</code> or <code>arn:aws:apigateway:us-west-2:s3:path/{bucket}/{key}</code>.")),
		mcp.WithString("passthroughBehavior", mcp.Description("Input parameter: Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the <code>requestTemplates</code> property on the Integration resource. There are three valid values: <code>WHEN_NO_MATCH</code>, <code>WHEN_NO_TEMPLATES</code>, and <code>NEVER</code>. ")),
		mcp.WithObject("requestTemplates", mcp.Description("Input parameter: Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. The content type value is the key in this map, and the template (as a String) is the value.")),
		mcp.WithArray("cacheKeyParameters", mcp.Description("Input parameter: A list of request parameters whose values API Gateway caches. To be valid values for <code>cacheKeyParameters</code>, these parameters must also be specified for Method <code>requestParameters</code>.")),
		mcp.WithString("connectionType", mcp.Description("Input parameter: The type of the network connection to the integration endpoint. The valid value is <code>INTERNET</code> for connections through the public routable internet or <code>VPC_LINK</code> for private connections between API Gateway and a network load balancer in a VPC. The default value is <code>INTERNET</code>.")),
		mcp.WithObject("requestParameters", mcp.Description("Input parameter: A key-value map specifying request parameters that are passed from the method request to the back end. The key is an integration request parameter name and the associated value is a method request parameter value or static value that must be enclosed within single quotes and pre-encoded as required by the back end. The method request parameter value must match the pattern of <code>method.request.{location}.{name}</code>, where <code>location</code> is <code>querystring</code>, <code>path</code>, or <code>header</code> and <code>name</code> must be a valid and unique method request parameter name.")),
		mcp.WithString("cacheNamespace", mcp.Description("Input parameter: Specifies a group of related cached parameters. By default, API Gateway uses the resource ID as the <code>cacheNamespace</code>. You can specify the same <code>cacheNamespace</code> across resources to return the same cached data for requests to different resources.")),
		mcp.WithString("contentHandling", mcp.Description("Input parameter: <p>Specifies how to handle request payload content type conversions. Supported values are <code>CONVERT_TO_BINARY</code> and <code>CONVERT_TO_TEXT</code>, with the following behaviors:</p> <p>If this property is not defined, the request payload will be passed through from the method request to integration request without modification, provided that the <code>passthroughBehavior</code> is configured to support payload pass-through.</p>")),
		mcp.WithString("credentials", mcp.Description("Input parameter: Specifies whether credentials are required for a put integration.")),
		mcp.WithString("httpMethod", mcp.Description("Input parameter: The HTTP method for the integration.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    PutintegrationHandler(cfg),
	}
}
