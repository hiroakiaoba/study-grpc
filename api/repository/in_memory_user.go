package repository

import (
	"sync"

	"api/model"
)

type UserBox struct {
	sync.Mutex
	data      map[string]*model.User // map[id]User
	idCounter int32
}

func (u *UserBox) Add(user *model.User) {
	u.Lock()
	user.ID = u.idCounter
	u.idCounter++
	u.data[user.LoginName] = user
	u.Unlock()
}

func (u *UserBox) FindByLoginName(loginName string) *model.User {
	u.Lock()
	user := u.data[loginName]
	u.Unlock()
	return user
}

func (u *UserBox) List() []*model.User {
	u.Lock()
	users := make([]*model.User, 0)
	for _, user := range u.data {
		users = append(users, user)
	}
	u.Unlock()
	return users
}

type InMemoryUserRepository struct {
	userBox *UserBox
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	mockUser1 := &model.User{
		ID:        1,
		LoginName: "PiyoUser",
		Password:  "pasword",
	}
	mockUser2 := &model.User{
		ID:        2,
		LoginName: "HogeUser",
		Password:  "pasword",
	}
	userData := make(map[string]*model.User)
	userData[mockUser1.LoginName] = mockUser1
	userData[mockUser2.LoginName] = mockUser2

	return &InMemoryUserRepository{
		userBox: &UserBox{
			data:      userData,
			idCounter: 3,
		},
	}
}

func (r *InMemoryUserRepository) Create(user *model.User) error {
	r.userBox.Add(user)
	return nil
}

func (r *InMemoryUserRepository) List() ([]*model.User, error) {
	users := r.userBox.List()
	return users, nil
}

func (r *InMemoryUserRepository) FindByLoginName(loginName string) (*model.User, error) {
	user := r.userBox.FindByLoginName(loginName)
	return user, nil
}
