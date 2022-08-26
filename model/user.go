package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserId    int     `json:"userId" db:"user_id"`
	Name      string  `json:"name" db:"name"`
	FullName  string  `json:"fullName" db:"full_name"`
	Photo     *string `json:"photo" db:"photo"`
	Mail      string  `json:"mail" db:"mail"`
	Password  *string `json:"password,omitempty" db:"password"`
	Phone     string  `json:"phone" db:"phone"`
	Specialty *string `json:"specialty, omitempty" db:"specialty"`
	Role      string  `json:"role" db:"role"`
	Status    int     `json:"status" db:"status"`
	CreatedAt string  `json:"-" db:"created_at"`
	UpdatedAt string  `json:"-" db:"updated_at"`
	Orcid  	  string  `json:"orcid" db:"orcid"`
}

type Auth struct {
	Mail     string `json:"mail" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
