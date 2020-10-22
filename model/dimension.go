package model

type Dimension struct {
	DimensionId int    `json:"dimensionId" db:"dimension_id"`
	ResearchId  int    `json:"researchId" db:"research_id"`
	Name        string `json:"name" db:"name"`
	Variable    string `json:"variable" db:"variable"`
	Status      string `json:"status" db:"status"`
	CreatedAt   string `json:"-" db:"created_at"`
	UpdatedAt   string `json:"-" db:"updated_at"`
}
