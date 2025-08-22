package main

import (
	"github.com/aws-systems-manager-incident-manager/mcp-server/config"
	"github.com/aws-systems-manager-incident-manager/mcp-server/models"
	tools_gettimelineevent_eventid_incidentrecordarn "github.com/aws-systems-manager-incident-manager/mcp-server/tools/gettimelineevent_eventid_incidentrecordarn"
	tools_createreplicationset "github.com/aws-systems-manager-incident-manager/mcp-server/tools/createreplicationset"
	tools_putresourcepolicy "github.com/aws-systems-manager-incident-manager/mcp-server/tools/putresourcepolicy"
	tools_startincident "github.com/aws-systems-manager-incident-manager/mcp-server/tools/startincident"
	tools_tags "github.com/aws-systems-manager-incident-manager/mcp-server/tools/tags"
	tools_updaterelateditems "github.com/aws-systems-manager-incident-manager/mcp-server/tools/updaterelateditems"
	tools_updateresponseplan "github.com/aws-systems-manager-incident-manager/mcp-server/tools/updateresponseplan"
	tools_deletetimelineevent "github.com/aws-systems-manager-incident-manager/mcp-server/tools/deletetimelineevent"
	tools_getreplicationset_arn "github.com/aws-systems-manager-incident-manager/mcp-server/tools/getreplicationset_arn"
	tools_listincidentrecords "github.com/aws-systems-manager-incident-manager/mcp-server/tools/listincidentrecords"
	tools_getincidentrecord_arn "github.com/aws-systems-manager-incident-manager/mcp-server/tools/getincidentrecord_arn"
	tools_getresourcepolicies_resourcearn "github.com/aws-systems-manager-incident-manager/mcp-server/tools/getresourcepolicies_resourcearn"
	tools_updatedeletionprotection "github.com/aws-systems-manager-incident-manager/mcp-server/tools/updatedeletionprotection"
	tools_updateincidentrecord "github.com/aws-systems-manager-incident-manager/mcp-server/tools/updateincidentrecord"
	tools_updatereplicationset "github.com/aws-systems-manager-incident-manager/mcp-server/tools/updatereplicationset"
	tools_createtimelineevent "github.com/aws-systems-manager-incident-manager/mcp-server/tools/createtimelineevent"
	tools_listrelateditems "github.com/aws-systems-manager-incident-manager/mcp-server/tools/listrelateditems"
	tools_listreplicationsets "github.com/aws-systems-manager-incident-manager/mcp-server/tools/listreplicationsets"
	tools_getresponseplan_arn "github.com/aws-systems-manager-incident-manager/mcp-server/tools/getresponseplan_arn"
	tools_createresponseplan "github.com/aws-systems-manager-incident-manager/mcp-server/tools/createresponseplan"
	tools_listresponseplans "github.com/aws-systems-manager-incident-manager/mcp-server/tools/listresponseplans"
	tools_deleteresourcepolicy "github.com/aws-systems-manager-incident-manager/mcp-server/tools/deleteresourcepolicy"
	tools_deleteresponseplan "github.com/aws-systems-manager-incident-manager/mcp-server/tools/deleteresponseplan"
	tools_listtimelineevents "github.com/aws-systems-manager-incident-manager/mcp-server/tools/listtimelineevents"
	tools_updatetimelineevent "github.com/aws-systems-manager-incident-manager/mcp-server/tools/updatetimelineevent"
	tools_deleteincidentrecord "github.com/aws-systems-manager-incident-manager/mcp-server/tools/deleteincidentrecord"
	tools_deletereplicationset_arn "github.com/aws-systems-manager-incident-manager/mcp-server/tools/deletereplicationset_arn"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_gettimelineevent_eventid_incidentrecordarn.CreateGettimelineeventTool(cfg),
		tools_createreplicationset.CreateCreatereplicationsetTool(cfg),
		tools_putresourcepolicy.CreatePutresourcepolicyTool(cfg),
		tools_startincident.CreateStartincidentTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_updaterelateditems.CreateUpdaterelateditemsTool(cfg),
		tools_updateresponseplan.CreateUpdateresponseplanTool(cfg),
		tools_deletetimelineevent.CreateDeletetimelineeventTool(cfg),
		tools_getreplicationset_arn.CreateGetreplicationsetTool(cfg),
		tools_listincidentrecords.CreateListincidentrecordsTool(cfg),
		tools_getincidentrecord_arn.CreateGetincidentrecordTool(cfg),
		tools_getresourcepolicies_resourcearn.CreateGetresourcepoliciesTool(cfg),
		tools_updatedeletionprotection.CreateUpdatedeletionprotectionTool(cfg),
		tools_updateincidentrecord.CreateUpdateincidentrecordTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_updatereplicationset.CreateUpdatereplicationsetTool(cfg),
		tools_createtimelineevent.CreateCreatetimelineeventTool(cfg),
		tools_listrelateditems.CreateListrelateditemsTool(cfg),
		tools_listreplicationsets.CreateListreplicationsetsTool(cfg),
		tools_getresponseplan_arn.CreateGetresponseplanTool(cfg),
		tools_createresponseplan.CreateCreateresponseplanTool(cfg),
		tools_listresponseplans.CreateListresponseplansTool(cfg),
		tools_deleteresourcepolicy.CreateDeleteresourcepolicyTool(cfg),
		tools_deleteresponseplan.CreateDeleteresponseplanTool(cfg),
		tools_listtimelineevents.CreateListtimelineeventsTool(cfg),
		tools_updatetimelineevent.CreateUpdatetimelineeventTool(cfg),
		tools_deleteincidentrecord.CreateDeleteincidentrecordTool(cfg),
		tools_deletereplicationset_arn.CreateDeletereplicationsetTool(cfg),
	}
}
