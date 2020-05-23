package files

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type filesRepository struct {
	conn *sql.DB
}

func NewFilesRepository(conn *sql.DB) *filesRepository {
	return &filesRepository{conn: conn}
}

func (files *filesRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.Files)
	return driver.GetById(files.conn, obj, id)
}

func (files *filesRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.Files)
	usr := obj.(model.Files)
	// result, err := driver.Create(files.conn, usr)
	result, err := driver.Create(files.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (files *filesRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	// usr := obj.(*model.Files)
	usr := obj.(model.Files)
	// err := driver.UpdateById(files.conn, usr)
	err := driver.UpdateById(files.conn, &usr)
	return obj, err
}

func (files *filesRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.Files{Id: id}
	return driver.DeleteById(files.conn, obj, id)
}

func (files *filesRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.Files{}
	return driver.GetAll(files.conn, obj, 0, 0)
}

func (files *filesRepository) GetFilesByPID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.Usergroup)
	// return driver.GetGroupById(usergroup.conn, obj, id)
	// return driver.GetUserByGId(files.conn, obj, id)
	return driver.GetFilesByPId(files.conn, obj, id)

}
