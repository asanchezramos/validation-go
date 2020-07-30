package model

type Answer struct {
	AnswerId     int     `json:"answerId" db:"answer_id"`
	Comments     string  `json:"comments" db:"comments"`
	File         *string `json:"file" db:"file"`
	SolicitudeId int     `json:"solicitudeId" db:"solicitude_id"`
	Status       int     `json:"status" db:"status"`
	CreatedAt    string  `json:"-" db:"created_at"`
	UpdatedAt    string  `json:"-" db:"updated_at"`
}

type ExpertAnswer struct {
	AnswerId  int     `json:"answerId" db:"answer_id"`
	FullName  string  `json:"fullName" db:"full_name"`
	Photo     *string `json:"photo" db:"photo"`
	Specialty *string `json:"specialty" db:"specialty"`
	File      *string `json:"file" db:"file"`
	Comments  string  `json:"comments" db:"comments"`
}
