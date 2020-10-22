package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goandreus/validation-go/datalayer/datastore"
	"github.com/goandreus/validation-go/model"
	"github.com/labstack/echo/v4"
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

	g.GET("/network/:userId", handler.FetchAllNetworkByUser)
	g.GET("/expertfind/:param", handler.FetchExpertFindParam)
	g.POST("/networkrequest", handler.CreateNetworkRequest)
	g.GET("/research/:userId", handler.FetchResearchByUser)
	g.GET("/dimension/:researchId", handler.FetchDimensionByResearch)
	g.POST("/dimension", handler.CreateDimension)
	g.POST("/research", handler.CreateResearch)
	g.DELETE("/dimension-delete/:dimensionId", handler.FetchDeleteDimension)
	g.PUT("/research-status/:researchId", handler.FethUpdateResearchStatus)
	g.PUT("/dimension-update/:dimensionId", handler.FethUpdateDimensionStatus)
	g.GET("/research-only/:researchId", handler.FetchResearchById)
	g.GET("/criterio/:speciality/:expertId", handler.FetchCriterioByExpert)
	g.PUT("/research-all/:researchId", handler.UpdateResearch)
	g.POST("/criterio-response", handler.CreateCriterioResponse)
	g.PUT("/criterio-response/:criterioResponseId", handler.UpdateCriterioResponse)
	g.DELETE("/research-delete/:researchId", handler.FetchDeleteResearch)
	g.GET("/research-revision/:userId", handler.FetchResearchByExpert)
	g.GET("/criterio-response-get/:researchId", handler.FetchCriterioResponseByResearchId)

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
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Experto no encontrado"})
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
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
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
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Solicitud no encontrado"})
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
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Solicitud no encontrado"})
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
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Solicitud no encontrado"})
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
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Solicitud no encontrado"})
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
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchAllNetworkByUser(c echo.Context) error {
	userIdStr := c.Param("userId")
	if userIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	userId, _ := strconv.Atoi(userIdStr)

	data, err := m.db.GetNetworkByUser(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchExpertFindParam(c echo.Context) error {
	userIdStr := c.Param("param")

	data, err := m.db.GetExpertFindParam(userIdStr)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) CreateNetworkRequest(c echo.Context) error {
	var networkRequest model.NetworkRequest
	response := model.ResponseMessage{}

	UserBaseIdStr := c.FormValue("userBaseId")
	UserBaseId, _ := strconv.Atoi(UserBaseIdStr)
	networkRequest.UserBaseId = UserBaseId

	UserRelationIdStr := c.FormValue("userRelationId")
	UserRelationId, _ := strconv.Atoi(UserRelationIdStr)
	networkRequest.UserRelationId = UserRelationId
	networkRequest.Status = 1

	id, err := m.db.CreateNetworkRequest(networkRequest)
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

func (m *MobileHandler) FetchResearchByUser(c echo.Context) error {
	userIdStr := c.Param("userId")
	if userIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	userId, _ := strconv.Atoi(userIdStr)

	data, err := m.db.GetResearchByUser(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchDimensionByResearch(c echo.Context) error {
	researchIdStr := c.Param("researchId")
	if researchIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	researchId, _ := strconv.Atoi(researchIdStr)

	data, err := m.db.GetDimensionByResearchId(researchId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) CreateDimension(c echo.Context) error {
	var dimension model.Dimension
	response := model.ResponseMessage{}

	researchIdStr := c.FormValue("researchId")
	researchId, _ := strconv.Atoi(researchIdStr)
	dimension.ResearchId = researchId

	nameStr := c.FormValue("name")
	dimension.Name = nameStr
	variableStr := c.FormValue("variable")
	dimension.Variable = variableStr

	dimension.Status = "P"

	id, err := m.db.CreateDimension(dimension)
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
func (m *MobileHandler) CreateResearch(c echo.Context) error {
	var research model.Research
	response := model.ResponseMessage{}

	researcherIdStr := c.FormValue("researcherId")
	researcherId, _ := strconv.Atoi(researcherIdStr)
	research.ResearcherId = researcherId

	expertIdStr := c.FormValue("expertId")
	expertId, _ := strconv.Atoi(expertIdStr)
	research.ExpertId = expertId

	titleStr := c.FormValue("title")
	research.Title = titleStr

	specialityStr := c.FormValue("speciality")
	research.Speciality = specialityStr

	authorsStr := c.FormValue("authors")
	research.Authors = authorsStr

	observationStr := c.FormValue("observation")
	research.Observation = observationStr

	// Attachment One
	attachmentOne, _ := c.FormFile("attachmentOne")
	var attachmentOneName string

	if attachmentOne != nil {
		src, err := attachmentOne.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		attachmentOneName = fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ReplaceAll(attachmentOne.Filename, " ", "_"))

		// Destination
		dst, err := os.Create("./public/" + attachmentOneName)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		research.AttachmentOne = &attachmentOneName
	}

	// Attachment Two
	attachmentTwo, _ := c.FormFile("attachmentTwo")
	var attachmentTwoName string

	if attachmentTwo != nil {
		src, err := attachmentTwo.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		attachmentTwoName = fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ReplaceAll(attachmentTwo.Filename, " ", "_"))

		// Destination
		dst, err := os.Create("./public/" + attachmentTwoName)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		research.AttachmentTwo = &attachmentTwoName
	}

	id, err := m.db.CreateResearch(research)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if id == nil {
		response.Success = false
		response.Message = "Error al crear respuesta"
	} else {
		response.Success = true
		response.Message = "Respuesta Creada"
		response.Data = id
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchDeleteDimension(c echo.Context) error {
	response := model.ResponseMessage{}
	dimensionIdStr := c.Param("dimensionId")
	fmt.Println("OK")
	fmt.Println(dimensionIdStr)
	fmt.Println("ERROR")
	if dimensionIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	dimensionId, _ := strconv.Atoi(dimensionIdStr)
	resp, err := m.db.DeleteDimension(dimensionId)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if resp == nil {
		response.Success = false
		response.Message = "Error al eliminar dimension"
	} else {
		response.Success = true
		response.Message = "Dimension eliminada"
	}
	return c.JSON(http.StatusOK, response)
}
func (m *MobileHandler) FethUpdateResearchStatus(c echo.Context) error {
	var research model.Research
	response := model.ResponseMessage{}

	researchIdStr := c.Param("researchId")
	researchId, _ := strconv.Atoi(researchIdStr)
	research.ResearchId = researchId

	observationStr := c.FormValue("observation")
	research.Observation = observationStr
	statusStr := c.FormValue("status")
	status, _ := strconv.Atoi(statusStr)
	research.Status = status

	id, err := m.db.UpdateResearchStatus(research)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if id == nil {
		response.Success = false
		response.Message = "Error al crear respuesta"
	} else {
		response.Success = true
		response.Message = "Respuesta Creada"
		response.Data = id
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FethUpdateDimensionStatus(c echo.Context) error {
	var dimension model.Dimension
	response := model.ResponseMessage{}

	researchIdStr := c.Param("researchId")
	researchId, _ := strconv.Atoi(researchIdStr)
	dimension.DimensionId = researchId

	statusStr := c.FormValue("status")
	dimension.Status = statusStr

	id, err := m.db.UpdateDimensionStatus(dimension)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if id == nil {
		response.Success = false
		response.Message = "Error al crear respuesta"
	} else {
		response.Success = true
		response.Message = "Respuesta Creada"
		response.Data = id
	}

	return c.JSON(http.StatusOK, response)
}
func (m *MobileHandler) FetchResearchById(c echo.Context) error {
	researchIdStr := c.Param("researchId")
	if researchIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	researchId, _ := strconv.Atoi(researchIdStr)
	data, err := m.db.GetResearchById(researchId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchCriterioByExpert(c echo.Context) error {
	specialityStr := c.Param("speciality")
	if specialityStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}
	expertIdStr := c.Param("expertId")
	if expertIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	expertId, _ := strconv.Atoi(expertIdStr)

	data, err := m.db.GetCriterioByExpert(specialityStr, expertId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) UpdateResearch(c echo.Context) error {
	var research model.Research
	response := model.ResponseMessage{}
	researchIdStr := c.Param("researchId")
	if researchIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	researchId, _ := strconv.Atoi(researchIdStr)
	research.ResearchId = researchId

	researcherIdStr := c.FormValue("researcherId")
	researcherId, _ := strconv.Atoi(researcherIdStr)
	research.ResearcherId = researcherId

	expertIdStr := c.FormValue("expertId")
	expertId, _ := strconv.Atoi(expertIdStr)
	research.ExpertId = expertId

	titleStr := c.FormValue("title")
	research.Title = titleStr

	specialityStr := c.FormValue("speciality")
	research.Speciality = specialityStr

	authorsStr := c.FormValue("authors")
	research.Authors = authorsStr

	observationStr := c.FormValue("observation")
	research.Observation = observationStr

	// Attachment One
	attachmentOne, _ := c.FormFile("attachmentOneFile")
	var attachmentOneName string

	if attachmentOne != nil {
		src, err := attachmentOne.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		attachmentOneName = fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ReplaceAll(attachmentOne.Filename, " ", "_"))

		// Destination
		dst, err := os.Create("./public/" + attachmentOneName)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		research.AttachmentOne = &attachmentOneName
	} else {
		attachmentOneAux := c.FormValue("attachmentOne")
		research.AttachmentOne = &attachmentOneAux
	}

	// Attachment Two
	attachmentTwo, _ := c.FormFile("attachmentTwoFile")
	var attachmentTwoName string

	if attachmentTwo != nil {
		src, err := attachmentTwo.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		attachmentTwoName = fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ReplaceAll(attachmentTwo.Filename, " ", "_"))

		// Destination
		dst, err := os.Create("./public/" + attachmentTwoName)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		research.AttachmentTwo = &attachmentTwoName
	} else {
		attachmentTwoAux := c.FormValue("attachmentTwo")
		research.AttachmentTwo = &attachmentTwoAux
	}

	statusStr := c.FormValue("status")
	status, _ := strconv.Atoi(statusStr)
	research.Status = status
	id, err := m.db.UpdateResearch(research)
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
func (m *MobileHandler) CreateCriterioResponse(c echo.Context) error {
	var criterioResponse model.CriterioResponse
	response := model.ResponseMessage{}

	CriterioIdStr := c.FormValue("criterioId")
	criterioId, _ := strconv.Atoi(CriterioIdStr)
	criterioResponse.CriterioId = criterioId

	researchIdStr := c.FormValue("researchId")
	researchId, _ := strconv.Atoi(researchIdStr)
	criterioResponse.ResearchId = researchId

	dimensionIdStr := c.FormValue("dimensionId")
	dimensionId, _ := strconv.Atoi(dimensionIdStr)
	criterioResponse.DimensionId = dimensionId

	statusStr := c.FormValue("status")
	criterioResponse.Status = statusStr

	id, err := m.db.CreateCriterioResponse(criterioResponse)
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
func (m *MobileHandler) UpdateCriterioResponse(c echo.Context) error {
	var criterioResponse model.CriterioResponse
	response := model.ResponseMessage{}
	criterioResponseIdStr := c.Param("criterioResponseId")
	if criterioResponseIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}
	criterioResponseId, _ := strconv.Atoi(criterioResponseIdStr)
	criterioResponse.CriterioResponseId = criterioResponseId

	CriterioIdStr := c.FormValue("criterioId")
	criterioId, _ := strconv.Atoi(CriterioIdStr)
	criterioResponse.CriterioId = criterioId

	researchIdStr := c.FormValue("researchId")
	researchId, _ := strconv.Atoi(researchIdStr)
	criterioResponse.ResearchId = researchId

	dimensionIdStr := c.FormValue("dimensionId")
	dimensionId, _ := strconv.Atoi(dimensionIdStr)
	criterioResponse.DimensionId = dimensionId

	statusStr := c.FormValue("status")
	criterioResponse.Status = statusStr
	id, err := m.db.UpdateCriterioResponse(criterioResponse)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if id == nil {
		response.Success = false
		response.Message = "Error al crear respuesta"
	} else {
		response.Success = true
		response.Message = "Respuesta Creada"
		response.Data = id
	}

	return c.JSON(http.StatusOK, response)
}
func (m *MobileHandler) FetchDeleteResearch(c echo.Context) error {
	response := model.ResponseMessage{}
	researchIdStr := c.Param("researchId")
	if researchIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	researchId, _ := strconv.Atoi(researchIdStr)
	resp, err := m.db.DeleteResearch(researchId)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if resp == nil {
		response.Success = false
		response.Message = "Error al eliminar dimension"
	} else {
		response.Success = true
		response.Message = "Dimension eliminada"
	}
	return c.JSON(http.StatusOK, response)
}
func (m *MobileHandler) FetchResearchByExpert(c echo.Context) error {
	userIdStr := c.Param("userId")
	if userIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}
	userId, _ := strconv.Atoi(userIdStr)

	data, err := m.db.GetResearchByExpert(userId, 2)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}
func (m *MobileHandler) FetchCriterioResponseByResearchId(c echo.Context) error {
	researchIdStr := c.Param("researchId")
	if researchIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	researchId, _ := strconv.Atoi(researchIdStr)

	data, err := m.db.GetCriterioResponseByResearchId(researchId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}
