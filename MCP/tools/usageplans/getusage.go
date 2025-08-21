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

func GetusageHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		usageplanIdVal, ok := args["usageplanId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: usageplanId"), nil
		}
		usageplanId, ok := usageplanIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: usageplanId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["keyId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("keyId=%v", val))
		}
		if val, ok := args["startDate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startDate=%v", val))
		}
		if val, ok := args["endDate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("endDate=%v", val))
		}
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
		url := fmt.Sprintf("%s/usageplans/%s/usage#startDate&endDate%s", cfg.BaseURL, usageplanId, queryString)
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
		var result models.Usage
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

func CreateGetusageTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_usageplans_usageplanId_usage#startDate&endDate",
		mcp.WithDescription("Gets the usage data of a usage plan in a specified time interval."),
		mcp.WithString("usageplanId", mcp.Required(), mcp.Description("The Id of the usage plan associated with the usage data.")),
		mcp.WithString("keyId", mcp.Description("The Id of the API key associated with the resultant usage data.")),
		mcp.WithString("startDate", mcp.Required(), mcp.Description("The starting date (e.g., 2016-01-01) of the usage data.")),
		mcp.WithString("endDate", mcp.Required(), mcp.Description("The ending date (e.g., 2016-12-31) of the usage data.")),
		mcp.WithString("position", mcp.Description("The current pagination position in the paged result set.")),
		mcp.WithNumber("limit", mcp.Description("The maximum number of returned results per page. The default value is 25 and the maximum value is 500.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetusageHandler(cfg),
	}
}
