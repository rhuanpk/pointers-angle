package clock

import "github.com/rhuanpk/pointers-angle/internal/infra/api"

func init() {
	clock := api.V1.Group("/clock")
	{
		clock.GET("/:hour", angle)
		clock.GET("/:hour/:minute", angle)
	}
}
