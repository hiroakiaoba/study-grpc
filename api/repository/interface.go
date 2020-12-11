package repository

import "api/model"

type IUserRepository interface {
	Create(*model.User) error
	List() ([]*model.User, error)
	FindByLoginName(loginName string) (*model.User, error)
	FindByID(id int32) (*model.User, error)
}
