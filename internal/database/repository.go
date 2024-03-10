package database

import "time"

type Repository struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Path      string `gorm:"uniqueIndex"`
	Paused    bool
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
