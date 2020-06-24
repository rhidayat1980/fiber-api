package database

import (
	"github.com/jinzhu/gorm"

	// dialect for sql
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// DBConn variable
	DBConn *gorm.DB
)
