package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS set CORS headers in request and also validade preflight request.
func CORS(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
	if ctx.Request.Method == http.MethodOptions {
		ctx.Status(http.StatusContinue)
	}
	ctx.Next()
}
