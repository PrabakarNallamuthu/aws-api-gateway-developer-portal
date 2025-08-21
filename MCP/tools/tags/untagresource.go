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

func UntagresourceHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		resource_arnVal, ok := args["resource_arn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resource_arn"), nil
		}
		resource_arn, ok := resource_arnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resource_arn"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["tagKeys"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tagKeys=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/tags/%s#tagKeys%s", cfg.BaseURL, resource_arn, queryString)
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

func CreateUntagresourceTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_tags_resource_arn#tagKeys",
		mcp.WithDescription("Removes a tag from a given resource."),
		mcp.WithString("resource_arn", mcp.Required(), mcp.Description("The ARN of a resource that can be tagged.")),
		mcp.WithArray("tagKeys", mcp.Required(), mcp.Description("The Tag keys to delete.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UntagresourceHandler(cfg),
	}
}
