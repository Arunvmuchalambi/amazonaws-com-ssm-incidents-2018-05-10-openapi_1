package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListRelatedItemsInput represents the ListRelatedItemsInput schema from the OpenAPI specification
type ListRelatedItemsInput struct {
	Incidentrecordarn interface{} `json:"incidentRecordArn"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// Action represents the Action schema from the OpenAPI specification
type Action struct {
	Ssmautomation interface{} `json:"ssmAutomation,omitempty"`
}

// DeleteResponsePlanOutput represents the DeleteResponsePlanOutput schema from the OpenAPI specification
type DeleteResponsePlanOutput struct {
}

// GetReplicationSetInput represents the GetReplicationSetInput schema from the OpenAPI specification
type GetReplicationSetInput struct {
}

// Filter represents the Filter schema from the OpenAPI specification
type Filter struct {
	Key interface{} `json:"key"`
	Condition interface{} `json:"condition"`
}

// CreateResponsePlanInput represents the CreateResponsePlanInput schema from the OpenAPI specification
type CreateResponsePlanInput struct {
	Actions interface{} `json:"actions,omitempty"`
	Incidenttemplate interface{} `json:"incidentTemplate"`
	Tags interface{} `json:"tags,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Engagements interface{} `json:"engagements,omitempty"`
	Chatchannel interface{} `json:"chatChannel,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Integrations interface{} `json:"integrations,omitempty"`
	Name interface{} `json:"name"`
}

// UpdateReplicationSetOutput represents the UpdateReplicationSetOutput schema from the OpenAPI specification
type UpdateReplicationSetOutput struct {
}

// Condition represents the Condition schema from the OpenAPI specification
type Condition struct {
	Before interface{} `json:"before,omitempty"`
	Equals interface{} `json:"equals,omitempty"`
	After interface{} `json:"after,omitempty"`
}

// DeleteResourcePolicyOutput represents the DeleteResourcePolicyOutput schema from the OpenAPI specification
type DeleteResourcePolicyOutput struct {
}

// SsmParameters represents the SsmParameters schema from the OpenAPI specification
type SsmParameters struct {
}

// TimelineEvent represents the TimelineEvent schema from the OpenAPI specification
type TimelineEvent struct {
	Eventid interface{} `json:"eventId"`
	Eventreferences interface{} `json:"eventReferences,omitempty"`
	Eventtime interface{} `json:"eventTime"`
	Eventtype interface{} `json:"eventType"`
	Eventupdatedtime interface{} `json:"eventUpdatedTime"`
	Incidentrecordarn interface{} `json:"incidentRecordArn"`
	Eventdata interface{} `json:"eventData"`
}

// PutResourcePolicyInput represents the PutResourcePolicyInput schema from the OpenAPI specification
type PutResourcePolicyInput struct {
	Policy interface{} `json:"policy"`
	Resourcearn interface{} `json:"resourceArn"`
}

// GetResponsePlanInput represents the GetResponsePlanInput schema from the OpenAPI specification
type GetResponsePlanInput struct {
}

// GetIncidentRecordOutput represents the GetIncidentRecordOutput schema from the OpenAPI specification
type GetIncidentRecordOutput struct {
	Incidentrecord interface{} `json:"incidentRecord"`
}

// EventSummary represents the EventSummary schema from the OpenAPI specification
type EventSummary struct {
	Eventupdatedtime interface{} `json:"eventUpdatedTime"`
	Incidentrecordarn interface{} `json:"incidentRecordArn"`
	Eventid interface{} `json:"eventId"`
	Eventreferences interface{} `json:"eventReferences,omitempty"`
	Eventtime interface{} `json:"eventTime"`
	Eventtype interface{} `json:"eventType"`
}

// GetResponsePlanOutput represents the GetResponsePlanOutput schema from the OpenAPI specification
type GetResponsePlanOutput struct {
	Engagements interface{} `json:"engagements,omitempty"`
	Incidenttemplate interface{} `json:"incidentTemplate"`
	Integrations interface{} `json:"integrations,omitempty"`
	Name interface{} `json:"name"`
	Actions interface{} `json:"actions,omitempty"`
	Arn interface{} `json:"arn"`
	Chatchannel interface{} `json:"chatChannel,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
}

