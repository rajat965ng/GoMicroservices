package repository

import (
	"fmt"
	"github.com/rajat965ng/p3Gorm/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type NameRepository struct {
	db *gorm.DB
}

func NewPostgreGorm() (*NameRepository, error) {
	db, err := gorm.Open(postgres.Open("postgres://user:password@localhost:5432/dbname?sslmode=disable"), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error in creating db connection: %s\n", err.Error())
		return nil, err
	}
	return &NameRepository{
		db: db,
	}, nil
}

func (pg *NameRepository) FindNamesById(nconst string) (*domain.Name, error) {
	name := &domain.GormName{}
	fmt.Printf("Nconst is %s\n",nconst)
	tx := pg.db.Where("nconst = ?", nconst).First(name)
	if tx.Error != nil {
		fmt.Printf("Error while FindNamesById: %s\n", tx.Error)
		return nil, tx.Error
	}

	return &domain.Name{
		Nconst:              name.Nconst,
		Primary_name:        name.Primary_name,
		Birth_year:          name.Birth_year,
		Death_year:          name.Death_year,
		Primary_professions: name.Primary_professions,
		Known_for_titles:    name.Known_for_titles,
	}, nil
}
