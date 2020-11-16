package datastore

import "github.com/goandreus/validation-go/model"

// Datastore is the main datastore implemented by datastore providers
type Datastore interface {
	Open() (Database, error)
}

// Database is the main interface implemented by database providers
type Database interface {
	Register(user model.User) (*int64, error)
	GetUserByEmail(email string) (*model.User, error)
	GetAllExpert() ([]*model.User, error)
	GetExpertById(expertId int) (*model.User, error)
	CreateSolicitude(solicitude model.Solicitude) (*int64, error)
	GetAllSolicitudeByUser(userId int) ([]*model.SolicitudeUser, error)
	GetAnswerBySolicitude(solicitudeId int) ([]*model.ExpertAnswer, error)
	GetAllUserByExpert(expertId int) ([]*model.UserSolicitude, error)
	GetSolicitudeStudentBySolicitudeId(solicitudeId int) ([]*model.StudentSolicitude, error)
	CreateAnswer(answer model.Answer) (*int64, error)
	UpdateStatusSolicitude(solicitudeId int, status string) (*int64, error)
	GetAllSolicitudeByExpert(expertId int) ([]*model.SolicitudeUser, error)

	GetNetworkByUser(userId int) ([]*model.User, error)
	GetExpertFindParam(param string) ([]*model.User, error)
	CreateNetworkRequest(networkRequest model.NetworkRequest) (*int64, error)
	GetResearchByUser(userId int) ([]*model.Research, error)
	GetDimensionByResearchId(researchId int) ([]*model.Dimension, error)
	CreateDimension(dimension model.Dimension) (*int64, error)
	CreateResearch(research model.Research) (*int64, error)
	DeleteDimension(dimensionId int) (*int64, error)
	UpdateResearchStatus(research model.Research) (*int64, error)
	UpdateDimensionStatus(dimension model.Dimension) (*int64, error)
	GetResearchById(researchId int) (*model.Research, error)
	GetCriterioByExpert(speciality string, expertId int) ([]*model.Criterio, error)
	UpdateResearch(research model.Research) (*int64, error)
	DeleteResearch(researchId int) (*int64, error)
	GetResearchByExpert(userId int, status int) ([]*model.Research, error)
	GetCriterioResponseByResearchId(researchId int) ([]*model.CriterioResponse, error)
	CreateCriterioResponse(criterioResponse model.CriterioResponse) (*int64, error)
	UpdateCriterioResponse(criterioResponse model.CriterioResponse) (*int64, error)
	UpdateDimension(dimension model.Dimension) (*int64, error)
	FetchNetworkRequestByUserId(userId int) ([]*model.User, error)
	FetchNetworkRequestByExpertId(userId int) ([]*model.User, error)
	CreateNetwork(network model.Network) (*int64, error)
	UpdateNetworkRequest(networkReuest model.NetworkRequest) (*int64, error)
	CreateResourceUser(resourceUser model.ResourceUser) (*int64, error)
	DeleteResourceUser(resourceUserId int) (*int64, error)
	FetchAllResourceUserById(userId int) ([]*model.ResourceUser, error)
	FetchCreateSigning(signing model.Signing) (*int64, error)
	FetchUpdateSigning(signing model.Signing) (*int64, error)
	FetchGetSigning(expertId int) ([]*model.Signing, error)
}
