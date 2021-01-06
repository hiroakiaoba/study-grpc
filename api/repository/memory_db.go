package repository

import (
	"api/model"
	"log"
	"sync"
)

var memoryDB *db

type db struct {
	userDB *userDB
}

type userDB struct {
	sync.Mutex
	data      map[string]*model.User // map[id]User
	idCounter int32
}

func (u *userDB) add(user *model.User) {
	u.Lock()
	user.ID = u.idCounter
	u.idCounter++
	u.data[user.LoginName] = user
	u.Unlock()
}

func (u *userDB) findByLoginName(loginName string) *model.User {
	u.Lock()
	user := u.data[loginName]
	u.Unlock()
	return user
}

func (u *userDB) findByID(id int32) *model.User {
	u.Lock()
	for _, user := range u.data {
		if user.ID == id {
			return user
		}
	}
	u.Unlock()
	return nil
}

func (u *userDB) list() []*model.User {
	u.Lock()
	users := make([]*model.User, 0)
	for _, user := range u.data {
		users = append(users, user)
	}
	u.Unlock()
	return users
}

func init() {
	mockUser1 := &model.User{
		ID:        1,
		LoginName: "PiyoUser",
		Password:  "password",
	}
	mockUser2 := &model.User{
		ID:        2,
		LoginName: "HogeUser",
		Password:  "password",
	}
	userData := make(map[string]*model.User)
	userData[mockUser1.LoginName] = mockUser1
	userData[mockUser2.LoginName] = mockUser2

	log.Println("setuped memoryDb!")

	memoryDB = &db{
		userDB: &userDB{
			data:      userData,
			idCounter: 3,
		},
	}
}
