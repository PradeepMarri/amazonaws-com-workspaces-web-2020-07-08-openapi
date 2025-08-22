package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-workspaces-web/mcp-server/config"
	"github.com/amazon-workspaces-web/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func AssociateusersettingsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		portalArnVal, ok := args["portalArn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: portalArn"), nil
		}
		portalArn, ok := portalArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: portalArn"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["userSettingsArn"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("userSettingsArn=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/portals/%s/userSettings#userSettingsArn%s", cfg.BaseURL, portalArn, queryString)
		req, err := http.NewRequest("PUT", url, nil)
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
		var result models.AssociateUserSettingsResponse
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

func CreateAssociateusersettingsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_portals_portalArn_userSettings#userSettingsArn",
		mcp.WithDescription("Associates a user settings resource with a web portal."),
		mcp.WithString("portalArn", mcp.Required(), mcp.Description("The ARN of the web portal.")),
		mcp.WithString("userSettingsArn", mcp.Required(), mcp.Description("The ARN of the user settings.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AssociateusersettingsHandler(cfg),
	}
}
