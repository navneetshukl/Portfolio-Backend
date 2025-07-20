package handlers

import (
	"context"
	"net/http"
	"portfolio/internals/core"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	emailHandler core.EmailUseCase
}

func NewHandler(eh core.EmailUseCase) Handler {
	return Handler{
		emailHandler: eh,
	}
}

func (h *Handler) SendEmailHandler() func(*gin.Context) {
	return func(ctx *gin.Context) {

		var req *core.SendEmail
		err := ctx.ShouldBindJSON(&req)

		if err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":    []interface{}{},
				"message": "invalid request",
				"status":  http.StatusBadRequest,
			})
			return
		}

		err = h.emailHandler.SendEmail(context.Background(), req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":    []interface{}{},
				"message": "something went wrong",
				"status":  http.StatusInternalServerError,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data":    []interface{}{},
			"message": "email sent successfully",
			"status":  http.StatusOK,
		})
		return

	}
}
