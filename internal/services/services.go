package services

import "github.com/AdiKhoironHasan/go-kampus-auth/pkg/dto"

type Services interface {
	Login(req *dto.UserLoginReqDTO) (*dto.UserResponse, error)
}
