package logs

import (
	"fmt"
)

const (
	paramNamespace        = "ParamNamespace"
	paramAppName          = "ParamAppName"
	paramAppId            = "ParamAppId"
	paramComponentName    = "ParamComponentName"
	paramJobComponentName = "ParamJobComponentName"
	paramJobName          = "ParamJobName"
	paramPipelineJobName  = "ParamPipelineJobName"
	paramPodName          = "ParamPodName"
	paramContainerId      = "ParamContainerId"
)

// KQL documentation: https://learn.microsoft.com/en-us/azure/data-explorer/kusto/query/

var (
	joinContainerLogV1 = `ContainerLog
	| where ContainerID in (containers)
	| project TimeGenerated, LogEntry
	| order by TimeGenerated desc`

	joinContainerLogV2 = `ContainerLogV2
	| where ContainerId  in (containers)
	| project TimeGenerated, LogEntry=LogMessage
	| order by TimeGenerated desc`

	joinContainerBoth = `let logv1=ContainerLog
	| where ContainerID in (containers)
	| project TimeGenerated, LogEntry;
	let logv2=ContainerLogV2
	| where ContainerId in (containers)
	| project TimeGenerated, LogEntry=tostring(LogMessage);
	union kind=outer logv1, logv2
	| order by TimeGenerated desc`
)

func getComponentInventoryQuery() string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-app-id"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| project TimeGenerated, Name, ContainerID, PodCreationTimeStamp, ContainerCreationTimeStamp=coalesce(ContainerCreationTimeStamp,todatetime(parse_json(ContainerLastStatus)["startedAt"]))
	| where isnotnull(ContainerCreationTimeStamp)
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp), ContainerCreationTimeStamp=min(ContainerCreationTimeStamp), LastTimeGenerated=max(TimeGenerated) by Name, ContainerID`,
		paramNamespace, paramAppName, paramAppId, paramComponentName)
}

func getComponentLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`let containers = KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-app-id"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| summarize by ContainerID;
	`+joinContainerLog,
		paramNamespace, paramAppName, paramAppId, paramComponentName)
}

func getComponentPodLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`let containers = KubePodInventory
	| where Namespace == %s and Name == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-app-id"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| summarize by ContainerID;
	`+joinContainerLog,
		paramNamespace, paramPodName, paramAppName, paramAppId, paramComponentName)
}

func getComponentContainerLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`let containers = KubePodInventory
	| where Namespace == %s and Name == %s and ContainerID == %s
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-app-id"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| summarize by ContainerID;
	`+joinContainerLog,
		paramNamespace, paramPodName, paramContainerId, paramAppName, paramAppId, paramComponentName)
}

func getJobInventoryQuery() string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-app-id"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp) by Name, ContainerID
    | join kind=inner ContainerInventory on ContainerID
    | project Name, PodCreationTimeStamp, ContainerID, ContainerLastKnownTimeStamp=coalesce(FinishedTime, TimeGenerated), CreatedTime
    | summarize PodCreationTimeStamp=min(PodCreationTimeStamp), ContainerCreationTimeStamp=min(CreatedTime), LastTimeGenerated=max(ContainerLastKnownTimeStamp) by Name, ContainerID`,
		paramNamespace, paramAppName, paramAppId, paramJobComponentName, paramJobName)
}

func getJobLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`let containers = KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-app-id"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s
	| summarize by ContainerID;
	`+joinContainerLog,
		paramNamespace, paramAppName, paramAppId, paramJobComponentName, paramJobName)
}

func getJobPodLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`let containers = KubePodInventory
	| where Namespace == %s and Name == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-app-id"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s
	| summarize by ContainerID;
	`+joinContainerLog,
		paramNamespace, paramPodName, paramAppName, paramAppId, paramJobComponentName, paramJobName)
}

func getJobContainerLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`let containers = KubePodInventory
	| where Namespace == %s and Name == %s and ContainerID == %s
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-app-id"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s
	| summarize by ContainerID;
	`+joinContainerLog,
		paramNamespace, paramPodName, paramContainerId, paramAppName, paramAppId, paramJobComponentName, paramJobName)
}

func getPipelineJobInventoryQuery() string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-job-name"] == %s and isempty(d["tekton.dev/task"])
	| extend ContainerNameShort=replace_string(ContainerName, strcat(PodUid,"/"), "")
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp) by Name, ContainerID, ContainerNameShort
	| join kind=inner ContainerInventory on ContainerID
	| project Name, PodCreationTimeStamp, ContainerID, ContainerNameShort, ContainerLastKnownTimeStamp=coalesce(FinishedTime, TimeGenerated), CreatedTime
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp), ContainerCreationTimeStamp=min(CreatedTime), LastTimeGenerated=max(ContainerLastKnownTimeStamp) by Name, ContainerID, ContainerNameShort`,
		paramNamespace, paramPipelineJobName)
}

func getPipelineJobContainerLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`let containers = KubePodInventory
	| where Namespace == %s and Name == %s and ContainerID == %s
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-job-name"] == %s and isempty(d["tekton.dev/task"])
	| summarize by ContainerID;
	`+joinContainerLog,
		paramNamespace, paramPodName, paramContainerId, paramPipelineJobName)
}
