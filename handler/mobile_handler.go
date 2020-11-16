package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goandreus/validation-go/datalayer/datastore"
	"github.com/goandreus/validation-go/model"
	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
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
	g.PUT("/dimension-update-only/:dimensionId", handler.UpdateDimension)
	g.GET("/certificate/:researchId", handler.FetchCertificateByResearchId)
	g.GET("/network-request/:userId", handler.FetchNetworkRequestByUserId)
	g.GET("/network-request-expert/:userId", handler.FetchNetworkRequestByExpertId)
	g.PUT("/network-request-response/:userBaseId/:userRelationId", handler.NetworkRequestResponse)
	g.DELETE("/resource-user/:resourceUserId", handler.FetchDeleteResourceUser)
	g.POST("/resource-user-post", handler.FetchCreateResourceUser)
	g.GET("/resource-user-post/:userId", handler.FetchResourcesUserBy)

	g.POST("/signing-create", handler.CreateSigning)
	g.PUT("/signing-update/:signingId", handler.UpdateSigning)
	g.GET("/signing-get/:expertId", handler.GetSigning)

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
func (m *MobileHandler) UpdateDimension(c echo.Context) error {
	var dimension model.Dimension
	response := model.ResponseMessage{}

	dimensionIdStr := c.Param("dimensionId")
	dimensionId, _ := strconv.Atoi(dimensionIdStr)
	dimension.DimensionId = dimensionId
	researchIdStr := c.FormValue("researchId")
	researchId, _ := strconv.Atoi(researchIdStr)
	dimension.ResearchId = researchId

	nameStr := c.FormValue("name")
	dimension.Name = nameStr
	variableStr := c.FormValue("variable")
	dimension.Variable = variableStr

	statusStr := c.FormValue("status")
	dimension.Status = statusStr

	id, err := m.db.UpdateDimension(dimension)
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
func (m *MobileHandler) FetchCertificateByResearchId(c echo.Context) error {
	researchIdStr := c.Param("researchId")
	if researchIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}
	researchId, _ := strconv.Atoi(researchIdStr)
	data, errC := m.db.GetResearchById(researchId)
	if errC != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	dataExpert, errB := m.db.GetExpertById(data.ExpertId)
	if errB != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Experto no encontrado"})
	}

	pdf := gofpdf.New("L", "mm", "A4", "")

	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetTitle("Certificado", true)
	pdf.SetTopMargin(10)
	pdf.AddPage()
	pdf.SetFont("Times", "B", 28)
	pdf.CellFormat(0, 10, "Certificado de validez", "", 0, "C", false, 0, "")
	pdf.Ln(12)
	pdf.SetFont("Times", "", 20)
	pdf.CellFormat(0, 10, tr("POR MEDIO DE LA PRESENTE SE CERTIFICA LA VALIDEZ DE LA INVESTIGACIÓN"), "", 0, "C", false, 0, "")
	pdf.Ln(20)
	pdf.SetFont("Times", "", 35)
	pdf.CellFormat(0, 20, tr(string(data.Title)), "", 2, "C", false, 0, "")

	pdf.Ln(12)
	pdf.SetFont("Times", "", 20)
	pdf.CellFormat(0, 10, "DE LOS AUTORES", "", 1, "C", false, 0, "")
	pdf.Ln(20)
	pdf.SetFont("Times", "", 35)
	pdf.CellFormat(0, 10, tr(data.Authors), "", 0, "C", false, 0, "")
	pdf.Ln(12)
	pdf.SetFont("Times", "", 20)
	pdf.CellFormat(0, 10, "APROBADO POR: "+dataExpert.Name+dataExpert.FullName, "", 0, "C", false, 0, "")

	pdf.Ln(20)
	pdf.SetFont("Times", "", 20)
	pdf.CellFormat(0, 10, data.UpdatedAt, "", 0, "C", false, 0, "")
	dataSigning, errA := m.db.FetchGetSigning(data.ExpertId)
	if errA != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "No hay error"})
	}
	url := "http://192.168.31.210:8080/public/" + dataSigning[0].Link
	httpimg.Register(pdf, url, "")
	pdf.Image(url, 15, 15, 267, 0, false, "", 0, "")
	//pdf table header
	//pdf = header(pdf, []string{"1st column", "2nd", "3rd", "4th", "5th", "6th"})

	//pdf table content
	//pdf = table(pdf, data)

	if pdf.Err() {
		log.Fatalf("failed ! %s", pdf.Error())
	}
	attachmentName := fmt.Sprintf("%d-%s", time.Now().Unix(), "certificado.pdf")
	//attachmentName := "attachment.pdf"
	//attachmentOneName = fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ReplaceAll(attachmentOne.Filename, " ", "_"))

	errx := pdf.OutputFileAndClose("./public/" + attachmentName)
	if errx != nil {
		log.Fatalf("error saving pdf file: %s", errx)
	}
	response := model.Response{
		Success: true,
		Data:    attachmentName,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchNetworkRequestByUserId(c echo.Context) error {
	userIdStr := c.Param("userId")
	if userIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	userId, _ := strconv.Atoi(userIdStr)

	data, err := m.db.FetchNetworkRequestByUserId(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}
func (m *MobileHandler) FetchNetworkRequestByExpertId(c echo.Context) error {

	userIdStr := c.Param("userId")
	if userIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	userId, _ := strconv.Atoi(userIdStr)

	data, err := m.db.FetchNetworkRequestByExpertId(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) NetworkRequestResponse(c echo.Context) error {

	userBaseIdStr := c.Param("userBaseId")
	if userBaseIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	userBaseId, _ := strconv.Atoi(userBaseIdStr)

	userRelationIdStr := c.Param("userRelationId")
	if userRelationIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	userRelationId, _ := strconv.Atoi(userRelationIdStr)
	var network model.Network
	network.UserBaseId = userBaseId
	network.UserRelationId = userRelationId
	id, err := m.db.CreateNetwork(network)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	var networkRequest model.NetworkRequest
	networkRequest.Status = 2
	networkRequest.UserBaseId = userBaseId
	networkRequest.UserRelationId = userRelationId
	id, errs := m.db.UpdateNetworkRequest(networkRequest)
	if errs != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    id,
	}

	return c.JSON(http.StatusOK, response)
}
func (m *MobileHandler) FetchResourcesUserBy(c echo.Context) error {

	userIdStr := c.Param("userId")
	if userIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	userId, _ := strconv.Atoi(userIdStr)

	data, err := m.db.FetchAllResourceUserById(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "Usuario no encontrado"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) FetchCreateResourceUser(c echo.Context) error {
	var resourceUser model.ResourceUser
	response := model.ResponseMessage{}
	titleStr := c.FormValue("title")
	resourceUser.Title = titleStr

	subtitleStr := c.FormValue("subtitle")
	resourceUser.Subtitle = subtitleStr

	linkStr := c.FormValue("link")
	resourceUser.Link = linkStr

	expertIdStr := c.FormValue("expertId")
	expertId, _ := strconv.Atoi(expertIdStr)
	resourceUser.ExpertId = expertId
	id, err := m.db.CreateResourceUser(resourceUser)
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
func (m *MobileHandler) FetchDeleteResourceUser(c echo.Context) error {

	response := model.ResponseMessage{}
	resourceUserIdStr := c.Param("resourceUserId")
	if resourceUserIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}
	resourceUserId, _ := strconv.Atoi(resourceUserIdStr)

	id, err := m.db.DeleteResourceUser(resourceUserId)
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

func (m *MobileHandler) CreateSigning(c echo.Context) error {
	var signing model.Signing
	response := model.ResponseMessage{}

	// Source
	file, _ := c.FormFile("fileSigning")
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

		signing.Link = fileName
	}
	expertIdStr := c.FormValue("expertId")
	expertId, _ := strconv.Atoi(expertIdStr)
	signing.ExpertId = expertId
	id, err := m.db.FetchCreateSigning(signing)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if id == nil {
		response.Success = false
		response.Message = "Error al crear firma"
	} else {
		response.Success = true
		response.Message = "Respuesta Creada"
		response.Data = id
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) UpdateSigning(c echo.Context) error {
	var signing model.Signing
	response := model.ResponseMessage{}
	signingIdStr := c.Param("signingId")
	if signingIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}
	signingId, _ := strconv.Atoi(signingIdStr)
	signing.SigningId = signingId
	// Source
	file, _ := c.FormFile("fileSigning")
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

		signing.Link = fileName
	} else {
		linkStr := c.FormValue("link")
		signing.Link = linkStr
	}
	expertIdStr := c.FormValue("expertId")
	expertId, _ := strconv.Atoi(expertIdStr)
	signing.ExpertId = expertId
	id, err := m.db.FetchUpdateSigning(signing)
	if err != nil {
		return c.JSON(getStatusCode(err), model.ResponseError{Message: err.Error()})
	}

	if id == nil {
		response.Success = false
		response.Message = "Error al crear firma"
	} else {
		response.Success = true
		response.Message = "Respuesta Creada"
		response.Data = id
	}

	return c.JSON(http.StatusOK, response)
}

func (m *MobileHandler) GetSigning(c echo.Context) error {
	expertIdStr := c.Param("expertId")
	if expertIdStr == "" {
		return c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}
	expertId, _ := strconv.Atoi(expertIdStr)

	data, err := m.db.FetchGetSigning(expertId)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseMessage{Success: false, Message: "No hay error"})
	}

	response := model.Response{
		Success: true,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}
