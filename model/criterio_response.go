package model

type CriterioResponse struct {
	CriterioResponseId int    `json:"criterioResponseId" db:"criterio_response_id"`
	CriterioId         int    `json:"criterioId" db:"criterio_id"`
	ResearchId         int    `json:"researchId" db:"research_id"`
	DimensionId        int    `json:"dimensionId" db:"dimension_id"`
	Status             string `json:"status" db:"status"`
	CreatedAt          string `json:"-" db:"created_at"`
	UpdatedAt          string `json:"updateAt" db:"updated_at"`
}
