package handler

import (
	"fmt"
	"github.com/goandreus/validation-go/datalayer/datastore"
	"github.com/goandreus/validation-go/model"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// MobileHandler  represent the httphandler for mobile endpoints
type MobileHandler struct {
	db datastore.Database
}

// NewMobileHandler will initialize the mobile/ resources endpoint
func NewMobileHandler(e *echo.Group, db datastore.Database) {
	handler := &MobileHandler{
		db,
	}
	g := e.Group("/mobile")
	g.GET("/expert", handler.FetchAllExpert)
	g.GET("/expert/:expertId", handler.FetchExpert)
	g.POST("/solicitude", handler.CreateSolicitude)
	g.GET("/solicitude/user/:userId", handler.FetchAllSolicitudeByUser)
	g.GET("/solicitude-answer/:solicitudeId", handler.FetchAllSolicitudeAnswerByUser)
	g.GET("/solicitude-user/:expertId", handler.FetchAllUserSolicitudeByExpert)
	g.GET("/solicitude-user-expert/:solicitudeId", handler.FetchAllUserSolicitudeByExpertDetail)
	g.PUT("/solicitude/:solicitudeId/:status", handler.UpdateSolicitudeStatus)
	g.POST("/answer", handler.CreateAnswer)
	g.GET("/solicitude/expert/:expertId", handler.FetchAllSolicitudeByExpert)
}

func (m *MobileHandler) FetchAllExpert(c echo.Context) error {
	data, err := m.db.GetAllExpert()
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchExpert(c echo.Context) error {
	expertIdStr := c.Param("expertId")
	if expertIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	expertId, _ := strconv.Atoi(expertIdStr)

	data, err := m.db.GetExpertById(expertId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false,Message: "Experto no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) CreateSolicitude(c echo.Context) error {
	var solicitude model.Solicitude
	response := model.ResponseMessage{}

	// Source
	repository, _ := c.FormFile("repository")
	var repositoryName string

	if repository != nil {
		src, err := repository.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		repositoryName = fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ReplaceAll(repository.Filename, " ", "_"))

		// Destination
		dst, err := os.Create("./public/" + repositoryName)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		solicitude.Repository = repositoryName
	}

	// Source
	investigation, _ := c.FormFile("investigation")
	var investigationName string

	if investigation != nil {
		src, err := investigation.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		investigationName = fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ReplaceAll(investigation.Filename, " ", "_"))

		// Destination
		dst, err := os.Create("./public/" + investigationName)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		solicitude.Investigation = &investigationName
	}

	userIdStr := c.FormValue("userId")
	expertIdStr := c.FormValue("expertId")
	userId, _ := strconv.Atoi(userIdStr)
	expertId, _ := strconv.Atoi(expertIdStr)
	solicitude.UserId = userId
	solicitude.ExpertId = expertId
	solicitude.Status = c.FormValue("status")

	id, err := m.db.CreateSolicitude(solicitude)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if id == nil {
		response.Success = false
		response.Message = "Error al crear solicitud"
	} else {
		response.Success = true
		response.Message = "Solicitud Creado"
	}


	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchAllSolicitudeByUser(c echo.Context) error {
	userIdStr := c.Param("userId")
	if userIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	userId, _ := strconv.Atoi(userIdStr)

	data, err := m.db.GetAllSolicitudeByUser(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false,Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchAllSolicitudeAnswerByUser(c echo.Context) error {
	solicitudeIdStr := c.Param("solicitudeId")
	if solicitudeIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	solicitudeId, _ := strconv.Atoi(solicitudeIdStr)

	data, err := m.db.GetAnswerBySolicitude(solicitudeId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false,Message: "Solicitud no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}


func (m *MobileHandler) FetchAllUserSolicitudeByExpert(c echo.Context) error {
	expertIdStr := c.Param("expertId")
	if expertIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	expertId, _ := strconv.Atoi(expertIdStr)

	data, err := m.db.GetAllUserByExpert(expertId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false,Message: "Solicitud no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) UpdateSolicitudeStatus(c echo.Context) error {
	response := model.ResponseMessage{}
	solicitudeIdStr := c.Param("solicitudeId")
	status := c.Param("status")
	if solicitudeIdStr == "" && status == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	solicitudeId, _ := strconv.Atoi(solicitudeIdStr)

	row, err := m.db.UpdateStatusSolicitude(solicitudeId, status)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false,Message: "Solicitud no encontrado"})
	}

	if *row == 0 {
		response.Success = false
		response.Message = "Error al actualizar solicitud"
	} else {
		response.Success = true
		response.Message = "Solicitud actualizada"
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchAllUserSolicitudeByExpertDetail(c echo.Context) error {
	solicitudeIdStr := c.Param("solicitudeId")
	if solicitudeIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	solicitudeId, _ := strconv.Atoi(solicitudeIdStr)

	data, err := m.db.GetSolicitudeStudentBySolicitudeId(solicitudeId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false,Message: "Solicitud no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) CreateAnswer(c echo.Context) error {
	var answer model.Answer
	response := model.ResponseMessage{}


	answer.Comments = c.FormValue("comments")

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

		answer.File = &fileName
	}

	solicitudeIdStr := c.FormValue("solicitudeId")
	solicitudeId, _ := strconv.Atoi(solicitudeIdStr)
	answer.SolicitudeId = solicitudeId

	id, err := m.db.CreateAnswer(answer)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if id == nil {
		response.Success = false
		response.Message = "Error al crear respuesta"
	} else {
		response.Success = true
		response.Message = "Respuesta Creada"
	}


	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchAllSolicitudeByExpert(c echo.Context) error {
	expertIdStr := c.Param("expertId")
	if expertIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	expertId, _ := strconv.Atoi(expertIdStr)

	data, err := m.db.GetAllSolicitudeByExpert(expertId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false,Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}