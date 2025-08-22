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

func UpdateresponseplanHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/updateResponsePlan", cfg.BaseURL)
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

func CreateUpdateresponseplanTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_updateResponsePlan",
		mcp.WithDescription("Updates the specified response plan."),
		mcp.WithString("incidentTemplateTitle", mcp.Description("Input parameter: The short format name of the incident. The title can't contain spaces.")),
		mcp.WithArray("actions", mcp.Description("Input parameter: The actions that this response plan takes at the beginning of an incident.")),
		mcp.WithString("clientToken", mcp.Description("Input parameter: A token ensuring that the operation is called only once with the specified details.")),
		mcp.WithString("incidentTemplateDedupeString", mcp.Description("Input parameter: The string Incident Manager uses to prevent duplicate incidents from being created by the same incident in the same account.")),
		mcp.WithNumber("incidentTemplateImpact", mcp.Description("Input parameter: <p>Defines the impact to the customers. Providing an impact overwrites the impact provided by a response plan.</p> <p class=\"title\"> <b>Possible impacts:</b> </p> <ul> <li> <p> <code>5</code> - Severe impact</p> </li> <li> <p> <code>4</code> - High impact</p> </li> <li> <p> <code>3</code> - Medium impact</p> </li> <li> <p> <code>2</code> - Low impact</p> </li> <li> <p> <code>1</code> - No impact</p> </li> </ul>")),
		mcp.WithString("incidentTemplateSummary", mcp.Description("Input parameter: A brief summary of the incident. This typically contains what has happened, what's currently happening, and next steps.")),
		mcp.WithObject("incidentTemplateTags", mcp.Description("Input parameter: Tags to assign to the template. When the <code>StartIncident</code> API action is called, Incident Manager assigns the tags specified in the template to the incident. To call this action, you must also have permission to call the <code>TagResource</code> API action for the incident record resource.")),
		mcp.WithArray("integrations", mcp.Description("Input parameter: Information about third-party services integrated into the response plan.")),
		mcp.WithString("arn", mcp.Required(), mcp.Description("Input parameter: The Amazon Resource Name (ARN) of the response plan.")),
		mcp.WithObject("chatChannel", mcp.Description("Input parameter: The Chatbot chat channel used for collaboration during an incident.")),
		mcp.WithString("displayName", mcp.Description("Input parameter: The long format name of the response plan. The display name can't contain spaces.")),
		mcp.WithArray("engagements", mcp.Description("Input parameter: The Amazon Resource Name (ARN) for the contacts and escalation plans that the response plan engages during an incident.")),
		mcp.WithArray("incidentTemplateNotificationTargets", mcp.Description("Input parameter: The Amazon SNS targets that are notified when updates are made to an incident.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateresponseplanHandler(cfg),
	}
}
