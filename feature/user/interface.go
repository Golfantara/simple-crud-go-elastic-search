package user

import (
	"elasticsearch/feature/user/dtos"
	"net/http"
)

type Repository interface {
	Paginate(page, size int) []User
	Insert(user *User) error
	FindByID(userID int) *User
	Update(user User) int64
	DeleteByID(userID int) int64
}

type RepositoryElasticsearch interface {
	Save(user User) error
    FindByID(id string) (User, error)
    SearchUsers(query string) ([]User, error)
    Delete(id string) error
    Update(id string, user User) error
}

type Usecase interface {
    CreateUser(newData dtos.InputUser) (*dtos.ResUser, error)
    GetUserByID(id string) (User, error)
    SearchUsers(query string) ([]User, error)
    DeleteUser(id string) error
    UpdateUser(id string, user User) error
}

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
    GetUserDetails(w http.ResponseWriter, r *http.Request)
    SearchUsers(w http.ResponseWriter, r *http.Request)
    DeleteUser(w http.ResponseWriter, r *http.Request)
    UpdateUser(w http.ResponseWriter, r *http.Request)
}