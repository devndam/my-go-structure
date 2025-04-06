package models

import "time"

// Migration struct to track applied migrations
type Migration struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
}
