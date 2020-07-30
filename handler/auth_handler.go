package handler

import (
	"fmt"
	"github.com/goandreus/validation-go/datalayer/datastore"
	"github.com/goandreus/validation-go/model"
	"github.com/goandreus/validation-go/pkg/encrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// MobileHandler  represent the httphandler for auth endpoints
type AuthHandler struct {
	db datastore.Database
}

// NewAuthHandler will initialize the auth/ resources endpoint
func NewAuthHandler(e *echo.Echo, db datastore.Database) {
	handler := &AuthHandler{
		db,
	}
	g := e.Group("/auth")
	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
}

func (a *AuthHandler) Register(c echo.Context) error {
	var user model.User
	response := model.ResponseMessage{}

	user.Name = c.FormValue("name")
	user.FullName = c.FormValue("fullName")

	// Source
	file, _ := c.FormFile("file")
	var fileName string

	if file != nil {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		fileName = fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ReplaceAll(file.Filename, " ", "_"))

		// Destination
		dst, err := os.Create("./public/" + fileName)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		user.Photo = &fileName
	}

	user.Mail = c.FormValue("mail")

	password := encrypt.Bcrypt(c.FormValue("password"))
	user.Password = &password

	user.Phone = c.FormValue("phone")
	specialty := c.FormValue("specialty")
	if specialty == "" {
		user.Specialty = nil
	} else {
		user.Specialty = &specialty
	}
	user.Role = c.FormValue("role")

	id, err := a.db.Register(user)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if id == nil {
		response.Success = false
		response.Message = "Error al crear usuario"
	} else {
		response.Success = true
		response.Message = "Usuario Creado"
	}

	return c.JSON(http.StatusOK, response)
}

func (a *AuthHandler) Login(c echo.Context) error {
	response := model.AuthResponse{}
	message := model.ResponseMessage{}
	var auth model.Auth
	err := c.Bind(&auth)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseError{Message: err.Error()})
	}

	var ok bool
	if ok, err = isRequestValid(&auth); !ok {
		return c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}

	user, err := a.db.GetUserByEmail(auth.Mail)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false,Message: "Experto no encontrado"})
	}

	compare := encrypt.BcryptCheck(*user.Password, auth.Password)

	if compare {
		user.Password = nil

		//token
		// Set custom claims
		claims := &model.JwtCustomClaims{
			ID: user.UserId,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(fmt.Sprintf("%s", viper.Get("VALIDATION_API_SECRET"))))
		if err != nil {
			return err
		}

		response.Success = compare
		response.Token = t
		response.Data = user
		return c.JSON(http.StatusOK, response)
	} else {
		message.Success = compare
		message.Message = "Usuario o Clave Incorrecto"
		return c.JSON(http.StatusNotFound, message)
	}

	return nil
}

func isRequestValid(m interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrNotFound:
		return http.StatusNotFound
	case model.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}