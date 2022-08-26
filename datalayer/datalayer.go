package datalayer

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/goandreus/validation-go/datalayer/datastore"
	"github.com/goandreus/validation-go/model"
)

type MysqlDatastore struct {
	databases map[string]*MysqlDatabase
}

// ArangoDatabase represents an ArangoDB database
type MysqlDatabase struct {
	db *sql.DB
}

// NewDatastore creates a new instance of the MySQL data provider
func NewDatastore() datastore.Datastore {
	return &MysqlDatastore{}
}
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

//Open
func (p *MysqlDatastore) Open() (datastore.Database, error) {
	/*
			uri := fmt.Sprintf(viper.Get("VALIDATION_DB_USER"), ":", viper.Get("VALIDATION_DB_PASSWORD"), "@tcp(localhost:3306)/%s?charset=utf8", viper.Get("VALIDATION_DB_USER"),
				viper.Get("VALIDATION_DB_PASSWORD"), viper.Get("VALIDATION_DB_HOST"),
				viper.Get("VALIDATION_DB_PORT"), viper.Get("VALIDATION_DB_NAME"))

		var (
			dbUser                 = mustGetenv("CLOUDSQL_USER")            // e.g. 'my-db-user'
			dbPwd                  = mustGetenv("CLOUDSQL_PASSWORD")        // e.g. 'my-db-password'
			instanceConnectionName = mustGetenv("CLOUDSQL_CONNECTION_NAME") // e.g. 'project:region:instance'
			dbName                 = mustGetenv("CLOUDSQL_DATABASE")        // e.g. 'my-database'
		)

		socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
		if !isSet {
			socketDir = "/cloudsql"
		}
	*/
	var connString string
	//connString = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	connString = "admin-mysql:admin2021!@tcp(umadev.pe:3306)/bd_juicio_experto"
	//connString = "root:@tcp(localhost:3306)/bd_juicio_experto"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	if p.databases == nil {
		p.databases = make(map[string]*MysqlDatabase)
	}

	result := &MysqlDatabase{
		db: db,
	}
	return result, nil
}

func (m *MysqlDatabase) Register(user model.User) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO `user` (`name`, `full_name`, `photo`, `mail`, `password`, `phone`, `specialty`, `role`, `orcid`) VALUES (?,?,?,?,?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.FullName, user.Photo, user.Mail, user.Password, user.Phone, user.Specialty, user.Role,user.Orcid)
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

	fmt.Println("0correo y estatuis: " +item.Name + "correo y estatuis: "    )
	if row != nil {
		err := row.Scan(ScanRow(item)...)
		if err != nil {
			fmt.Println(err)
			return nil, err
			tx.Rollback()
		}

	}
	fmt.Println("1correo y estatuis: " +item.Name + "correo y estatuis: "    )

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	fmt.Println("2correo y estatuis: " +item.Name + "correo y estatuis: "    )
	return item, nil
}