// UpdateReplicationSetInput represents the UpdateReplicationSetInput schema from the OpenAPI specification
type UpdateReplicationSetInput struct {
	Actions interface{} `json:"actions"`
	Arn interface{} `json:"arn"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// ListTimelineEventsOutput represents the ListTimelineEventsOutput schema from the OpenAPI specification
type ListTimelineEventsOutput struct {
	Eventsummaries interface{} `json:"eventSummaries"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// RegionMapInputValue represents the RegionMapInputValue schema from the OpenAPI specification
type RegionMapInputValue struct {
	Ssekmskeyid interface{} `json:"sseKmsKeyId,omitempty"`
}

// ListReplicationSetsOutput represents the ListReplicationSetsOutput schema from the OpenAPI specification
type ListReplicationSetsOutput struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Replicationsetarns interface{} `json:"replicationSetArns"`
}

// EmptyChatChannel represents the EmptyChatChannel schema from the OpenAPI specification
type EmptyChatChannel struct {
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// ListRelatedItemsOutput represents the ListRelatedItemsOutput schema from the OpenAPI specification
type ListRelatedItemsOutput struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Relateditems interface{} `json:"relatedItems"`
}

// RegionInfoMap represents the RegionInfoMap schema from the OpenAPI specification
type RegionInfoMap struct {
}

// RegionMapInput represents the RegionMapInput schema from the OpenAPI specification
type RegionMapInput struct {
}

// DeleteTimelineEventOutput represents the DeleteTimelineEventOutput schema from the OpenAPI specification
type DeleteTimelineEventOutput struct {
}

// CreateReplicationSetOutput represents the CreateReplicationSetOutput schema from the OpenAPI specification
type CreateReplicationSetOutput struct {
	Arn interface{} `json:"arn"`
}

// GetTimelineEventOutput represents the GetTimelineEventOutput schema from the OpenAPI specification
type GetTimelineEventOutput struct {
	Event interface{} `json:"event"`
}

// IncidentRecordSummary represents the IncidentRecordSummary schema from the OpenAPI specification
type IncidentRecordSummary struct {
	Title interface{} `json:"title"`
	Arn interface{} `json:"arn"`
	Creationtime interface{} `json:"creationTime"`
	Impact interface{} `json:"impact"`
	Incidentrecordsource interface{} `json:"incidentRecordSource"`
	Resolvedtime interface{} `json:"resolvedTime,omitempty"`
	Status interface{} `json:"status"`
}

