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
	paramPodName          = "ParamPodName"
	paramContainerId      = "ParamContainerId"
)

// KQL documentation: https://learn.microsoft.com/en-us/azure/data-explorer/kusto/query/

var (
	joinContainerLog = `| join kind=inner ContainerLog on $left.ContainerID==$right.ContainerID
	| project TimeGenerated, Name, ContainerID, LogEntry
	| sort by TimeGenerated desc`
)

var (
	componentLogQuery string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramAppName, paramComponentName)

	componentPodLogQuery string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and Name == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramPodName, paramAppName, paramComponentName)

	componentContainerLogQuery string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and Name == %s and ContainerID == %s
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramPodName, paramContainerId, paramAppName, paramComponentName)

	componentInventoryQuery string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| project TimeGenerated, Name, ContainerID, PodCreationTimeStamp, ContainerCreationTimeStamp=coalesce(ContainerCreationTimeStamp,todatetime(parse_json(ContainerLastStatus)["startedAt"]))
	| where isnotnull(ContainerCreationTimeStamp)
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp), ContainerCreationTimeStamp=min(ContainerCreationTimeStamp), LastTimeGenerated=max(TimeGenerated) by Name, ContainerID`,
		paramNamespace, paramAppName, paramComponentName)

	jobInventoryQuery string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s 
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp) by Name, ContainerID
    | join kind=inner ContainerInventory on ContainerID 
    | project Name, PodCreationTimeStamp, ContainerID, ContainerLastKnownTimeStamp=coalesce(FinishedTime, TimeGenerated), CreatedTime
    | summarize PodCreationTimeStamp=min(PodCreationTimeStamp), ContainerCreationTimeStamp=min(CreatedTime), LastTimeGenerated=max(ContainerLastKnownTimeStamp) by Name, ContainerID`,
		paramNamespace, paramAppName, paramJobComponentName, paramJobName)

	jobLogQuery string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s 
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramAppName, paramJobComponentName, paramJobName)

	jobPodLogQuery string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and Name == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s 
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramPodName, paramAppName, paramJobComponentName, paramJobName)

	jobContainerLogQuery string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and Name == %s and ContainerID == %s
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and d["radix-job-type"] == "job-scheduler" and d["job-name"] == %s 
	| summarize by ContainerID
	`+joinContainerLog,
		paramNamespace, paramPodName, paramContainerId, paramAppName, paramJobComponentName, paramJobName)
)
