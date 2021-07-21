package domain

import "github.com/lib/pq"

type Name struct {
	Nconst              string   `json:"nconst"`
	Primary_name        string   `json:"primary_name"`
	Birth_year          string   `json:"birth_year"`
	Death_year          string   `json:"column:death_year"`
	Primary_professions []string `json:"column:primary_professions"`
	Known_for_titles    []string `json:"column:known_for_titles"`
}

type GormName struct {
	Nconst              string   `gorm:"primaryKey;column:nconst"`
	Primary_name        string   `gorm:"column:primary_name"`
	Birth_year          string
	Death_year          string
	Primary_professions pq.StringArray `gorm:"type:text[]"`
	Known_for_titles    pq.StringArray `gorm:"type:text[]"`
}

func (GormName) TableName() string {
	return "names"
}
