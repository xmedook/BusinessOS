// Package agents provides stub types for the agent subsystem.
// All implementations are no-ops so the server compiles with OSA_ENABLED=false.
package agents

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rhl/businessos-backend/internal/config"
	"github.com/rhl/businessos-backend/internal/feedback"
	"github.com/rhl/businessos-backend/internal/services"
	bossignal "github.com/rhl/businessos-backend/internal/signal"
	"github.com/rhl/businessos-backend/internal/streaming"
	"github.com/rhl/businessos-backend/internal/tools"
)

// AgentType identifies a specialist agent.
type AgentType string

const (
	AgentTypeDocument     AgentType = "document"
	AgentTypeAnalyst      AgentType = "analyst"
	AgentTypeTask         AgentType = "task"
	AgentTypeProject      AgentType = "project"
	AgentTypeClient       AgentType = "client"
	AgentTypeOrchestrator AgentType = "orchestrator"
)

// AgentInput is the input payload for Agent.Run.
type AgentInput struct {
	Messages       []services.ChatMessage
	Context        *services.TieredContext
	FocusMode      string
	ConversationID uuid.UUID
	UserID         string
	UserName       string
	MemoryContext  string
	RoleContext    string
	SignalEnvelope *bossignal.SignalEnvelope
}

// Agent is the interface every agent must satisfy.
type Agent interface {
	Run(ctx context.Context, input AgentInput) (<-chan streaming.StreamEvent, <-chan error)
	GetSystemPrompt() string
	SetModel(model string)
	SetOptions(opts services.LLMOptions)
	SetRoleContextPrompt(prompt string)
	SetProfileContext(ctx string)
	SetMemoryContext(ctx string)
	SetSkillsPrompt(prompt string)
	SetFocusModePrompt(prompt string)
	SetCustomSystemPrompt(prompt string)
	SetOutputStylePrompt(prompt string)
	SetGenreContext(ctx string)
	SetTieredContext(ctx *services.TieredContext)
	RegisterExternalTool(tool *tools.EscalateToOSATool)
}

// BaseAgent is a concrete type used in type assertions by handler code.
type BaseAgent struct{}

func (b *BaseAgent) Run(_ context.Context, _ AgentInput) (<-chan streaming.StreamEvent, <-chan error) {
	events := make(chan streaming.StreamEvent)
	errs := make(chan error)
	close(events)
	close(errs)
	return events, errs
}

func (b *BaseAgent) GetSystemPrompt() string                          { return "" }
func (b *BaseAgent) SetModel(_ string)                                {}
func (b *BaseAgent) SetOptions(_ services.LLMOptions)                 {}
func (b *BaseAgent) SetRoleContextPrompt(_ string)                    {}
func (b *BaseAgent) SetProfileContext(_ string)                       {}
func (b *BaseAgent) SetMemoryContext(_ string)                        {}
func (b *BaseAgent) SetSkillsPrompt(_ string)                        {}
func (b *BaseAgent) SetFocusModePrompt(_ string)                     {}
func (b *BaseAgent) SetCustomSystemPrompt(_ string)                  {}
func (b *BaseAgent) SetOutputStylePrompt(_ string)                   {}
func (b *BaseAgent) SetGenreContext(_ string)                        {}
func (b *BaseAgent) SetTieredContext(_ *services.TieredContext)       {}
func (b *BaseAgent) RegisterExternalTool(_ *tools.EscalateToOSATool) {}

// AgentRegistry creates and returns agents by type.
type AgentRegistry struct{}

// NewAgentRegistry returns a no-op registry.
func NewAgentRegistry(_ *pgxpool.Pool, _ *config.Config, _ *services.EmbeddingService, _ *services.PromptPersonalizer, _ feedback.SignalHintProvider) *AgentRegistry {
	return &AgentRegistry{}
}

// GetAgent returns a stub BaseAgent regardless of the requested type.
func (r *AgentRegistry) GetAgent(_ AgentType, _ string, _ string, _ *uuid.UUID, _ *services.TieredContext) Agent {
	return &BaseAgent{}
}

// OrchestratorCOT is a chain-of-thought orchestrator stub.
type OrchestratorCOT struct{}

// NewOrchestratorCOT returns a no-op COT orchestrator.
func NewOrchestratorCOT(_ *pgxpool.Pool, _ *config.Config, _ *AgentRegistry) *OrchestratorCOT {
	return &OrchestratorCOT{}
}

// ProcessWithCOT returns closed channels (no-op).
func (o *OrchestratorCOT) ProcessWithCOT(_ context.Context, _ AgentInput, _ string, _ string, _ *uuid.UUID, _ services.LLMOptions) (<-chan streaming.StreamEvent, <-chan error, error) {
	events := make(chan streaming.StreamEvent)
	errs := make(chan error)
	close(events)
	close(errs)
	return events, errs, nil
}

// BuildSignalAnnotation returns a genre-aware annotation string for prompt injection.
func BuildSignalAnnotation(_ *bossignal.SignalEnvelope, _ *services.TieredContext) string {
	return ""
}
