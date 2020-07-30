package datalayer

import (
	"database/sql"
	"fmt"
	"github.com/goandreus/validation-go/datalayer/datastore"
	"github.com/goandreus/validation-go/model"
	"github.com/spf13/viper"
	"reflect"
)

type MysqlDatastore struct {
	databases map[string]*MysqlDatabase
}

// ArangoDatabase represents an ArangoDB database
type MysqlDatabase struct {
	db          *sql.DB
}

// NewDatastore creates a new instance of the MySQL data provider
func NewDatastore() datastore.Datastore {
	return &MysqlDatastore{
	}
}

//Open
func (p *MysqlDatastore) Open() (datastore.Database, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", viper.Get("VALIDATION_DB_USER"),
								viper.Get("VALIDATION_DB_PASSWORD"), viper.Get("VALIDATION_DB_HOST"),
								viper.Get("VALIDATION_DB_PORT"), viper.Get("VALIDATION_DB_NAME"))
	connString := uri
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	if p.databases == nil {
		p.databases = make(map[string]*MysqlDatabase)
	}

	result := &MysqlDatabase{
		db:          db,
	}
	return result, nil
}

func (m *MysqlDatabase) Register(user model.User) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO `user` (`name`, `full_name`, `photo`, `mail`, `password`, `phone`, `specialty`, `role`) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.FullName, user.Photo, user.Mail, user.Password, user.Phone, user.Specialty, user.Role)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (m *MysqlDatabase) GetUserByEmail(email string) (*model.User, error) {
	status := 1
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
		tx.Rollback()
	}

	item := &model.User{}

	row := tx.QueryRow("SELECT * FROM `user` WHERE mail = ? AND `status` = ?", email, status)

	if row != nil {
		err := row.Scan(ScanRow(item)...)
		if err != nil {
			return nil, err
			tx.Rollback()
		}

		if item.Photo != nil {
			*item.Photo = fmt.Sprintf("%s%s/%s", viper.Get("VALIDATION_BASE_URL"), viper.Get("VALIDATION_PUBLIC"), *item.Photo)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *MysqlDatabase) GetAllExpert() ([]*model.User, error)  {
	status := 1
	expert := "E"
	result := make([]*model.User, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT * FROM `user` WHERE `status` = ? AND role = ?", status, expert)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.User{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		items.Password = nil

		if items.Photo != nil {
			*items.Photo = fmt.Sprintf("%s%s/%s", viper.Get("VALIDATION_BASE_URL"), viper.Get("VALIDATION_PUBLIC"), *items.Photo)
		}

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MysqlDatabase) GetExpertById(expertId int) (*model.User, error) {
	status := 1
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
		tx.Rollback()
	}

	item := &model.User{}

	row := tx.QueryRow("SELECT * FROM `user` WHERE user_id = ? AND role = 'E' AND `status` = ?", expertId, status)

	if row != nil {
		err := row.Scan(ScanRow(item)...)
		if err != nil {
			return nil, err
			tx.Rollback()
		}

		if item.Photo != nil {
			*item.Photo = fmt.Sprintf("%s%s/%s", viper.Get("VALIDATION_BASE_URL"), viper.Get("VALIDATION_PUBLIC"), *item.Photo)
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (m *MysqlDatabase) CreateSolicitude(solicitude model.Solicitude) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO solicitude (`repository`, investigation, user_id, expert_id, status) VALUES (?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(solicitude.Repository, solicitude.Investigation, solicitude.UserId, solicitude.ExpertId, solicitude.Status)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (m *MysqlDatabase) GetAllSolicitudeByUser(userId int) ([]*model.SolicitudeUser, error) {
	result := make([]*model.SolicitudeUser, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT s.solicitude_id, CONCAT(u.`name`, ' ', u.full_name) as full_name, u.specialty, s.status " +
								" FROM solicitude s " +
								" INNER JOIN `user` u ON s.expert_id = u.user_id AND u.`role` = 'E' " +
								" WHERE s.user_id = ?", userId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.SolicitudeUser{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		switch items.Status {
		case "P":
			items.Status = "Pendiente"
			break
		case "C":
			items.Status = "Completado"
			break
		case "D":
			items.Status = "Denegado"
			break
		default:
			items.Status = "Not found"
		}

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MysqlDatabase) GetAnswerBySolicitude(solicitudeId int) ([]*model.ExpertAnswer, error)  {
	result := make([]*model.ExpertAnswer, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT a.answer_id , CONCAT(u.`name`, ' ', u.full_name) as full_name, u.photo, u.specialty, a.`file`, a.comments " +
		"FROM answer a " +
		"INNER JOIN solicitude s ON a.solicitude_id = s.solicitude_id " +
		"INNER JOIN `user` u ON s.expert_id = u.user_id AND u.`role` = 'E'" +
		"WHERE a.solicitude_id = ?", solicitudeId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.ExpertAnswer{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if items.Photo != nil {
			*items.Photo = fmt.Sprintf("%s%s/%s", viper.Get("VALIDATION_BASE_URL"), viper.Get("VALIDATION_PUBLIC"), *items.Photo)
		}

		if items.File != nil {
			*items.File = fmt.Sprintf("%s%s/%s", viper.Get("VALIDATION_BASE_URL"), viper.Get("VALIDATION_PUBLIC"), *items.File)
		}

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MysqlDatabase) GetAllUserByExpert(expertId int) ([]*model.UserSolicitude, error) {
	status := "P"
	result := make([]*model.UserSolicitude, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT s.solicitude_id , u.* FROM solicitude s " +
							"INNER JOIN `user` u ON s.user_id = u.user_id AND u.`role` = 'U' " +
							"WHERE s.expert_id = ? AND s.`status` = ?", expertId, status)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.UserSolicitude{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if items.Photo != nil {
			*items.Photo = fmt.Sprintf("%s%s/%s", viper.Get("VALIDATION_BASE_URL"), viper.Get("VALIDATION_PUBLIC"), *items.Photo)
		}
		items.Password = nil

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MysqlDatabase) GetSolicitudeStudentBySolicitudeId(solicitudeId int) ([]*model.StudentSolicitude, error) {
	result := make([]*model.StudentSolicitude, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT s.solicitude_id, s.repository, s.investigation, CONCAT(u.`name`, ' ', u.full_name) AS full_name, u.photo FROM solicitude s " +
								"INNER JOIN `user` u ON s.user_id = u.user_id AND u.`role` = 'U' " +
								"WHERE s.solicitude_id = ?", solicitudeId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.StudentSolicitude{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if items.Investigation != nil {
			*items.Investigation = fmt.Sprintf("%s%s/%s", viper.Get("VALIDATION_BASE_URL"), viper.Get("VALIDATION_PUBLIC"), *items.Investigation)
		}

		if items.Repository != nil {
			*items.Repository = fmt.Sprintf("%s%s/%s", viper.Get("VALIDATION_BASE_URL"), viper.Get("VALIDATION_PUBLIC"), *items.Repository)
		}

		if items.Photo != nil {
			*items.Photo = fmt.Sprintf("%s%s/%s", viper.Get("VALIDATION_BASE_URL"), viper.Get("VALIDATION_PUBLIC"), *items.Photo)
		}

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MysqlDatabase) CreateAnswer(answer model.Answer) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO answer (`comments`, `file`, solicitude_id) VALUES (?,?,?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(answer.Comments, answer.File, answer.SolicitudeId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (m *MysqlDatabase) UpdateStatusSolicitude(solicitudeId int, status string) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE solicitude SET `status` = ? WHERE solicitude_id = ?;")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(status, solicitudeId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &rows, nil
}

func (m *MysqlDatabase) GetAllSolicitudeByExpert(expertId int) ([]*model.SolicitudeUser, error) {
	result := make([]*model.SolicitudeUser, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT s.solicitude_id, CONCAT(u.`name`, ' ', u.full_name) as full_name, u.specialty, s.status " +
		" FROM solicitude s " +
		" INNER JOIN `user` u ON s.user_id = u.user_id AND u.`role` = 'U' " +
		" WHERE s.expert_id = ?", expertId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.SolicitudeUser{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		switch items.Status {
		case "P":
			items.Status = "Pendiente"
			break
		case "C":
			items.Status = "Completado"
			break
		case "D":
			items.Status = "Denegado"
			break
		default:
			items.Status = "Not found"
		}

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}

//ScanRow attributes from sql resulset
func ScanRow(row interface{}) []interface{} {

	rt := reflect.TypeOf(row).Elem()
	rv := reflect.ValueOf(row).Elem()

	numCols := rv.NumField()

	var count int = 0

	for i := 0; i < numCols; i++ {
		f := rt.Field(i)
		// fmt.Println(f.Tag)
		// fmt.Println(f.Tag.Get("db"))
		if f.Tag.Get("db") != "-" {
			count++
		}
	}

	columns := make([]interface{}, count)
	for i := 0; i < count; i++ {
		field := rv.Field(i)
		f := rt.Field(i)
		if f.Tag.Get("db") != "-" {
			columns[i] = field.Addr().Interface()
		}
	}

	return columns
}