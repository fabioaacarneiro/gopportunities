package handler

import (
	"net/http"

	"github.com/fabioaacarneiro/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context)  {
    openings := []schemas.Opening{}

    if err := db.Find(&openings).Error; err != nil {
        sendError(ctx, http.StatusInternalServerError, "error on listing openings")
        return
    }

    sendSuccess(ctx, "success on listring openings", openings)
}
