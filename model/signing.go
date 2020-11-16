package model

type Signing struct {
	SigningId int    `json:"signingId" db:"signing_id"`
	ExpertId  int    `json:"expertId" db:"expert_id"`
	Link      string `json:"link" db:"link"`
	CreatedAt string `json:"createdAt" db:"created_at"`
	UpdatedAt string `json:"-" db:"updated_at"`
}
