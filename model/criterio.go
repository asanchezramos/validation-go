package model

type Criterio struct {
	CriterioId int    `json:"criterioId" db:"criterio_id"`
	Name       string `json:"name" db:"name"`
	Speciality string `json:"speciality" db:"speciality"`
	ExpertId   int    `json:"expertId" db:"expert_id"`
	CreatedAt  string `json:"-" db:"created_at"`
	UpdatedAt  string `json:"-" db:"updated_at"`
}
