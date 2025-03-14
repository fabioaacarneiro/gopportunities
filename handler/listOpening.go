package handler

import (
	"net/http"

	"github.com/fabioaacarneiro/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context)  {
    openings := []schemas.Opening{}

    if err := db.Find(&openings).Error; err != nil {
        sendError(ctx, http.StatusInternalServerError, "error on listing openings")
        return
    }

    sendSuccess(ctx, "success on listring openings", openings)
}
