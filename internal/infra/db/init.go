package db

import (
	"fmt"
	"time"

	"github.com/rhuanpk/pointers-angle/internal/config"
	"github.com/rhuanpk/pointers-angle/internal/models"
	"github.com/rhuanpk/pointers-angle/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	// This variable is necessary or in next "Tx, err :=", the Tx variable is locally created instead of use the global already existing.
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		*config.DB.Host,
		*config.DB.User,
		*config.DB.Pass,
		*config.DB.Name,
		*config.DB.Port,
		*config.DB.SSL,
		*config.DB.TZ,
	)

	Tx, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal.Fatalln(err.Error())
	}

	db, err := Tx.DB()
	if err != nil {
		logger.Fatal.Fatalln(err.Error())
	}
	db.SetMaxIdleConns(int(*config.DB.MaxIdleConns))
	db.SetMaxOpenConns(int(*config.DB.MaxOpenConns))
	db.SetConnMaxLifetime(time.Second * time.Duration(*config.DB.MaxLifeSecConn))

	if err := Tx.AutoMigrate(&models.Clock{}); err != nil {
		logger.Debug.Fatalln(err.Error())
	}
}
