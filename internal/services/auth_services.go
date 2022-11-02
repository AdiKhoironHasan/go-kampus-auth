package services

import (
	"fmt"

	integ "github.com/AdiKhoironHasan/go-kampus-auth/internal/integration"
	"github.com/AdiKhoironHasan/go-kampus-auth/internal/repository"
	"github.com/AdiKhoironHasan/go-kampus-auth/pkg/dto"
	"github.com/AdiKhoironHasan/go-kampus-auth/pkg/dto/assembler"
)

type service struct {
	repo      repository.Repository
	IntegServ integ.IntegServices
}

func NewService(repo repository.Repository, IntegServ integ.IntegServices) Services {
	return &service{repo, IntegServ}
}

func (s *service) Login(req *dto.UserLoginReqDTO) (*dto.UserResponse, error) {
	var dataRes *dto.UserResponse
	dtLogin := assembler.ToLogin(req)

	dataUser, err := s.repo.Login(dtLogin)

	if err != nil {
		return nil, err
	}

	if dataUser == nil {
		return nil, nil
	}

	dataRes = &dto.UserResponse{
		ID:   dataUser[0].ID,
		Name: dataUser[0].Name,
	}

	fmt.Println(*dataRes)
	fmt.Println(dataUser[0].ID)

	return dataRes, nil
}
