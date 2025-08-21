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

func GetclientcertificateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		clientcertificate_idVal, ok := args["clientcertificate_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: clientcertificate_id"), nil
		}
		clientcertificate_id, ok := clientcertificate_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: clientcertificate_id"), nil
		}
		url := fmt.Sprintf("%s/clientcertificates/%s", cfg.BaseURL, clientcertificate_id)
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
		var result models.ClientCertificate
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

func CreateGetclientcertificateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_clientcertificates_clientcertificate_id",
		mcp.WithDescription("Gets information about the current ClientCertificate resource."),
		mcp.WithString("clientcertificate_id", mcp.Required(), mcp.Description("The identifier of the ClientCertificate resource to be described.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetclientcertificateHandler(cfg),
	}
}
