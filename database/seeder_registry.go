package database

import (
	"gorm.io/gorm"
)

// SeederFunc defines the signature for seeders
type SeederFunc func(*gorm.DB) error

var SeederList = []SeederFunc{}
