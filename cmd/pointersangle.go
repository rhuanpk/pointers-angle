package main

import (
	// Built-in imports.
	"fmt"

	// Project imports.
	"github.com/rhuanpk/pointers-angle/internal/config"
	"github.com/rhuanpk/pointers-angle/internal/infra/api"
	"github.com/rhuanpk/pointers-angle/pkg/logger"

	// Load imports.
	_ "github.com/rhuanpk/pointers-angle/internal/process"
)

// Swagger:
//
//	@title			Pointer Angles API
//	@version		0.0.0
//	@description	API for calculate the less angle between hour and minute pointer.
//
//	@contact.name	Rhuan Patriky
//	@contact.url	https://linktr.ee/rhuanpk
//	@contact.email	support@rhuanpk.com
//
//	@host			localhost:8080
//	@BasePath		/v0/rest
func main() {
	logger.Fatal.Fatalln(api.Router.Run(fmt.Sprintf(":%d", *config.API.Port)).Error())
}
