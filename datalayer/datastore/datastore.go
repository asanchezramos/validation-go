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
}