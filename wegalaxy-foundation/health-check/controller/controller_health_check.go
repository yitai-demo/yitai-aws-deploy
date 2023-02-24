package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController struct {
}

func NewHealthCheckController() (*HealthCheckController, error) {
	return &HealthCheckController{}, nil
}

// Health Check godoc
// @Summary      Health Check
// @Description  Health Check
// @ID           health-check
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /health-check [get]
func (c *HealthCheckController) HealthCheck(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusOK)
}
