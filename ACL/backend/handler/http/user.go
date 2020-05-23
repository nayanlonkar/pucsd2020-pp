package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/pucsd2020-pp/rest-api/handler"
	"github.com/pucsd2020-pp/rest-api/model"
	"github.com/pucsd2020-pp/rest-api/repository"
	"github.com/pucsd2020-pp/rest-api/repository/user"
)

type User struct {
	handler.HTTPHandler
	repo  repository.IRepository
	repo1 repository.JRepository
}

func NewUserHandler(conn *sql.DB) *User {
	return &User{
		repo:  user.NewUserRepository(conn),
		repo1: user.NewUserRepository(conn),
	}
}

func (user *User) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "user/{id}", Func: user.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "user", Func: user.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "user/{id}", Func: user.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "user/{id}", Func: user.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "user", Func: user.GetAll},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "login", Func: user.Login},
	}
}

func (user *User) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = user.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (user *User) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = user.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (user *User) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.User{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in user id for tracking update
		// usr.UpdatedBy = 0

		iUsr, err = user.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.User)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (user *User) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = user.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "User deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (user *User) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := user.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}

func (user *User) Login(w http.ResponseWriter, r *http.Request) {
	var usr1 model.User
	err := json.NewDecoder(r.Body).Decode(&usr1)
	id := usr1.Id
	// fmt.Printf("id is %d\n", id)
	password := usr1.Password
	var usr interface{}
	for {
		if nil != err {
			break
		}

		// usr, err = user.repo.Login(r.Context(), id, password)
		usr, err = user.repo1.Login(r.Context(), id, password)
		break
	}
	if nil != err {
		handler.WriteJSONResponse(w, r, usr, http.StatusNotFound, err)
	} else {
		handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
	}

}
