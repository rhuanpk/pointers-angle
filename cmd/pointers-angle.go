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

func main() {
	logger.Fatal.Fatalln(api.Router.Run(fmt.Sprintf(":%d", *config.Config.API.Port)).Error())
}
