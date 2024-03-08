package database

import "gorm.io/gorm"

type Repository struct {
	gorm.Model
	Path   string `gorm:"uniqueIndex"`
	Paused bool
}
