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

func DeletegatewayresponseHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		response_typeVal, ok := args["response_type"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: response_type"), nil
		}
		response_type, ok := response_typeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: response_type"), nil
		}
		url := fmt.Sprintf("%s/restapis/%s/gatewayresponses/%s", cfg.BaseURL, restapi_id, response_type)
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

func CreateDeletegatewayresponseTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_restapis_restapi_id_gatewayresponses_response_type",
		mcp.WithDescription("Clears any customization of a GatewayResponse of a specified response type on the given RestApi and resets it with the default settings."),
		mcp.WithString("restapi_id", mcp.Required(), mcp.Description("The string identifier of the associated RestApi.")),
		mcp.WithString("response_type", mcp.Required(), mcp.Description("The response type of the associated GatewayResponse.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletegatewayresponseHandler(cfg),
	}
}
