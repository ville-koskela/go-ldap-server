package domain_test

import (
	"errors"
	"testing"

	"github.com/ville-koskela/go-ldap-server/domain"
	"github.com/ville-koskela/go-ldap-server/test"
)

type MockDatabase struct {
	users map[string]domain.User
}

func (m *MockDatabase) AddUser(user domain.User) error {
	if _, exists := m.users[user.Username]; exists {
		return errors.New("user already exists")
	}
	m.users[user.Username] = user
	return nil
}

func (m *MockDatabase) FindUserByName(username string) (domain.User, error) {
	user, exists := m.users[username]
	if !exists {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{users: make(map[string]domain.User)}
}

func TestAuthenticateUser(t *testing.T) {
	mockDB := NewMockDatabase()
	useCases := domain.NewUseCases(mockDB)

	user := domain.User{Username: "testuser", Password: "password"}
	mockDB.AddUser(user)

	authenticated := useCases.AuthenticateUser("testuser", "password")
	test.Assert(t, true, authenticated)
}

/*
func TestAddUser(t *testing.T) {
    mockDB := NewMockDatabase()
    useCases := domain.NewUseCases(mockDB)

    user := domain.User{Username: "newuser", Password: "newpassword"}
    err := useCases.AddUser(user)
    test.Assert(t, nil, err)

    storedUser, err := mockDB.FindUserByName("newuser")
    test.Assert(t, nil, err)
    test.Assert(t, user, storedUser)
}

func TestFindUserByName(t *testing.T) {
    mockDB := NewMockDatabase()
    useCases := domain.NewUseCases(mockDB)

    user := domain.User{Username: "existinguser", Password: "password"}
    mockDB.AddUser(user)

    foundUser, err := useCases.FindUserByName("existinguser")
    test.Assert(t, nil, err)
    test.Assert(t, user, foundUser)
}
*/
