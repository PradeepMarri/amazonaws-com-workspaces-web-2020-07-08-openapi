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

func GettruststorecertificateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		trustStoreArnVal, ok := args["trustStoreArn"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: trustStoreArn"), nil
		}
		trustStoreArn, ok := trustStoreArnVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: trustStoreArn"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["thumbprint"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("thumbprint=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/trustStores/%s/certificate#thumbprint%s", cfg.BaseURL, trustStoreArn, queryString)
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
		var result models.GetTrustStoreCertificateResponse
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

func CreateGettruststorecertificateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_trustStores_trustStoreArn_certificate#thumbprint",
		mcp.WithDescription("Gets the trust store certificate."),
		mcp.WithString("thumbprint", mcp.Required(), mcp.Description("The thumbprint of the trust store certificate.")),
		mcp.WithString("trustStoreArn", mcp.Required(), mcp.Description("The ARN of the trust store certificate.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GettruststorecertificateHandler(cfg),
	}
}
