package handler

import (
	"fmt"
	"net/http"

	"github.com/fabioaacarneiro/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context)  {
    id := ctx.Query("id")

    if id == "" {
        sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
        return
    }

    opening := schemas.Opening{}

    if err := db.First(&opening, id).Error; err != nil {
        sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id: %v is not found", id))
        return
    }

    if err := db.Delete(&opening).Error; err != nil {
        sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting openin with id %v", id))
    }

    sendSuccess(ctx, "delete-opening", opening)
}

