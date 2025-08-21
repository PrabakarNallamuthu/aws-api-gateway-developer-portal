package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amazon-api-gateway/mcp-server/config"
	"github.com/amazon-api-gateway/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DeletebasepathmappingHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		domain_nameVal, ok := args["domain_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: domain_name"), nil
		}
		domain_name, ok := domain_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: domain_name"), nil
		}
		base_pathVal, ok := args["base_path"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: base_path"), nil
		}
		base_path, ok := base_pathVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: base_path"), nil
		}
		url := fmt.Sprintf("%s/domainnames/%s/basepathmappings/%s", cfg.BaseURL, domain_name, base_path)
		req, err := http.NewRequest("DELETE", url, nil)
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
		var result map[string]interface{}
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

func CreateDeletebasepathmappingTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_domainnames_domain_name_basepathmappings_base_path",
		mcp.WithDescription("Deletes the BasePathMapping resource."),
		mcp.WithString("domain_name", mcp.Required(), mcp.Description("The domain name of the BasePathMapping resource to delete.")),
		mcp.WithString("base_path", mcp.Required(), mcp.Description("<p>The base path name of the BasePathMapping resource to delete.</p> <p>To specify an empty base path, set this parameter to <code>'(none)'</code>.</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletebasepathmappingHandler(cfg),
	}
}
