package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// AgenticRAGResponse holds the result of an agentic RAG query.
type AgenticRAGResponse struct {
	Answer  string   `json:"answer"`
	Sources []string `json:"sources,omitempty"`
}

// AgenticRAGRequest is a request for agentic RAG search.
type AgenticRAGRequest struct {
	Query              string     `json:"query"`
	WorkspaceID        string     `json:"workspace_id,omitempty"`
	UserID             string     `json:"user_id,omitempty"`
	MaxResults         int        `json:"max_results,omitempty"`
	MinQualityScore    float64    `json:"min_quality_score,omitempty"`
	ProjectContext     *uuid.UUID `json:"project_context,omitempty"`
	TaskContext        *uuid.UUID `json:"task_context,omitempty"`
	UsePersonalization bool       `json:"use_personalization,omitempty"`
}

// QueryIntent classifies the intent behind a user query.
type QueryIntent string

const (
	IntentFactualLookup  QueryIntent = "factual_lookup"
	IntentProcedural     QueryIntent = "procedural"
	IntentAnalytical     QueryIntent = "analytical"
	IntentConversational QueryIntent = "conversational"
	IntentComparison     QueryIntent = "comparison"
	IntentRecent         QueryIntent = "recent"
	IntentExhaustive     QueryIntent = "exhaustive"
)

// AgenticRAGService provides intelligent, adaptive RAG retrieval.
type AgenticRAGService struct{}

// NewAgenticRAGService creates a new agentic RAG service (stub).
func NewAgenticRAGService(_ *pgxpool.Pool, _ *HybridSearchService, _ *ReRankerService, _ *EmbeddingService, _ *LearningService) *AgenticRAGService {
	return &AgenticRAGService{}
}

// Query performs an agentic RAG search (stub).
func (s *AgenticRAGService) Query(_ context.Context, _ AgenticRAGRequest) (*AgenticRAGResponse, error) {
	return &AgenticRAGResponse{}, nil
}

// Retrieve performs an agentic RAG retrieval (stub).
func (s *AgenticRAGService) Retrieve(_ context.Context, _ AgenticRAGRequest) (*AgenticRAGResponse, error) {
	return &AgenticRAGResponse{}, nil
}

// SetCache sets the RAG cache on the service (stub).
func (s *AgenticRAGService) SetCache(_ *RAGCacheService) {}

// SetQueryExpansion sets the query expansion service (stub).
func (s *AgenticRAGService) SetQueryExpansion(_ *QueryExpansionService) {}
