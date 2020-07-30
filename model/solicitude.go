package model

type Solicitude struct {
	SolicitudeId  int     `json:"solicitudeId" db:"solicitude_id"`
	Repository    string  `json:"repository" db:"repository"`
	Investigation *string `json:"investigation" db:"investigation"`
	UserId        int     `json:"userId" db:"user_id"`
	ExpertId      int     `json:"expertId" db:"expert_id"`
	Status        string  `json:"status" db:"status"`
	CreatedAt     string  `json:"-" db:"created_at"`
	UpdatedAt     string  `json:"-" db:"updated_at"`
}

type SolicitudeUser struct {
	SolicitudeId int     `json:"solicitudeId" db:"solicitude_id"`
	FullName     string  `json:"fullName" db:"full_name"`
	Specialty    *string `json:"specialty,omitempty" db:"specialty"`
	Status       string  `json:"status" db:"status"`
}

type UserSolicitude struct {
	SolicitudeId int     `json:"solicitudeId" db:"solicitude_id"`
	UserId       int     `json:"userId" db:"user_id"`
	Name         string  `json:"name" db:"name"`
	FullName     string  `json:"fullName" db:"full_name"`
	Photo        *string `json:"photo" db:"photo"`
	Mail         string  `json:"mail" db:"mail"`
	Password     *string `json:"password,omitempty" db:"password"`
	Phone        string  `json:"phone" db:"phone"`
	Specialty    *string `json:"specialty" db:"specialty"`
	Role         string  `json:"role" db:"role"`
	Status       int     `json:"status" db:"status"`
	CreatedAt    string  `json:"-" db:"created_at"`
	UpdatedAt    string  `json:"-" db:"updated_at"`
}

type StudentSolicitude struct {
	SolicitudeId  int     `json:"solicitudeId" db:"solicitude_id"`
	Repository    *string  `json:"repository" db:"repository"`
	Investigation *string `json:"investigation" db:"investigation"`
	FullName      string  `json:"fullName" db:"full_name"`
	Photo         *string `json:"photo" db:"photo"`
}
