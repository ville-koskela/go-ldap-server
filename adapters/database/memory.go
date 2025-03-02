package database

import (
	"errors"

	"github.com/ville-koskela/go-ldap-server/domain"
)

type InMemoryDatabase struct {
	users map[string]domain.User
}

func NewInMemoryDatabase() *InMemoryDatabase {
	db := &InMemoryDatabase{
		users: make(map[string]domain.User),
	}
	return db
}

func (db *InMemoryDatabase) AddUser(user domain.User) (domain.User, error) {
	db.users[user.Username] = user
	return db.users[user.Username], nil
}

func (db *InMemoryDatabase) FindUserByUsername(username string) (domain.User, error) {
	user, ok := db.users[username]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (db *InMemoryDatabase) Close() error {
	return nil
}
