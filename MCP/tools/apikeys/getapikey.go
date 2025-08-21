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

func GetapikeyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		api_KeyVal, ok := args["api_Key"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: api_Key"), nil
		}
		api_Key, ok := api_KeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: api_Key"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["includeValue"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeValue=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/apikeys/%s%s", cfg.BaseURL, api_Key, queryString)
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
		var result models.ApiKey
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

func CreateGetapikeyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_apikeys_api_Key",
		mcp.WithDescription("Gets information about the current ApiKey resource."),
		mcp.WithString("api_Key", mcp.Required(), mcp.Description("The identifier of the ApiKey resource.")),
		mcp.WithBoolean("includeValue", mcp.Description("A boolean flag to specify whether (<code>true</code>) or not (<code>false</code>) the result contains the key value.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetapikeyHandler(cfg),
	}
}
