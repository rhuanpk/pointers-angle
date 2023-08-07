package clock

import "github.com/rhuanpk/pointers-angle/internal/infra/api"

func init() {
	clock := api.V0.Group("/clock")
	{
		clock.GET("/:hour", angle)
		clock.GET("/:hour/:minute", angle)
	}
}
