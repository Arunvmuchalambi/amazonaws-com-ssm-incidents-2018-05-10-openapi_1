package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws-systems-manager-incident-manager/mcp-server/config"
	"github.com/aws-systems-manager-incident-manager/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GettimelineeventHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["eventId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("eventId=%v", val))
		}
		if val, ok := args["incidentRecordArn"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("incidentRecordArn=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/getTimelineEvent#eventId&incidentRecordArn%s", cfg.BaseURL, queryString)
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
		var result models.GetTimelineEventOutput
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

func CreateGettimelineeventTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_getTimelineEvent#eventId&incidentRecordArn",
		mcp.WithDescription("Retrieves a timeline event based on its ID and incident record."),
		mcp.WithString("eventId", mcp.Required(), mcp.Description("The ID of the event. You can get an event's ID when you create it, or by using <code>ListTimelineEvents</code>.")),
		mcp.WithString("incidentRecordArn", mcp.Required(), mcp.Description("The Amazon Resource Name (ARN) of the incident that includes the timeline event.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GettimelineeventHandler(cfg),
	}
}
