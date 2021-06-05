package app

import (
	"github.com/mrityunjaygr8/go-pass/users"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (a App) migrate() {
	a.DB.AutoMigrate(&users.User{})
}

func (a App) Initialize() {
	a.migrate()
}
