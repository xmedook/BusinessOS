// Package sorx provides stub types for the SORX skill execution engine.
package sorx

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rhl/businessos-backend/internal/carrier"
)

// Engine is the SORX skill execution engine.
type Engine struct {
	pool   *pgxpool.Pool
	logger *slog.Logger
}

// NewEngine creates a new no-op SORX engine.
func NewEngine(pool *pgxpool.Pool, logger *slog.Logger) *Engine {
	return &Engine{pool: pool, logger: logger}
}

// SetCarrierClient attaches a carrier client to the engine (no-op stub).
func (e *Engine) SetCarrierClient(_ *carrier.Client) {}

// ExecuteSkill starts a skill execution (stub returns disabled status).
func (e *Engine) ExecuteSkill(_ context.Context, req ExecuteRequest) (*Execution, error) {
	return &Execution{
		ID:        uuid.New(),
		SkillID:   req.SkillID,
		Status:    "disabled",
		StartedAt: time.Now(),
	}, nil
}

// GetExecution retrieves a skill execution by ID (stub always returns not-found).
func (e *Engine) GetExecution(_ uuid.UUID) (*Execution, bool) {
	return nil, false
}

// ListSkills returns all available skill definitions (stub returns empty).
func (e *Engine) ListSkills() []*SkillDefinition {
	return nil
}

// ExecuteAction runs a named action with params (used by carrier/optimal).
func (e *Engine) ExecuteAction(_ context.Context, _ string, _ map[string]interface{}) (any, error) {
	return nil, nil
}

// ExecuteRequest is the input for ExecuteSkill.
type ExecuteRequest struct {
	SkillID string
	UserID  string
	Params  map[string]interface{}
}

// Execution tracks the state of a running skill.
type Execution struct {
	ID          uuid.UUID
	SkillID     string
	Status      string
	CurrentStep string
	Params      map[string]interface{}
	Result      interface{}
	Error       string
	StepResults interface{}
	StartedAt   time.Time
	CompletedAt *time.Time
}

// SkillDefinition describes an available skill.
type SkillDefinition struct {
	ID                   string
	Name                 string
	Description          string
	Category             string
	RequiredIntegrations []string
	Steps                []interface{}
}

// SkillCommand maps a chat command to a skill.
type SkillCommand struct {
	Name        string
	DisplayName string
	Description string
	Icon        string
	Category    string
	SkillID     string
	Params      map[string]interface{}
}

// ListSkillCommands returns all skill-based commands (stub returns empty).
func ListSkillCommands() []SkillCommand {
	return nil
}

// GetSkillCommand retrieves a skill command by name (stub always returns not-found).
func GetSkillCommand(_ string) (SkillCommand, bool) {
	return SkillCommand{}, false
}

// Scheduler periodically triggers proactive skill executions.
type Scheduler struct{}

// NewScheduler creates a new no-op scheduler.
func NewScheduler(_ *Engine, _ *pgxpool.Pool, _ *slog.Logger) *Scheduler {
	return &Scheduler{}
}

// Start begins the scheduler (no-op).
func (s *Scheduler) Start() error {
	return nil
}

// Stop halts the scheduler (no-op).
func (s *Scheduler) Stop() {}
