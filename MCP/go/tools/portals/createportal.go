package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-workspaces-web/mcp-server/config"
	"github.com/amazon-workspaces-web/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateportalHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/portals", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.CreatePortalResponse
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

func CreateCreateportalTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_portals",
		mcp.WithDescription("Creates a web portal."),
		mcp.WithArray("tags", mcp.Description("Input parameter: The tags to add to the web portal. A tag is a key-value pair.")),
		mcp.WithObject("additionalEncryptionContext", mcp.Description("Input parameter: The additional encryption context of the portal.")),
		mcp.WithString("authenticationType", mcp.Description("Input parameter: <p>The type of authentication integration points used when signing into the web portal. Defaults to <code>Standard</code>.</p> <p> <code>Standard</code> web portals are authenticated directly through your identity provider. You need to call <code>CreateIdentityProvider</code> to integrate your identity provider with your web portal. User and group access to your web portal is controlled through your identity provider.</p> <p> <code>IAM_Identity_Center</code> web portals are authenticated through AWS IAM Identity Center (successor to AWS Single Sign-On). They provide additional features, such as IdP-initiated authentication. Identity sources (including external identity provider integration), plus user and group access to your web portal, can be configured in the IAM Identity Center.</p>")),
		mcp.WithString("customerManagedKey", mcp.Description("Input parameter: The customer managed key of the web portal.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: The name of the web portal. This is not visible to users who log into the web portal.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateportalHandler(cfg),
	}
}
