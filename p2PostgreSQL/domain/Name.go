package domain

type Name struct {
	Nconst              string   `json:"nconst"`
	Primary_name        string   `json:"primary_name"`
	Birth_year          string   `json:"birth_year"`
	Death_year          string   `json:"death_year"`
	Primary_professions []string `json:"primary_professions"`
	Known_for_titles    []string `json:"known_for_titles"`
} 