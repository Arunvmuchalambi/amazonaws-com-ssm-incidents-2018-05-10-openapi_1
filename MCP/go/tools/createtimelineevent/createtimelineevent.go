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

func CreatetimelineeventHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/createTimelineEvent", cfg.BaseURL)
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
		var result models.CreateTimelineEventOutput
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

func CreateCreatetimelineeventTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_createTimelineEvent",
		mcp.WithDescription("Creates a custom timeline event on the incident details page of an incident record. Incident Manager automatically creates timeline events that mark key moments during an incident. You can create custom timeline events to mark important events that Incident Manager can detect automatically."),
		mcp.WithString("clientToken", mcp.Description("Input parameter: A token that ensures that a client calls the action only once with the specified details.")),
		mcp.WithString("eventData", mcp.Required(), mcp.Description("Input parameter: A short description of the event.")),
		mcp.WithArray("eventReferences", mcp.Description("Input parameter: Adds one or more references to the <code>TimelineEvent</code>. A reference is an Amazon Web Services resource involved or associated with the incident. To specify a reference, enter its Amazon Resource Name (ARN). You can also specify a related item associated with a resource. For example, to specify an Amazon DynamoDB (DynamoDB) table as a resource, use the table's ARN. You can also specify an Amazon CloudWatch metric associated with the DynamoDB table as a related item.")),
		mcp.WithString("eventTime", mcp.Required(), mcp.Description("Input parameter: The time that the event occurred.")),
		mcp.WithString("eventType", mcp.Required(), mcp.Description("Input parameter: The type of event. You can create timeline events of type <code>Custom Event</code>.")),
		mcp.WithString("incidentRecordArn", mcp.Required(), mcp.Description("Input parameter: The Amazon Resource Name (ARN) of the incident record that the action adds the incident to.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatetimelineeventHandler(cfg),
	}
}
