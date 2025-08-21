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

func GetapikeysHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["position"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("position=%v", val))
		}
		if val, ok := args["limit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("limit=%v", val))
		}
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["customerId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("customerId=%v", val))
		}
		if val, ok := args["includeValues"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("includeValues=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/apikeys%s", cfg.BaseURL, queryString)
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
		var result models.ApiKeys
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

func CreateGetapikeysTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_apikeys",
		mcp.WithDescription("Gets information about the current ApiKeys resource."),
		mcp.WithString("position", mcp.Description("The current pagination position in the paged result set.")),
		mcp.WithNumber("limit", mcp.Description("The maximum number of returned results per page. The default value is 25 and the maximum value is 500.")),
		mcp.WithString("name", mcp.Description("The name of queried API keys.")),
		mcp.WithString("customerId", mcp.Description("The identifier of a customer in AWS Marketplace or an external system, such as a developer portal.")),
		mcp.WithBoolean("includeValues", mcp.Description("A boolean flag to specify whether (<code>true</code>) or not (<code>false</code>) the result contains key values.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetapikeysHandler(cfg),
	}
}