// CreateReplicationSetInput represents the CreateReplicationSetInput schema from the OpenAPI specification
type CreateReplicationSetInput struct {
	Regions interface{} `json:"regions"`
	Tags interface{} `json:"tags,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// UpdateRelatedItemsInput represents the UpdateRelatedItemsInput schema from the OpenAPI specification
type UpdateRelatedItemsInput struct {
	Incidentrecordarn interface{} `json:"incidentRecordArn"`
	Relateditemsupdate interface{} `json:"relatedItemsUpdate"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// UpdateIncidentRecordInput represents the UpdateIncidentRecordInput schema from the OpenAPI specification
type UpdateIncidentRecordInput struct {
	Summary interface{} `json:"summary,omitempty"`
	Title interface{} `json:"title,omitempty"`
	Arn interface{} `json:"arn"`
	Chatchannel interface{} `json:"chatChannel,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Impact interface{} `json:"impact,omitempty"`
	Notificationtargets interface{} `json:"notificationTargets,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// RegionInfo represents the RegionInfo schema from the OpenAPI specification
type RegionInfo struct {
	Ssekmskeyid interface{} `json:"sseKmsKeyId,omitempty"`
	Status interface{} `json:"status"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Statusupdatedatetime interface{} `json:"statusUpdateDateTime"`
}

// GetResourcePoliciesInput represents the GetResourcePoliciesInput schema from the OpenAPI specification
type GetResourcePoliciesInput struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ResponsePlanSummary represents the ResponsePlanSummary schema from the OpenAPI specification
type ResponsePlanSummary struct {
	Arn interface{} `json:"arn"`
	Displayname interface{} `json:"displayName,omitempty"`
	Name interface{} `json:"name"`
}

// GetResourcePoliciesOutput represents the GetResourcePoliciesOutput schema from the OpenAPI specification
type GetResourcePoliciesOutput struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Resourcepolicies interface{} `json:"resourcePolicies"`
}

// DeleteResponsePlanInput represents the DeleteResponsePlanInput schema from the OpenAPI specification
type DeleteResponsePlanInput struct {
	Arn interface{} `json:"arn"`
}

// DeleteReplicationSetOutput represents the DeleteReplicationSetOutput schema from the OpenAPI specification
type DeleteReplicationSetOutput struct {
}

// DeleteIncidentRecordOutput represents the DeleteIncidentRecordOutput schema from the OpenAPI specification
type DeleteIncidentRecordOutput struct {
}

// GetReplicationSetOutput represents the GetReplicationSetOutput schema from the OpenAPI specification
type GetReplicationSetOutput struct {
	Replicationset interface{} `json:"replicationSet"`
}

// UpdateIncidentRecordOutput represents the UpdateIncidentRecordOutput schema from the OpenAPI specification
type UpdateIncidentRecordOutput struct {
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// EventReference represents the EventReference schema from the OpenAPI specification
type EventReference struct {
	Relateditemid interface{} `json:"relatedItemId,omitempty"`
	Resource interface{} `json:"resource,omitempty"`
}

// PagerDutyIncidentDetail represents the PagerDutyIncidentDetail schema from the OpenAPI specification
type PagerDutyIncidentDetail struct {
	Autoresolve interface{} `json:"autoResolve,omitempty"`
	Id interface{} `json:"id"`
	Secretid interface{} `json:"secretId,omitempty"`
}

// StartIncidentInput represents the StartIncidentInput schema from the OpenAPI specification
type StartIncidentInput struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Impact interface{} `json:"impact,omitempty"`
	Relateditems interface{} `json:"relatedItems,omitempty"`
	Responseplanarn interface{} `json:"responsePlanArn"`
	Title interface{} `json:"title,omitempty"`
	Triggerdetails interface{} `json:"triggerDetails,omitempty"`
}

// UpdateReplicationSetAction represents the UpdateReplicationSetAction schema from the OpenAPI specification
type UpdateReplicationSetAction struct {
	Addregionaction interface{} `json:"addRegionAction,omitempty"`
	Deleteregionaction interface{} `json:"deleteRegionAction,omitempty"`
}

// DeleteReplicationSetInput represents the DeleteReplicationSetInput schema from the OpenAPI specification
type DeleteReplicationSetInput struct {
}

// DeleteIncidentRecordInput represents the DeleteIncidentRecordInput schema from the OpenAPI specification
type DeleteIncidentRecordInput struct {
	Arn interface{} `json:"arn"`
}

// AutomationExecution represents the AutomationExecution schema from the OpenAPI specification
type AutomationExecution struct {
	Ssmexecutionarn interface{} `json:"ssmExecutionArn,omitempty"`
}

// ReplicationSet represents the ReplicationSet schema from the OpenAPI specification
type ReplicationSet struct {
	Lastmodifiedby interface{} `json:"lastModifiedBy"`
	Lastmodifiedtime interface{} `json:"lastModifiedTime"`
	Regionmap interface{} `json:"regionMap"`
	Status interface{} `json:"status"`
	Arn interface{} `json:"arn,omitempty"`
	Createdby interface{} `json:"createdBy"`
	Createdtime interface{} `json:"createdTime"`
	Deletionprotected interface{} `json:"deletionProtected"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// GetIncidentRecordInput represents the GetIncidentRecordInput schema from the OpenAPI specification
type GetIncidentRecordInput struct {
}

// ListIncidentRecordsInput represents the ListIncidentRecordsInput schema from the OpenAPI specification
type ListIncidentRecordsInput struct {
	Filters interface{} `json:"filters,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListTimelineEventsInput represents the ListTimelineEventsInput schema from the OpenAPI specification
type ListTimelineEventsInput struct {
	Incidentrecordarn interface{} `json:"incidentRecordArn"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Sortby interface{} `json:"sortBy,omitempty"`
	Sortorder interface{} `json:"sortOrder,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
}

// ListIncidentRecordsOutput represents the ListIncidentRecordsOutput schema from the OpenAPI specification
type ListIncidentRecordsOutput struct {
	Incidentrecordsummaries interface{} `json:"incidentRecordSummaries"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UpdateResponsePlanInput represents the UpdateResponsePlanInput schema from the OpenAPI specification
type UpdateResponsePlanInput struct {
	Incidenttemplatenotificationtargets interface{} `json:"incidentTemplateNotificationTargets,omitempty"`
	Incidenttemplatetags interface{} `json:"incidentTemplateTags,omitempty"`
	Integrations interface{} `json:"integrations,omitempty"`
	Arn interface{} `json:"arn"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Incidenttemplatetitle interface{} `json:"incidentTemplateTitle,omitempty"`
	Actions interface{} `json:"actions,omitempty"`
	Chatchannel interface{} `json:"chatChannel,omitempty"`
	Engagements interface{} `json:"engagements,omitempty"`
	Incidenttemplateimpact interface{} `json:"incidentTemplateImpact,omitempty"`
	Incidenttemplatesummary interface{} `json:"incidentTemplateSummary,omitempty"`
	Incidenttemplatededupestring interface{} `json:"incidentTemplateDedupeString,omitempty"`
}

// PagerDutyIncidentConfiguration represents the PagerDutyIncidentConfiguration schema from the OpenAPI specification
type PagerDutyIncidentConfiguration struct {
	Serviceid interface{} `json:"serviceId"`
}

// UpdateTimelineEventOutput represents the UpdateTimelineEventOutput schema from the OpenAPI specification
type UpdateTimelineEventOutput struct {
}

// IncidentRecord represents the IncidentRecord schema from the OpenAPI specification
type IncidentRecord struct {
	Chatchannel interface{} `json:"chatChannel,omitempty"`
	Dedupestring interface{} `json:"dedupeString"`
	Resolvedtime interface{} `json:"resolvedTime,omitempty"`
	Status interface{} `json:"status"`
	Arn interface{} `json:"arn"`
	Lastmodifiedby interface{} `json:"lastModifiedBy"`
	Notificationtargets interface{} `json:"notificationTargets,omitempty"`
	Title interface{} `json:"title"`
	Automationexecutions interface{} `json:"automationExecutions,omitempty"`
	Lastmodifiedtime interface{} `json:"lastModifiedTime"`
	Creationtime interface{} `json:"creationTime"`
	Impact interface{} `json:"impact"`
	Incidentrecordsource interface{} `json:"incidentRecordSource"`
	Summary interface{} `json:"summary,omitempty"`
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	Pagerdutyconfiguration interface{} `json:"pagerDutyConfiguration,omitempty"`
}

// StartIncidentOutput represents the StartIncidentOutput schema from the OpenAPI specification
type StartIncidentOutput struct {
	Incidentrecordarn interface{} `json:"incidentRecordArn"`
}

// IncidentRecordSource represents the IncidentRecordSource schema from the OpenAPI specification
type IncidentRecordSource struct {
	Invokedby interface{} `json:"invokedBy,omitempty"`
	Resourcearn interface{} `json:"resourceArn,omitempty"`
	Source interface{} `json:"source"`
	Createdby interface{} `json:"createdBy"`
}

// DeleteRegionAction represents the DeleteRegionAction schema from the OpenAPI specification
type DeleteRegionAction struct {
	Regionname interface{} `json:"regionName"`
}

// UpdateRelatedItemsOutput represents the UpdateRelatedItemsOutput schema from the OpenAPI specification
type UpdateRelatedItemsOutput struct {
}

// GetTimelineEventInput represents the GetTimelineEventInput schema from the OpenAPI specification
type GetTimelineEventInput struct {
}

// ListReplicationSetsInput represents the ListReplicationSetsInput schema from the OpenAPI specification
type ListReplicationSetsInput struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DynamicSsmParameters represents the DynamicSsmParameters schema from the OpenAPI specification
type DynamicSsmParameters struct {
}

// ItemIdentifier represents the ItemIdentifier schema from the OpenAPI specification
type ItemIdentifier struct {
	Value interface{} `json:"value"`
	TypeField interface{} `json:"type"`
}

// CreateTimelineEventInput represents the CreateTimelineEventInput schema from the OpenAPI specification
type CreateTimelineEventInput struct {
	Eventreferences interface{} `json:"eventReferences,omitempty"`
	Eventtime interface{} `json:"eventTime"`
	Eventtype interface{} `json:"eventType"`
	Incidentrecordarn interface{} `json:"incidentRecordArn"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Eventdata interface{} `json:"eventData"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// PutResourcePolicyOutput represents the PutResourcePolicyOutput schema from the OpenAPI specification
type PutResourcePolicyOutput struct {
	Policyid interface{} `json:"policyId"`
}

// IncidentTemplate represents the IncidentTemplate schema from the OpenAPI specification
type IncidentTemplate struct {
	Summary interface{} `json:"summary,omitempty"`
	Title interface{} `json:"title"`
	Dedupestring interface{} `json:"dedupeString,omitempty"`
	Impact interface{} `json:"impact"`
	Incidenttags interface{} `json:"incidentTags,omitempty"`
	Notificationtargets interface{} `json:"notificationTargets,omitempty"`
}

// CreateTimelineEventOutput represents the CreateTimelineEventOutput schema from the OpenAPI specification
type CreateTimelineEventOutput struct {
	Eventid interface{} `json:"eventId"`
	Incidentrecordarn interface{} `json:"incidentRecordArn"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
}

// TriggerDetails represents the TriggerDetails schema from the OpenAPI specification
type TriggerDetails struct {
	Timestamp interface{} `json:"timestamp"`
	Triggerarn interface{} `json:"triggerArn,omitempty"`
	Rawdata interface{} `json:"rawData,omitempty"`
	Source interface{} `json:"source"`
}

// RelatedItem represents the RelatedItem schema from the OpenAPI specification
type RelatedItem struct {
	Generatedid interface{} `json:"generatedId,omitempty"`
	Identifier interface{} `json:"identifier"`
	Title interface{} `json:"title,omitempty"`
}

// DeleteTimelineEventInput represents the DeleteTimelineEventInput schema from the OpenAPI specification
type DeleteTimelineEventInput struct {
	Eventid interface{} `json:"eventId"`
	Incidentrecordarn interface{} `json:"incidentRecordArn"`
}

// RelatedItemsUpdate represents the RelatedItemsUpdate schema from the OpenAPI specification
type RelatedItemsUpdate struct {
	Itemtoadd interface{} `json:"itemToAdd,omitempty"`
	Itemtoremove interface{} `json:"itemToRemove,omitempty"`
}

// ItemValue represents the ItemValue schema from the OpenAPI specification
type ItemValue struct {
	Metricdefinition interface{} `json:"metricDefinition,omitempty"`
	Pagerdutyincidentdetail interface{} `json:"pagerDutyIncidentDetail,omitempty"`
	Url interface{} `json:"url,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// NotificationTargetItem represents the NotificationTargetItem schema from the OpenAPI specification
type NotificationTargetItem struct {
	Snstopicarn interface{} `json:"snsTopicArn,omitempty"`
}

// ChatChannel represents the ChatChannel schema from the OpenAPI specification
type ChatChannel struct {
	Chatbotsns interface{} `json:"chatbotSns,omitempty"`
	Empty interface{} `json:"empty,omitempty"`
}

// ListResponsePlansOutput represents the ListResponsePlansOutput schema from the OpenAPI specification
type ListResponsePlansOutput struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Responseplansummaries interface{} `json:"responsePlanSummaries"`
}

// DynamicSsmParameterValue represents the DynamicSsmParameterValue schema from the OpenAPI specification
type DynamicSsmParameterValue struct {
	Variable interface{} `json:"variable,omitempty"`
}

// TagMapUpdate represents the TagMapUpdate schema from the OpenAPI specification
type TagMapUpdate struct {
}

// ListResponsePlansInput represents the ListResponsePlansInput schema from the OpenAPI specification
type ListResponsePlansInput struct {
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// AddRegionAction represents the AddRegionAction schema from the OpenAPI specification
type AddRegionAction struct {
	Regionname interface{} `json:"regionName"`
	Ssekmskeyid interface{} `json:"sseKmsKeyId,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags"`
}

// UpdateDeletionProtectionInput represents the UpdateDeletionProtectionInput schema from the OpenAPI specification
type UpdateDeletionProtectionInput struct {
	Deletionprotected interface{} `json:"deletionProtected"`
	Arn interface{} `json:"arn"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// UpdateResponsePlanOutput represents the UpdateResponsePlanOutput schema from the OpenAPI specification
type UpdateResponsePlanOutput struct {
}

// UpdateDeletionProtectionOutput represents the UpdateDeletionProtectionOutput schema from the OpenAPI specification
type UpdateDeletionProtectionOutput struct {
}

// AttributeValueList represents the AttributeValueList schema from the OpenAPI specification
type AttributeValueList struct {
	Integervalues interface{} `json:"integerValues,omitempty"`
	Stringvalues interface{} `json:"stringValues,omitempty"`
}

// DeleteResourcePolicyInput represents the DeleteResourcePolicyInput schema from the OpenAPI specification
type DeleteResourcePolicyInput struct {
	Policyid interface{} `json:"policyId"`
	Resourcearn interface{} `json:"resourceArn"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// ResourcePolicy represents the ResourcePolicy schema from the OpenAPI specification
type ResourcePolicy struct {
	Policydocument interface{} `json:"policyDocument"`
	Policyid interface{} `json:"policyId"`
	Ramresourceshareregion interface{} `json:"ramResourceShareRegion"`
}

// SsmAutomation represents the SsmAutomation schema from the OpenAPI specification
type SsmAutomation struct {
	Parameters interface{} `json:"parameters,omitempty"`
	Rolearn interface{} `json:"roleArn"`
	Targetaccount interface{} `json:"targetAccount,omitempty"`
	Documentname interface{} `json:"documentName"`
	Documentversion interface{} `json:"documentVersion,omitempty"`
	Dynamicparameters interface{} `json:"dynamicParameters,omitempty"`
}

// CreateResponsePlanOutput represents the CreateResponsePlanOutput schema from the OpenAPI specification
type CreateResponsePlanOutput struct {
	Arn interface{} `json:"arn"`
}

// UpdateTimelineEventInput represents the UpdateTimelineEventInput schema from the OpenAPI specification
type UpdateTimelineEventInput struct {
	Eventreferences interface{} `json:"eventReferences,omitempty"`
	Eventtime interface{} `json:"eventTime,omitempty"`
	Eventtype interface{} `json:"eventType,omitempty"`
	Incidentrecordarn interface{} `json:"incidentRecordArn"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Eventdata interface{} `json:"eventData,omitempty"`
	Eventid interface{} `json:"eventId"`
}

// PagerDutyConfiguration represents the PagerDutyConfiguration schema from the OpenAPI specification
type PagerDutyConfiguration struct {
	Secretid interface{} `json:"secretId"`
	Name interface{} `json:"name"`
	Pagerdutyincidentconfiguration interface{} `json:"pagerDutyIncidentConfiguration"`
}
