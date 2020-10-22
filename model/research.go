package model

type Research struct {
	ResearchId    int     `json:"researchId" db:"research_id"`
	ResearcherId  int     `json:"researcherId" db:"researcher_id"`
	ExpertId      int     `json:"expertId" db:"expert_id"`
	Title         string  `json:"title" db:"title"`
	Speciality    string  `json:"speciality" db:"speciality"`
	Authors       string  `json:"authors" db:"authors"`
	Observation   string  `json:"observation" db:"observation"`
	AttachmentOne *string `json:"attachmentOne" db:"attachment_one"`
	AttachmentTwo *string `json:"attachmentTwo" db:"attachment_two"`
	Status        int     `json:"status" db:"status"`
	CreatedAt     string  `json:"-" db:"created_at"`
	UpdatedAt     string  `json:"updatedAt" db:"updated_at"`
}
