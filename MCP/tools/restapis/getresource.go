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

func GetresourceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["embed"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("embed=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/restapis/%s/resources/%s%s", cfg.BaseURL, restapi_id, resource_id, queryString)
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
		var result models.Resource
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

func CreateGetresourceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_restapis_restapi_id_resources_resource_id",
		mcp.WithDescription("Lists information about a resource."),
		mcp.WithString("restapi_id", mcp.Required(), mcp.Description("The string identifier of the associated RestApi.")),
		mcp.WithString("resource_id", mcp.Required(), mcp.Description("The identifier for the Resource resource.")),
		mcp.WithArray("embed", mcp.Description("A query parameter to retrieve the specified resources embedded in the returned Resource representation in the response. This <code>embed</code> parameter value is a list of comma-separated strings. Currently, the request supports only retrieval of the embedded Method resources this way. The query parameter value must be a single-valued list and contain the <code>\"methods\"</code> string. For example, <code>GET /restapis/{restapi_id}/resources/{resource_id}?embed=methods</code>.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetresourceHandler(cfg),
	}
}
