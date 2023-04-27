package logs

import "fmt"

const (
	paramNamespace     = "ParamNamespace"
	paramAppName       = "ParamAppName"
	paramComponentName = "ParamComponentName"
)

var (
	componentLogQuery string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s
	| where ContainerID != ""
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| summarize by ContainerID
	| join kind=inner ContainerLog on $left.ContainerID==$right.ContainerID
	| project TimeGenerated, Name, ContainerID, LogEntry
	| sort by TimeGenerated asc`,
		paramNamespace, paramAppName, paramComponentName)

	componentPodInventory string = fmt.Sprintf(`KubePodInventory
	| where Namespace == %s
	| where ContainerID != ""
	| extend d=parse_json(PodLabel)[0]
	| where d["radix-app"] == %s and d["radix-component"] == %s and isempty(d["is-job-scheduler-pod"]) and isempty(d["radix-job-type"])
	| project Name, ContainerID, PodCreationTimeStamp, ContainerCreationTimeStamp=coalesce(ContainerCreationTimeStamp,todatetime(parse_json(ContainerLastStatus)["startedAt"]))
	| summarize PodCreationTimeStamp=min(PodCreationTimeStamp), ContainerCreationTimeStamp=min(ContainerCreationTimeStamp) by Name, ContainerID`,
		paramNamespace, paramAppName, paramComponentName)
)
