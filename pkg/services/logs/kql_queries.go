package logs

import (
	"fmt"
)

const (
	paramNamespace     = "ParamNamespace"
	paramAppName       = "ParamAppName"
	paramComponentName = "ParamComponentName"
	paramPodName       = "ParamPodName"
	paramContainerId   = "ParamContainerId"
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

	componentInventory string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s and isnotempty(ContainerID) == true
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| extend s=parse_json(ContainerLastStatus)
	| project TimeGenerated, Name, ContainerID, PodCreationTimeStamp, ContainerCreationTimeStamp=coalesce(ContainerCreationTimeStamp,todatetime(parse_json(ContainerLastStatus)["startedAt"]))
	| where isnotnull(ContainerCreationTimeStamp)
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp), ContainerCreationTimeStamp=min(ContainerCreationTimeStamp), LastTimeGenerated=max(TimeGenerated) by Name, ContainerID`,
		paramNamespace, paramAppName, paramComponentName)
)
