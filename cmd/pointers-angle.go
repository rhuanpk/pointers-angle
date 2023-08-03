package main

import (
	"github.com/rhuanpk/pointers-angle/internal/infra/api"
	"github.com/rhuanpk/pointers-angle/pkg/logger"

	_ "github.com/rhuanpk/pointers-angle/internal/process"
)

func main() {
	logger.Fatal.Fatalln(api.Router.Run(":8080").Error())
}
