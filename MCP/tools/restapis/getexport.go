package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-api-gateway/mcp-server/config"
	"github.com/amazon-api-gateway/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetexportHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		stage_nameVal, ok := args["stage_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: stage_name"), nil
		}
		stage_name, ok := stage_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: stage_name"), nil
		}
		export_typeVal, ok := args["export_type"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: export_type"), nil
		}
		export_type, ok := export_typeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: export_type"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["parameters"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("parameters=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/restapis/%s/stages/%s/exports/%s%s", cfg.BaseURL, restapi_id, stage_name, export_type, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["Accept"]; ok {
			req.Header.Set("Accept", fmt.Sprintf("%v", val))
		}

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
		var result models.ExportResponse
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

func CreateGetexportTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_restapis_restapi_id_stages_stage_name_exports_export_type",
		mcp.WithDescription("Exports a deployed version of a RestApi in a specified format."),
		mcp.WithString("restapi_id", mcp.Required(), mcp.Description("The string identifier of the associated RestApi.")),
		mcp.WithString("stage_name", mcp.Required(), mcp.Description("The name of the Stage that will be exported.")),
		mcp.WithString("export_type", mcp.Required(), mcp.Description("The type of export. Acceptable values are 'oas30' for OpenAPI 3.0.x and 'swagger' for Swagger/OpenAPI 2.0.")),
		mcp.WithObject("parameters", mcp.Description("A key-value map of query string parameters that specify properties of the export, depending on the requested <code>exportType</code>. For <code>exportType</code> <code>oas30</code> and <code>swagger</code>, any combination of the following parameters are supported: <code>extensions='integrations'</code> or <code>extensions='apigateway'</code> will export the API with x-amazon-apigateway-integration extensions. <code>extensions='authorizers'</code> will export the API with x-amazon-apigateway-authorizer extensions. <code>postman</code> will export the API with Postman extensions, allowing for import to the Postman tool")),
		mcp.WithString("Accept", mcp.Description("The content-type of the export, for example <code>application/json</code>. Currently <code>application/json</code> and <code>application/yaml</code> are supported for <code>exportType</code> of<code>oas30</code> and <code>swagger</code>. This should be specified in the <code>Accept</code> header for direct API requests.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetexportHandler(cfg),
	}
}
