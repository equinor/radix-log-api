package models

import "time"

type Pod struct {
	Name              string      `json:"name"`
	CreationTimestamp time.Time   `json:"creationTimestamp,omitempty"`
	Containers        []Container `json:"containers"`
}

type Container struct {
	Id                string    `json:"id"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}
