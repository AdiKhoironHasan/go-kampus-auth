package services

import (
	integ "github.com/AdiKhoironHasan/matkul/internal/integration"
	"github.com/AdiKhoironHasan/matkul/internal/repository"
	"github.com/AdiKhoironHasan/matkul/pkg/dto"
	"github.com/AdiKhoironHasan/matkul/pkg/dto/assembler"
)

type service struct {
	repo      repository.Repository
	IntegServ integ.IntegServices
}

func NewService(repo repository.Repository, IntegServ integ.IntegServices) Services {
	return &service{repo, IntegServ}
}

func (s *service) Login(req *dto.UserLoginReqDTO) error {

	dtLogin := assembler.ToLogin(req)

	err := s.repo.Login(dtLogin)
	if err != nil {
		return err
	}

	return nil
}
