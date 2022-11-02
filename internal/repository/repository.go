package repository

import "github.com/AdiKhoironHasan/matkul/internal/models"

type Repository interface {
	Login(dataMahasiswa *models.UserModels) (bool, error)
}
