package database

import "gorm.io/gorm"

type Repository struct {
	gorm.Model
	Path   string `gorm:"uniqueIndex"`
	Paused bool
}

func GetRepositoryByAbsPath(path string) (Repository, error) {
	var repository Repository
	result := Client.First(&repository, "path = ?", path)
	return repository, result.Error
}

func ToggleRepositoryPaused(repository Repository) {
	repository.Paused = !repository.Paused
	Client.Save(&repository)
}
