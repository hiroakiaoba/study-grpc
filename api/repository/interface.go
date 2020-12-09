package repository

import "api/model"

type IUserRepository interface {
	Create(*model.User) error
	List() ([]*model.User, error)
	FindByLoginName(loginName string) (*model.User, error)
}
