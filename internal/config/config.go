package config

import "github.com/rhuanpk/pointers-angle/internal/models"

var (
	// API is the global struct that store all API configuration.
	API *models.API

	// DB is the global struct that store all DB configuration.
	DB *models.DB
)
