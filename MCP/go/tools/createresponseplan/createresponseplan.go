package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/aws-systems-manager-incident-manager/mcp-server/config"
	"github.com/aws-systems-manager-incident-manager/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateresponseplanHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/createResponsePlan", cfg.BaseURL)
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
		var result models.CreateResponsePlanOutput
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

func CreateCreateresponseplanTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_createResponsePlan",
		mcp.WithDescription("Creates a response plan that automates the initial response to incidents. A response plan engages contacts, starts chat channel collaboration, and initiates runbooks at the beginning of an incident."),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: The short format name of the response plan. Can't include spaces.")),
		mcp.WithObject("chatChannel", mcp.Description("Input parameter: The Chatbot chat channel used for collaboration during an incident.")),
		mcp.WithString("clientToken", mcp.Description("Input parameter: A token ensuring that the operation is called only once with the specified details.")),
		mcp.WithArray("actions", mcp.Description("Input parameter: The actions that the response plan starts at the beginning of an incident.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: The long format of the response plan name. This field can contain spaces.")),
		mcp.WithObject("incidentTemplate", mcp.Required(), mcp.Description("Input parameter: Basic details used in creating a response plan. The response plan is then used to create an incident record.")),
		mcp.WithArray("integrations", mcp.Description("Input parameter: Information about third-party services integrated into the response plan.")),
		mcp.WithObject("tags", mcp.Description("Input parameter: A list of tags that you are adding to the response plan.")),
		mcp.WithArray("engagements", mcp.Description("Input parameter: The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateresponseplanHandler(cfg),
	}
}
