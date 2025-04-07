package database

import "github.com/devndam/go-starter/app/models"

var MigratableModels = []interface {
}{
	&models.User{},
}
