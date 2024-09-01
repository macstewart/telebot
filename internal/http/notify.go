package http

import (
	"github.com/gin-gonic/gin"
	"github.com/macstewart/telebot/internal/http/dto"
	"github.com/macstewart/telebot/internal/telegram"
	"golang.org/x/exp/slog"
)

func NotifyHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.NotifyRequest
		if err := ctx.Bind(&req); err != nil {
			slog.Error("Failed to bind request", "error", err)
			return
		}
		slog.Info("Go request notify", "data", req)
		telegram.Notify(req.Message)
	}
}
