package services

import "github.com/AdiKhoironHasan/matkul/pkg/dto"

type Services interface {
	Login(req *dto.UserLoginReqDTO) error
}
