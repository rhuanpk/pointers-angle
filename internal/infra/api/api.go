package api

import (
	"github.com/rhuanpk/pointers-angle/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

var (
	// Router api router base.
	Router = gin.Default()

	// V0 global route groups.
	V0 = Router.Group(
		"/v0/rest/",
		middlewares.CORS,
	)
)
