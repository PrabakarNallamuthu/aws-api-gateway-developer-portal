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

func DeleteusageplankeyHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		keyIdVal, ok := args["keyId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: keyId"), nil
		}
		keyId, ok := keyIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: keyId"), nil
		}
		url := fmt.Sprintf("%s/usageplans/%s/keys/%s", cfg.BaseURL, usageplanId, keyId)
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

func CreateDeleteusageplankeyTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_usageplans_usageplanId_keys_keyId",
		mcp.WithDescription("Deletes a usage plan key and remove the underlying API key from the associated usage plan."),
		mcp.WithString("usageplanId", mcp.Required(), mcp.Description("The Id of the UsagePlan resource representing the usage plan containing the to-be-deleted UsagePlanKey resource representing a plan customer.")),
		mcp.WithString("keyId", mcp.Required(), mcp.Description("The Id of the UsagePlanKey resource to be deleted.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteusageplankeyHandler(cfg),
	}
}
