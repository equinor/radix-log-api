package logs

import (
	"fmt"
)

const (
	paramNamespace        = "ParamNamespace"
	paramAppName          = "ParamAppName"
	paramComponentName    = "ParamComponentName"
	paramJobComponentName = "ParamJobComponentName"
	paramJobName          = "ParamJobName"
	paramPipelineJobName  = "ParamPipelineJobName"
	paramPodName          = "ParamPodName"
	paramContainerId      = "ParamContainerId"
)

// KQL documentation: https://learn.microsoft.com/en-us/azure/data-explorer/kusto/query/

var (
	joinContainerLogV1 = `| join kind=inner ContainerLog on $left.ContainerID==$right.ContainerID
	| project TimeGenerated, Name, ContainerID, LogEntry
	| sort by TimeGenerated desc`

	joinContainerLogV2 = `| join kind=inner ContainerLog on $left.ContainerID==$right.ContainerID
	| project TimeGenerated, Name, ContainerID, LogEntry
	| sort by TimeGenerated desc`

	joinContainerBoth = `| join kind=inner ContainerLog on $left.ContainerID==$right.ContainerID
	| project TimeGenerated, Name, ContainerID, LogEntry
	| sort by TimeGenerated desc`
)

func getComponentInventoryQuery() string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| project TimeGenerated, Name, ContainerID, PodCreationTimeStamp, ContainerCreationTimeStamp=coalesce(ContainerCreationTimeStamp,todatetime(parse_json(ContainerLastStatus)["startedAt"]))
	| where isnotnull(ContainerCreationTimeStamp)
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp), ContainerCreationTimeStamp=min(ContainerCreationTimeStamp), LastTimeGenerated=max(TimeGenerated) by Name, ContainerID`,
		paramNamespace, paramAppName, paramComponentName)
}

func getComponentLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramAppName, paramComponentName)
}

func getComponentPodLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and Name == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramPodName, paramAppName, paramComponentName)
}

func getComponentContainerLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and Name == %s and ContainerID == %s
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramPodName, paramContainerId, paramAppName, paramComponentName)
}

func getJobInventoryQuery() string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp) by Name, ContainerID
    | join kind=inner ContainerInventory on ContainerID
    | project Name, PodCreationTimeStamp, ContainerID, ContainerLastKnownTimeStamp=coalesce(FinishedTime, TimeGenerated), CreatedTime
    | summarize PodCreationTimeStamp=min(PodCreationTimeStamp), ContainerCreationTimeStamp=min(CreatedTime), LastTimeGenerated=max(ContainerLastKnownTimeStamp) by Name, ContainerID`,
		paramNamespace, paramAppName, paramJobComponentName, paramJobName)
}

func getJobLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramAppName, paramJobComponentName, paramJobName)
}

func getJobPodLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and Name == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramPodName, paramAppName, paramJobComponentName, paramJobName)
}

func getJobContainerLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and Name == %s and ContainerID == %s
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramPodName, paramContainerId, paramAppName, paramJobComponentName, paramJobName)
}

func getPipelineJobInventoryQuery() string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-job-name"] == %s and isempty(d["tekton.dev/task"])
	| extend ContainerNameShort=replace_string {
ontainerName, strcat(PodUid,"/"), "")
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp) by Name, ContainerID, ContainerNameShort
	| join kind=inner ContainerInventory on ContainerID
	| project Name, PodCreationTimeStamp, ContainerID, ContainerNameShort, ContainerLastKnownTimeStamp=coalesce(FinishedTime, TimeGenerated), CreatedTime
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp), ContainerCreationTimeStamp=min(CreatedTime), LastTimeGenerated=max(ContainerLastKnownTimeStamp) by Name, ContainerID, ContainerNameShort`,
		paramNamespace, paramPipelineJobName)
}

func getPipelineJobContainerLogQuery(joinContainerLog string) string {
	return fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and Name == %s and ContainerID == %s
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-job-name"] == %s and isempty(d["tekton.dev/task"])
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramPodName, paramContainerId, paramPipelineJobName)
}
