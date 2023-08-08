#!/bin/sh

# Script for API run standalone.

swag fmt && swag init -g ./cmd/pointersangle.go -o ./docs/
go run -ldflags '-s -w' ./cmd/
