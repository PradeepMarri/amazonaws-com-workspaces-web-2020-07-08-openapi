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

func CreateidentityproviderHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/identityProviders", cfg.BaseURL)
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
		var result models.CreateIdentityProviderResponse
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

func CreateCreateidentityproviderTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_identityProviders",
		mcp.WithDescription("Creates an identity provider resource that is then associated with a web portal."),
		mcp.WithString("identityProviderName", mcp.Required(), mcp.Description("Input parameter: The identity provider name.")),
		mcp.WithString("identityProviderType", mcp.Required(), mcp.Description("Input parameter: The identity provider type.")),
		mcp.WithString("portalArn", mcp.Required(), mcp.Description("Input parameter: The ARN of the web portal.")),
		mcp.WithObject("identityProviderDetails", mcp.Required(), mcp.Description("Input parameter: <p>The identity provider details. The following list describes the provider detail keys for each identity provider type. </p> <ul> <li> <p>For Google and Login with Amazon:</p> <ul> <li> <p> <code>client_id</code> </p> </li> <li> <p> <code>client_secret</code> </p> </li> <li> <p> <code>authorize_scopes</code> </p> </li> </ul> </li> <li> <p>For Facebook:</p> <ul> <li> <p> <code>client_id</code> </p> </li> <li> <p> <code>client_secret</code> </p> </li> <li> <p> <code>authorize_scopes</code> </p> </li> <li> <p> <code>api_version</code> </p> </li> </ul> </li> <li> <p>For Sign in with Apple:</p> <ul> <li> <p> <code>client_id</code> </p> </li> <li> <p> <code>team_id</code> </p> </li> <li> <p> <code>key_id</code> </p> </li> <li> <p> <code>private_key</code> </p> </li> <li> <p> <code>authorize_scopes</code> </p> </li> </ul> </li> <li> <p>For OIDC providers:</p> <ul> <li> <p> <code>client_id</code> </p> </li> <li> <p> <code>client_secret</code> </p> </li> <li> <p> <code>attributes_request_method</code> </p> </li> <li> <p> <code>oidc_issuer</code> </p> </li> <li> <p> <code>authorize_scopes</code> </p> </li> <li> <p> <code>authorize_url</code> <i>if not available from discovery URL specified by <code>oidc_issuer</code> key</i> </p> </li> <li> <p> <code>token_url</code> <i>if not available from discovery URL specified by <code>oidc_issuer</code> key</i> </p> </li> <li> <p> <code>attributes_url</code> <i>if not available from discovery URL specified by <code>oidc_issuer</code> key</i> </p> </li> <li> <p> <code>jwks_uri</code> <i>if not available from discovery URL specified by <code>oidc_issuer</code> key</i> </p> </li> </ul> </li> <li> <p>For SAML providers:</p> <ul> <li> <p> <code>MetadataFile</code> OR <code>MetadataURL</code> </p> </li> <li> <p> <code>IDPSignout</code> (boolean) <i>optional</i> </p> </li> </ul> </li> </ul>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateidentityproviderHandler(cfg),
	}
}
