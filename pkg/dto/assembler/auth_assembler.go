package assembler

import (
	"github.com/AdiKhoironHasan/go-kampus-auth/internal/models"
	"github.com/AdiKhoironHasan/go-kampus-auth/pkg/dto"
)

func ToLogin(d *dto.UserLoginReqDTO) *models.UserModels {
	return &models.UserModels{
		Email:    d.Email,
		Password: d.Password,
	}
}
