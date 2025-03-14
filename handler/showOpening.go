package handler

import (
	"fmt"
	"net/http"

	"github.com/fabioaacarneiro/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a opening by id
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context)  {
    id := ctx.Query("id")

    if id == "" {
        sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
        return
    }

    opening := schemas.Opening{}

    if err := db.First(&opening, id).Error; err != nil {
        sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v is not found", id))
        return
    }

    sendSuccess(ctx, "success on search opening by id", opening)
}

