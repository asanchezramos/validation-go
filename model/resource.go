package model

type ResourceUser struct {
	ResourceUserId int    `json:"resourceUserId" db:"resource_user_id"`
	ExpertId       int    `json:"expertId" db:"expert_id"`
	Title          string `json:"title" db:"title"`
	Subtitle       string `json:"subtitle" db:"subtitle"`
	Link           string `json:"link" db:"link"`
	CreatedAt      string `json:"-" db:"created_at"`
	UpdatedAt      string `json:"updatedAt" db:"updated_at"`
}
