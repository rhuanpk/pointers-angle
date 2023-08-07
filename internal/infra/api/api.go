package api

import (
	"github.com/rhuanpk/pointers-angle/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

var (
	// Router is the base router of api.
	Router = gin.Default()

	// V0 is the global route groups.
	V0 = Router.Group(
		"/v0/rest",
		middlewares.CORS,
	)
)
