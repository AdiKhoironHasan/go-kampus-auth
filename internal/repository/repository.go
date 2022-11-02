package repository

import "github.com/AdiKhoironHasan/go-kampus-auth/internal/models"

type Repository interface {
	Login(dataMahasiswa *models.UserModels) ([]*models.UserModels, error)
}
