package params

type Replica struct {
	ReplicaName string `uri:"replicaName" binding:"required"`
}
