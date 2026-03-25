// Package appgen provides stub types for the multi-agent app generation subsystem.
package appgen

import "time"

// AgentType identifies a specialist code-generation agent.
type AgentType string

const (
	AgentFrontend AgentType = "frontend"
	AgentBackend  AgentType = "backend"
	AgentDatabase AgentType = "database"
	AgentTest     AgentType = "test"
)

// ProgressEvent represents progress from an agent during generation.
type ProgressEvent struct {
	AgentType AgentType
	Progress  int
	Message   string
	TaskID    string
	Status    string
	Timestamp time.Time
}

// Plan represents an app generation plan.
type Plan struct {
	Tasks []PlanTask
}

// PlanTask is a single task within a plan.
type PlanTask struct {
	Type        AgentType
	Description string
}

// GeneratedApp is the result of a multi-agent app generation run.
type GeneratedApp struct {
	Success bool
	Results []AgentResult
}

// AgentResult holds the output of a single agent.
type AgentResult struct {
	AgentType  AgentType
	CodeBlocks map[string]string
}
