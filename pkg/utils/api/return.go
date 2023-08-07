package api

import (
	"github.com/rhuanpk/pointers-angle/internal/models"
	"github.com/rhuanpk/pointers-angle/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Return is the standard return of API and system log.
func Return(ctx *gin.Context, code int, message string, optionalLogs ...string) {
	ctx.JSON(
		code,
		&models.HTTP{
			Code:  code,
			Error: message,
		},
	)
	ctx.Abort()
	for _, log := range optionalLogs {
		logger.Error.Println(log)
	}
}
