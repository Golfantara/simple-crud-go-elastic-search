package usecase

import (
	"elasticsearch/feature/user"
	"elasticsearch/feature/user/dtos"
	"fmt"

	"github.com/labstack/gommon/log"
)

type service struct {
	modelElastic user.RepositoryElasticsearch
	model user.Repository
}

func New(modelElastic user.RepositoryElasticsearch, model user.Repository) user.Usecase {
	return &service{
		modelElastic: modelElastic,
		model: model,
	}
}


func (svc *service) CreateUser(newData dtos.InputUser) (*dtos.ResUser, error) {
	data := user.User{
		ID: newData.ID,
		Name: newData.Name,
		Email: newData.Email,
		Address: newData.Address,
	}
	if data.Name == ""{
		log.Error("failed to insert")
	}
	err := svc.model.Insert(&data)
	if err != nil {
		fmt.Printf("usecase: failed to insert data to mysql: %v", err)
	}

	err = svc.modelElastic.Save(data)
	if err != nil {
		fmt.Printf("usecase: failed to save data to elastic: %v", err)
	}

	result := dtos.ResUser{
		ID: data.ID,
		Name: data.Name,
		Email: data.Email,
		Address: data.Address,
	}

	return &result, nil
}


func (svc *service) GetUserByID(id string) (user.User, error) {
	return svc.modelElastic.FindByID(id)
}

func (svc *service) SearchUsers(query string) ([]user.User, error) {
    return svc.modelElastic.SearchUsers(query)
}

func (svc *service) DeleteUser(id string) error {
    return svc.modelElastic.Delete(id)
}

func (svc *service) UpdateUser(id string, user user.User) error {
    return svc.modelElastic.Update(id, user)
}