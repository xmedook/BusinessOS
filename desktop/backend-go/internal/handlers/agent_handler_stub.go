package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rhl/businessos-backend/internal/config"
)

// AgentHandler handles agent-related endpoints (presets, custom agents).
type AgentHandler struct {
	pool *pgxpool.Pool
	cfg  *config.Config
}

// NewAgentHandler creates a new AgentHandler.
func NewAgentHandler(pool *pgxpool.Pool, cfg *config.Config) *AgentHandler {
	return &AgentHandler{pool: pool, cfg: cfg}
}

func (h *AgentHandler) ListAgentPresets(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"presets": []interface{}{}, "count": 0})
}

func (h *AgentHandler) GetAgentPreset(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Agent presets not available"})
}

func (h *AgentHandler) ListCustomAgents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"agents": []interface{}{}, "count": 0})
}

func (h *AgentHandler) CreateCustomAgent(c *gin.Context) {
	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Agent management not available"})
}

func (h *AgentHandler) TestCustomAgent(c *gin.Context) {
	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Agent testing not available"})
}

func (h *AgentHandler) ListCustomAgentsByCategory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"agents": []interface{}{}, "count": 0})
}

func (h *AgentHandler) CreateAgentFromPreset(c *gin.Context) {
	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Agent management not available"})
}

func (h *AgentHandler) GetCustomAgent(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
}

func (h *AgentHandler) UpdateCustomAgent(c *gin.Context) {
	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Agent management not available"})
}

func (h *AgentHandler) DeleteCustomAgent(c *gin.Context) {
	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Agent management not available"})
}
