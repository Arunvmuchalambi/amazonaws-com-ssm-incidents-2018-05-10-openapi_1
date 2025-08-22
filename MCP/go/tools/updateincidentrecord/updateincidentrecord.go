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

func UpdateincidentrecordHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/updateIncidentRecord", cfg.BaseURL)
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

func CreateUpdateincidentrecordTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_updateIncidentRecord",
		mcp.WithDescription("Update the details of an incident record. You can use this operation to update an incident record from the defined chat channel. For more information about using actions in chat channels, see <a href="https://docs.aws.amazon.com/incident-manager/latest/userguide/chat.html#chat-interact">Interacting through chat</a>."),
		mcp.WithString("clientToken", mcp.Description("Input parameter: A token that ensures that a client calls the operation only once with the specified details.")),
		mcp.WithNumber("impact", mcp.Description("Input parameter: <p>Defines the impact of the incident to customers and applications. If you provide an impact for an incident, it overwrites the impact provided by the response plan.</p> <p class=\"title\"> <b>Possible impacts:</b> </p> <ul> <li> <p> <code>1</code> - Critical impact, full application failure that impacts many to all customers. </p> </li> <li> <p> <code>2</code> - High impact, partial application failure with impact to many customers.</p> </li> <li> <p> <code>3</code> - Medium impact, the application is providing reduced service to customers.</p> </li> <li> <p> <code>4</code> - Low impact, customer aren't impacted by the problem yet.</p> </li> <li> <p> <code>5</code> - No impact, customers aren't currently impacted but urgent action is needed to avoid impact.</p> </li> </ul>")),
		mcp.WithArray("notificationTargets", mcp.Description("Input parameter: <p>The Amazon SNS targets that Incident Manager notifies when a client updates an incident.</p> <p>Using multiple SNS topics creates redundancy in the event that a Region is down during the incident.</p>")),
		mcp.WithString("status", mcp.Description("Input parameter: The status of the incident. Possible statuses are <code>Open</code> or <code>Resolved</code>.")),
		mcp.WithString("summary", mcp.Description("Input parameter: A longer description of what occurred during the incident.")),
		mcp.WithString("title", mcp.Description("Input parameter: A brief description of the incident.")),
		mcp.WithString("arn", mcp.Required(), mcp.Description("Input parameter: The Amazon Resource Name (ARN) of the incident record you are updating.")),
		mcp.WithObject("chatChannel", mcp.Description("Input parameter: The Chatbot chat channel used for collaboration during an incident.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateincidentrecordHandler(cfg),
	}
}
