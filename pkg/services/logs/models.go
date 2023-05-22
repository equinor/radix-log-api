package logs

import "time"

type Pod struct {
	Name              string
	CreationTimestamp time.Time
	LastKnown         time.Time
	Containers        []Container
}

type Container struct {
	Id                string
	CreationTimestamp time.Time
	LastKnown         time.Time
}
