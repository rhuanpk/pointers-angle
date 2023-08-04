package db

import "gorm.io/gorm"

// Tx is the global database connection.
var Tx *gorm.DB
