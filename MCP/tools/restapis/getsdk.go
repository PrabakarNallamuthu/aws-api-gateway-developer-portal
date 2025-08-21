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

func GetsdkHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		stage_nameVal, ok := args["stage_name"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: stage_name"), nil
		}
		stage_name, ok := stage_nameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: stage_name"), nil
		}
		sdk_typeVal, ok := args["sdk_type"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: sdk_type"), nil
		}
		sdk_type, ok := sdk_typeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: sdk_type"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["parameters"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("parameters=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/restapis/%s/stages/%s/sdks/%s%s", cfg.BaseURL, restapi_id, stage_name, sdk_type, queryString)
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
		var result models.SdkResponse
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

func CreateGetsdkTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_restapis_restapi_id_stages_stage_name_sdks_sdk_type",
		mcp.WithDescription("Generates a client SDK for a RestApi and Stage."),
		mcp.WithString("restapi_id", mcp.Required(), mcp.Description("The string identifier of the associated RestApi.")),
		mcp.WithString("stage_name", mcp.Required(), mcp.Description("The name of the Stage that the SDK will use.")),
		mcp.WithString("sdk_type", mcp.Required(), mcp.Description("The language for the generated SDK. Currently <code>java</code>, <code>javascript</code>, <code>android</code>, <code>objectivec</code> (for iOS), <code>swift</code> (for iOS), and <code>ruby</code> are supported.")),
		mcp.WithObject("parameters", mcp.Description("A string-to-string key-value map of query parameters <code>sdkType</code>-dependent properties of the SDK. For <code>sdkType</code> of <code>objectivec</code> or <code>swift</code>, a parameter named <code>classPrefix</code> is required. For <code>sdkType</code> of <code>android</code>, parameters named <code>groupId</code>, <code>artifactId</code>, <code>artifactVersion</code>, and <code>invokerPackage</code> are required. For <code>sdkType</code> of <code>java</code>, parameters named <code>serviceName</code> and <code>javaPackageName</code> are required. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetsdkHandler(cfg),
	}
}
