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

func TestinvokemethodHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/restapis/%s/resources/%s/methods/%s", cfg.BaseURL, restapi_id, resource_id, http_method)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
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
		var result models.TestInvokeMethodResponse
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

func CreateTestinvokemethodTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_restapis_restapi_id_resources_resource_id_methods_http_method",
		mcp.WithDescription("Simulate the invocation of a Method in your RestApi with headers, parameters, and an incoming request body."),
		mcp.WithString("restapi_id", mcp.Required(), mcp.Description("The string identifier of the associated RestApi.")),
		mcp.WithString("resource_id", mcp.Required(), mcp.Description("Specifies a test invoke method request's resource ID.")),
		mcp.WithString("http_method", mcp.Required(), mcp.Description("Specifies a test invoke method request's HTTP method.")),
		mcp.WithString("pathWithQueryString", mcp.Description("Input parameter: The URI path, including query string, of the simulated invocation request. Use this to specify path parameters and query string parameters.")),
		mcp.WithObject("stageVariables", mcp.Description("Input parameter: A key-value map of stage variables to simulate an invocation on a deployed Stage.")),
		mcp.WithString("body", mcp.Description("Input parameter: The simulated request body of an incoming invocation request.")),
		mcp.WithString("clientCertificateId", mcp.Description("Input parameter: A ClientCertificate identifier to use in the test invocation. API Gateway will use the certificate when making the HTTPS request to the defined back-end endpoint.")),
		mcp.WithObject("headers", mcp.Description("Input parameter: A key-value map of headers to simulate an incoming invocation request.")),
		mcp.WithObject("multiValueHeaders", mcp.Description("Input parameter: The headers as a map from string to list of values to simulate an incoming invocation request.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    TestinvokemethodHandler(cfg),
	}
}
