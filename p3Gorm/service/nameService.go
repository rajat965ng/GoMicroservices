package service

import (
	"fmt"
	"github.com/rajat965ng/p3Gorm/domain"
	"github.com/rajat965ng/p3Gorm/repository"
)

type NameService struct {
	nameRepo *repository.NameRepository
}

func NewNameService() (*NameService,error)  {
	nr,err := repository.NewPostgreGorm()
	if err != nil {
		fmt.Printf("Error creating  Name repository: %s\n",err.Error())
		return nil,err
	}
	return &NameService{
		nameRepo:nr,
	}, nil
}

func (nameSvc *NameService) FindNameService(id string) (*domain.Name, error) {
	return nameSvc.nameRepo.FindNamesById(id)
}