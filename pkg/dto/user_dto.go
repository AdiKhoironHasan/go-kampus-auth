package dto

import (
	"github.com/AdiKhoironHasan/matkul/pkg/common/validator"
)

type UserLoginReqDTO struct {
	Email    string `json:"email" valid:"required" validname:"email"`
	Password string `json:"password" valid:"required" validname:"password"`
}

func (dto *UserLoginReqDTO) Validate() error {
	v := validator.NewValidate(dto)

	return v.Validate()
}
