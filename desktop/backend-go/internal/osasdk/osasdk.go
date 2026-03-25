// Package osasdk provides stub types that replace the proprietary github.com/Miosa-osa/sdk-go SDK.
// All types are zero-value safe. Factory functions return errors when called at runtime
// since OSA_ENABLED defaults to false.
package osasdk

import (
	"context"
	"fmt"
	"time"
)

// Event types (string constants matching the original SDK).
const (
	EventThinking       = "thinking"
	EventResponse       = "response"
	EventSkillStarted   = "skill_started"
	EventSkillCompleted = "skill_completed"
	EventSkillFailed    = "skill_failed"
	EventError          = "error"
	EventConnected      = "connected"
	EventSignal         = "signal"
)

// Event represents a real-time event from the OSA stream.
type Event struct {
	Type string
	Data map[string]interface{}
}

// Client is the interface satisfied by both local and cloud SDK clients.
type Client interface {
	GenerateApp(ctx context.Context, req AppGenerationRequest) (*AppGenerationResponse, error)
	GetAppStatus(ctx context.Context, appID string) (*AppStatusResponse, error)
	Orchestrate(ctx context.Context, req OrchestrateRequest) (*OrchestrateResponse, error)
	GetWorkspaces(ctx context.Context) (*WorkspacesResponse, error)
	Health(ctx context.Context) (*HealthResponse, error)
	GenerateAppFromTemplate(ctx context.Context, req GenerateFromTemplateRequest) (*AppGenerationResponse, error)
	Stream(ctx context.Context, sessionID string) (<-chan Event, error)
	LaunchSwarm(ctx context.Context, req SwarmRequest) (*SwarmResponse, error)
	ListSwarms(ctx context.Context) ([]SwarmStatus, error)
	GetSwarm(ctx context.Context, swarmID string) (*SwarmStatus, error)
	CancelSwarm(ctx context.Context, swarmID string) error
	DispatchInstruction(ctx context.Context, agentID string, instruction Instruction) error
	ListTools(ctx context.Context) ([]ToolDefinition, error)
	ExecuteTool(ctx context.Context, toolName string, params map[string]interface{}) (*ToolResult, error)
	Close() error
}

// LocalConfig configures a local-mode SDK client.
type LocalConfig struct {
	BaseURL      string
	SharedSecret string
	Timeout      time.Duration
	Resilience   *ResilienceConfig
}

// CloudConfig configures a cloud-mode SDK client.
type CloudConfig struct {
	APIKey  string
	BaseURL string
	Timeout time.Duration
}

// ResilienceConfig configures SDK-level resilience (circuit breaker, retry).
type ResilienceConfig struct {
	Enabled bool
}

// NewLocalClient returns an error since the SDK is not available.
func NewLocalClient(_ LocalConfig) (Client, error) {
	return nil, fmt.Errorf("OSA SDK not available: MIOSA proprietary dependency removed")
}

// NewCloudClient returns an error since the SDK is not available.
func NewCloudClient(_ CloudConfig) (Client, error) {
	return nil, fmt.Errorf("OSA SDK not available: MIOSA proprietary dependency removed")
}

// APIError is an error type returned by the SDK.
type APIError struct {
	StatusCode int
	ErrorCode  string
	Details    string
}

func (e *APIError) Error() string {
	if e.Details != "" {
		return e.Details
	}
	return e.ErrorCode
}

// Request/response types

type AppGenerationRequest struct {
	UserID      string
	WorkspaceID string
	Name        string
	Description string
	Type        string
	Parameters  map[string]interface{}
}

type AppGenerationResponse struct {
	AppID       string
	Status      string
	WorkspaceID string
	Message     string
	Data        map[string]interface{}
	CreatedAt   string
}

type AppStatusResponse struct {
	AppID       string
	Status      string
	Progress    float64
	CurrentStep string
	Output      string
	Error       string
	Metadata    map[string]interface{}
	UpdatedAt   string
}

type OrchestrateRequest struct {
	UserID      string
	Input       string
	SessionID   string
	Phase       string
	Context     map[string]interface{}
	WorkspaceID string
}

type OrchestrateResponse struct {
	Success     bool
	Output      string
	AgentsUsed  []string
	ExecutionMS int64
	Metadata    map[string]interface{}
	NextStep    string
	SessionID   string
}

type WorkspacesResponse struct {
	Workspaces []WorkspaceInfo
	Total      int
}

type WorkspaceInfo struct {
	ID          string
	Name        string
	Description string
	OwnerID     string
	CreatedAt   string
	UpdatedAt   string
}

type HealthResponse struct {
	Status   string
	Version  string
	Provider string
}

type GenerateFromTemplateRequest struct {
	TemplateName string
	Variables    map[string]interface{}
	UserID       string
	WorkspaceID  string
}

type SwarmRequest struct {
	Pattern   string
	Task      string
	Config    map[string]interface{}
	MaxAgents int
	SessionID string
}

type SwarmResponse struct {
	SwarmID string
	Status  string
}

type SwarmStatus struct {
	SwarmID string
	Status  string
	Agents  []string
}

type Instruction struct {
	SpecVersion string
	Type        string
	Source      string
	ID          string
	Data        map[string]interface{}
}

type ToolDefinition struct {
	Name        string
	Description string
	Parameters  map[string]interface{}
}

type ToolResult struct {
	Output string
	Error  string
}
