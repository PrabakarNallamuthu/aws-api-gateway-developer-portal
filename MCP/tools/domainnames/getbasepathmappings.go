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

func GetbasepathmappingsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["position"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("position=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/domainnames/%s/basepathmappings%s", cfg.BaseURL, domain_name, queryString)
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
		var result models.BasePathMappings
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

func CreateGetbasepathmappingsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_domainnames_domain_name_basepathmappings",
		mcp.WithDescription("Represents a collection of BasePathMapping resources."),
		mcp.WithString("domain_name", mcp.Required(), mcp.Description("The domain name of a BasePathMapping resource.")),
		mcp.WithString("position", mcp.Description("The current pagination position in the paged result set.")),
		mcp.WithNumber("limit", mcp.Description("The maximum number of returned results per page. The default value is 25 and the maximum value is 500.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetbasepathmappingsHandler(cfg),
	}
}
