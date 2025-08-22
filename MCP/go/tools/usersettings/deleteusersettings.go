package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amazon-workspaces-web/mcp-server/config"
	"github.com/amazon-workspaces-web/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DeleteusersettingsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		userSettingsArnVal, ok := args["userSettingsArn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: userSettingsArn"), nil
		}
		userSettingsArn, ok := userSettingsArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: userSettingsArn"), nil
		}
		url := fmt.Sprintf("%s/userSettings/%s", cfg.BaseURL, userSettingsArn)
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
		var result models.DeleteUserSettingsResponse
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

func CreateDeleteusersettingsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_userSettings_userSettingsArn",
		mcp.WithDescription("Deletes user settings."),
		mcp.WithString("userSettingsArn", mcp.Required(), mcp.Description("The ARN of the user settings.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeleteusersettingsHandler(cfg),
	}
}
