package handler

import (
	"github.com/Felipe-Janazi/gooportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func InitializerHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}