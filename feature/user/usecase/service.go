package usecase

import "elasticsearch/feature/user"

type service struct {
	model user.RepositoryElasticsearch
}

func New(model user.RepositoryElasticsearch) user.Usecase {
	return &service{
		model: model,
	}
}


func (svc *service) CreateUser(user user.User) error {
	return svc.model.Save(user)
}


func (svc *service) GetUserByID(id string) (user.User, error) {
	return svc.model.FindByID(id)
}

func (svc *service) SearchUsers(query string) ([]user.User, error) {
    return svc.model.SearchUsers(query)
}

func (svc *service) DeleteUser(id string) error {
    return svc.model.Delete(id)
}

func (svc *service) UpdateUser(id string, user user.User) error {
    return svc.model.Update(id, user)
}