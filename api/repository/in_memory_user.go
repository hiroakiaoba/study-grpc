package repository

import (
	"api/model"
)

type InMemoryUserRepository struct{}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}

func (r *InMemoryUserRepository) Create(user *model.User) error {
	memoryDB.userDB.add(user)
	return nil
}

func (r *InMemoryUserRepository) List() ([]*model.User, error) {
	users := memoryDB.userDB.list()
	return users, nil
}

func (r *InMemoryUserRepository) FindByLoginName(loginName string) (*model.User, error) {
	user := memoryDB.userDB.findByLoginName(loginName)
	return user, nil
}

func (r *InMemoryUserRepository) FindByID(id int32) (*model.User, error) {
	user := memoryDB.userDB.findByID(id)
	return user, nil
}
