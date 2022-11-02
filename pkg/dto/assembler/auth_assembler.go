package assembler

import (
	"github.com/AdiKhoironHasan/matkul/internal/models"
	"github.com/AdiKhoironHasan/matkul/pkg/dto"
)

func ToLogin(d *dto.UserLoginReqDTO) *models.UserModels {
	return &models.UserModels{
		Email:    d.Email,
		Password: d.Password,
	}
}
