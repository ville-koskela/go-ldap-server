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

	defaultUser := domain.User{
		Username: "test",
		Password: "test",
		Email:    "default@example.com",
		FullName: "Default User",
		UID:      10000,
		GID:      10000,
	}
	db.AddUser(defaultUser)

	return db
}

func (db *InMemoryDatabase) AddUser(user domain.User) error {
	db.users[user.Username] = user
	return nil
}

func (db *InMemoryDatabase) FindUserByName(username string) (domain.User, error) {
	user, ok := db.users[username]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (db *InMemoryDatabase) Close() error {
	return nil
}
