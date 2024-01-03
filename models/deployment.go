package models

import (
	"time"

	"github.com/google/uuid"
)

type Deployment struct {
	Id        uuid.UUID `json:"id"`
	Replicas  uint      `json:"replicas"`
	Image     string    `json:"image"`
	Labels    Label     `json:"labels"`
	Ports     []Port    `json:"ports"`
	CreatedAt time.Time `json:"createdAt"`
}

type Label struct {
	MonitorLabel  string `json:"monitor-label"`
	DeploymentTag string `json:"deployment-tag"`
}

type Port struct {
	Name string `json:"name"`
	Port uint   `json:"port"`
}
