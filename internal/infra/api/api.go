package api

import (
	"github.com/rhuanpk/pointers-angle/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

var (
	// Router is the base router of api.
	Router = gin.Default()

	// V1 is the global route groups.
	V1 = Router.Group(
		"/v1/rest",
		middlewares.CORS,
	)
)