func (m *MysqlDatabase) GetAllExpert() ([]*model.User, error) {
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

	row := tx.QueryRow("SELECT * FROM `user` WHERE user_id = ? AND `status` = ?", expertId, status)

	if row != nil {
		err := row.Scan(ScanRow(item)...)
		if err != nil {
			return nil, err
			tx.Rollback()
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

	rows, err := tx.Query("SELECT s.solicitude_id, CONCAT(u.`name`, ' ', u.full_name) as full_name, u.specialty, s.status "+
		" FROM solicitude s "+
		" INNER JOIN `user` u ON s.expert_id = u.user_id AND u.`role` = 'E' "+
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

func (m *MysqlDatabase) GetAnswerBySolicitude(solicitudeId int) ([]*model.ExpertAnswer, error) {
	result := make([]*model.ExpertAnswer, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT a.answer_id , CONCAT(u.`name`, ' ', u.full_name) as full_name, u.photo, u.specialty, a.`file`, a.comments "+
		"FROM answer a "+
		"INNER JOIN solicitude s ON a.solicitude_id = s.solicitude_id "+
		"INNER JOIN `user` u ON s.expert_id = u.user_id AND u.`role` = 'E'"+
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

	rows, err := tx.Query("SELECT s.solicitude_id , u.* FROM solicitude s "+
		"INNER JOIN `user` u ON s.user_id = u.user_id AND u.`role` = 'U' "+
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

	rows, err := tx.Query("SELECT s.solicitude_id, s.repository, s.investigation, CONCAT(u.`name`, ' ', u.full_name) AS full_name, u.photo FROM solicitude s "+
		"INNER JOIN `user` u ON s.user_id = u.user_id AND u.`role` = 'U' "+
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

	rows, err := tx.Query("SELECT s.solicitude_id, CONCAT(u.`name`, ' ', u.full_name) as full_name, u.specialty, s.status "+
		" FROM solicitude s "+
		" INNER JOIN `user` u ON s.user_id = u.user_id AND u.`role` = 'U' "+
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
func (m *MysqlDatabase) GetNetworkByUser(userId int) ([]*model.User, error) {
	status := 1
	expert := "E"
	result := make([]*model.User, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT `user`.* FROM `user` inner join `network` on `user`.`user_id` = `network`.`user_relation_id` where `network`.`user_base_id` = ? and `user`.`status` = ? and `user`.`role`= ?", userId, status, expert)
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

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *MysqlDatabase) GetExpertFindParam(param string) ([]*model.User, error) {
	status := 1
	expert := "E"
	result := make([]*model.User, 0)
	paramAux := "%" + param + "%"

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT `user`.* FROM `user` where CONCAT(`user`.`name`,' ',`user`.`full_name`, ' ' ,  COALESCE(`user`.`specialty`,'') ) like ? and `user`.`status` = ? and `user`.`role`= ?", paramAux, status, expert)
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

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *MysqlDatabase) CreateNetworkRequest(networkRequest model.NetworkRequest) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO `network_request`(  `user_base_id`, `user_relation_id`, `status`) VALUES (?,?,?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(networkRequest.UserBaseId, networkRequest.UserRelationId, networkRequest.Status)
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
func (m *MysqlDatabase) GetResearchByUser(userId int) ([]*model.Research, error) {
	result := make([]*model.Research, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT `research_id`, `researcher_id`, `expert_id`, `title`, `speciality`, `authors`, `observation`, `attachment_one`, `attachment_two`, `status`, `created_at`, `updated_at` FROM `research` WHERE  `researcher_id` = ?", userId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.Research{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *MysqlDatabase) GetDimensionByResearchId(researchId int) ([]*model.Dimension, error) {
	result := make([]*model.Dimension, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT `dimension_id`, `research_id`, `name`, `variable`, `status`, `created_at`, `updated_at` FROM `dimension` WHERE `research_id` = ?", researchId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.Dimension{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *MysqlDatabase) CreateDimension(dimension model.Dimension) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO `dimension`( `research_id`, `name`, `variable`) VALUES (?,?,?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(dimension.ResearchId, dimension.Name, dimension.Variable)
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
func (m *MysqlDatabase) UpdateDimension(dimension model.Dimension) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.Prepare("UPDATE `dimension` SET `research_id`= ?,`name`= ?,`variable`= ?,`status`=?  WHERE `dimension_id` = ?")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(dimension.ResearchId, dimension.Name, dimension.Variable, dimension.Status, dimension.DimensionId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.RowsAffected()
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
func (m *MysqlDatabase) CreateResearch(research model.Research) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO `research`( `researcher_id`, `expert_id`, `title`, `speciality`, `authors`, `observation`, `attachment_one`, `attachment_two`) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(research.ResearcherId, research.ExpertId, research.Title, research.Speciality, research.Authors, research.Observation, research.AttachmentOne, research.AttachmentTwo)
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
func (m *MysqlDatabase) DeleteDimension(dimensionId int) (*int64, error) {

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("DELETE FROM `dimension` WHERE `dimension_id` = ?")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(dimensionId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.RowsAffected()
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
func (m *MysqlDatabase) UpdateResearchStatus(research model.Research) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE `research` SET  `observation` = ? , `status` = ? where `research_id` = ?")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(research.Observation, research.Status, research.ResearchId)
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
func (m *MysqlDatabase) UpdateResearch(research model.Research) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE `research` SET `researcher_id`=?,`expert_id`=?,`title`=?,`speciality`=?,`authors`=?,`observation`=?,`attachment_one`=?,`attachment_two`=?,`status`=? WHERE `research_id`=?")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(research.ResearcherId, research.ExpertId, research.Title, research.Speciality, research.Authors, research.Observation, research.AttachmentOne, research.AttachmentTwo, research.Status, research.ResearchId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.RowsAffected()
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
func (m *MysqlDatabase) UpdateDimensionStatus(dimension model.Dimension) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE `dimension` SET  `status` = ? WHERE `dimension_id` = ?")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(dimension.Status, dimension.DimensionId)
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
func (m *MysqlDatabase) GetResearchById(researchId int) (*model.Research, error) {

	fmt.Println(researchId)
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
		tx.Rollback()
	}

	item := &model.Research{}

	rows := tx.QueryRow("SELECT `research_id`, `researcher_id`, `expert_id`, `title`, `speciality`, `authors`, `observation`, `attachment_one`, `attachment_two`, `status`, `created_at`, `updated_at` FROM `research` WHERE  `research_id` = ?", researchId)

	if rows != nil {
		err := rows.Scan(ScanRow(item)...)
		fmt.Println(err)
		if err != nil {
			return nil, err
			tx.Rollback()
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return item, nil
}
func (m *MysqlDatabase) GetCriterioByExpert(speciality string, expertId int) ([]*model.Criterio, error) {
	result := make([]*model.Criterio, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT `criterio_id`, `name`, `speciality`, `expert_id`, `created_at`, `updated_at` FROM `criterio` WHERE `speciality` = ?  ", speciality)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.Criterio{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *MysqlDatabase) CreateCriterioResponse(criterioResponse model.CriterioResponse) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO `criterio_response`( `research_id`, `dimension_id`, `status`, `criterio_id`) VALUES (?,?,?,?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(criterioResponse.ResearchId, criterioResponse.DimensionId, criterioResponse.Status, criterioResponse.CriterioId)
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
func (m *MysqlDatabase) UpdateCriterioResponse(criterioResponse model.CriterioResponse) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE `criterio_response` SET `research_id`=?,`dimension_id`=?,`status`=? , `criterio_id` = ? WHERE  `criterio_response_id` = ?")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(criterioResponse.ResearchId, criterioResponse.DimensionId, criterioResponse.Status, criterioResponse.CriterioId, criterioResponse.CriterioResponseId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.RowsAffected()
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
func (m *MysqlDatabase) DeleteResearch(researchId int) (*int64, error) {

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("DELETE FROM `research` WHERE `research_id` = ?")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(researchId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.RowsAffected()
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
func (m *MysqlDatabase) GetResearchByExpert(userId int, status int) ([]*model.Research, error) {
	result := make([]*model.Research, 0)
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("SELECT `research_id`, `researcher_id`, `expert_id`, `title`, `speciality`, `authors`, `observation`, `attachment_one`, `attachment_two`, `status`, `created_at`, `updated_at` FROM `research` WHERE  `expert_id` = ? and `status` in (?,5,6)", userId, status)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		items := &model.Research{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *MysqlDatabase) GetCriterioResponseByResearchId(researchId int) ([]*model.CriterioResponse, error) {
	result := make([]*model.CriterioResponse, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT `criterio_response_id`, `criterio_id`, `research_id`, `dimension_id`, `status`, `created_at`, `updated_at` FROM `criterio_response` WHERE  `research_id` = ?", researchId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.CriterioResponse{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *MysqlDatabase) FetchNetworkRequestByUserId(userId int) ([]*model.User, error) {
	status := 1
	expert := "U"
	result := make([]*model.User, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT DISTINCT `user`.* FROM `user` inner join `network` on `user`.`user_id` = `network`.`user_base_id` where `network`.`user_relation_id` = ? and `user`.`status` = ? and `user`.`role`= ?", userId, status, expert)
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

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *MysqlDatabase) FetchNetworkRequestByExpertId(userId int) ([]*model.User, error) {
	status := 1
	expert := "U"
	result := make([]*model.User, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT DISTINCT `user`.* FROM `user` inner join `network_request` on `user`.`user_id` = `network_request`.`user_base_id` where `network_request`.`user_relation_id` = ? and `user`.`status` = ? and `user`.`role`= ? and `network_request`.`status` = ?", userId, status, expert, status)
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

		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *MysqlDatabase) CreateNetwork(network model.Network) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO `network`( `user_base_id`, `user_relation_id`) VALUES (?,?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(network.UserBaseId, network.UserRelationId)
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
func (m *MysqlDatabase) UpdateNetworkRequest(networkRequest model.NetworkRequest) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE `network_request` SET `status`= ? WHERE `user_base_id` = ? and `user_relation_id` = ? ")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(networkRequest.Status, networkRequest.UserBaseId, networkRequest.UserRelationId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.RowsAffected()
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

func (m *MysqlDatabase) CreateResourceUser(resourceUser model.ResourceUser) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO `resource_user`(`expert_id`, `title`, `subtitle`, `link`) VALUES (?,?,?,?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(resourceUser.ExpertId, resourceUser.Title, resourceUser.Subtitle, resourceUser.Link)
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
func (m *MysqlDatabase) DeleteResourceUser(resourceUserId int) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("DELETE FROM `resource_user` WHERE `resource_user_id`= ? ")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(resourceUserId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.RowsAffected()
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
func (m *MysqlDatabase) FetchAllResourceUserById(userId int) ([]*model.ResourceUser, error) {

	result := make([]*model.ResourceUser, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT `resource_user_id`, `expert_id`, `title`, `subtitle`, `link`, `created_at`, `updated_at` FROM `resource_user` WHERE `expert_id`= ?", userId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.ResourceUser{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (m *MysqlDatabase) FetchCreateSigning(signing model.Signing) (*int64, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO `signing`(`expert_id`, `link`) VALUES (?,? )")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(signing.ExpertId, signing.Link)
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
func (m *MysqlDatabase) FetchUpdateSigning(signing model.Signing) (*int64, error) {

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE `signing` SET `link`= ? WHERE `signing_id` = ? ")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(signing.Link, signing.SigningId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, err := res.RowsAffected()
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

func (m *MysqlDatabase) FetchGetSigning(expertId int) ([]*model.Signing, error) {

	result := make([]*model.Signing, 0)

	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query("SELECT `signing_id`,`expert_id`, `link`,`created_at`, `updated_at` FROM `signing` WHERE `expert_id`= ?", expertId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		items := &model.Signing{}
		// Scan the result into the column pointers...

		err = rows.Scan(ScanRow(items)...)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		result = append(result, items)
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return result, nil
}
