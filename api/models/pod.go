package models

import "time"

type Replica struct {
	Name              string      `json:"name" example:"web-7db5f9c99b-nwn2w"`
	CreationTimestamp time.Time   `json:"creationTimestamp,omitempty" example:"2023-01-31T08:00:00Z"`
	LastKnown         time.Time   `json:"lastKnown,omitempty" example:"2023-01-31T08:00:00Z"`
	Containers        []Container `json:"containers"`
}

type Container struct {
	Id                string    `json:"id" example:"d40ba550f05b252da60e5b873c32204368ecb5b9e17ea123738d318be4e5295b"`
	Name              string    `json:"name,omitempty" example:"web"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty" example:"2023-01-31T08:00:00Z"`
	LastKnown         time.Time `json:"lastKnown,omitempty" example:"2023-01-31T08:00:00Z"`
}
