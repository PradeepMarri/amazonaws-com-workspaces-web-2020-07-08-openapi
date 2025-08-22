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

func GetidentityproviderHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		identityProviderArnVal, ok := args["identityProviderArn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: identityProviderArn"), nil
		}
		identityProviderArn, ok := identityProviderArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: identityProviderArn"), nil
		}
		url := fmt.Sprintf("%s/identityProviders/%s", cfg.BaseURL, identityProviderArn)
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
		var result models.GetIdentityProviderResponse
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

func CreateGetidentityproviderTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_identityProviders_identityProviderArn",
		mcp.WithDescription("Gets the identity provider."),
		mcp.WithString("identityProviderArn", mcp.Required(), mcp.Description("The ARN of the identity provider.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetidentityproviderHandler(cfg),
	}
}
