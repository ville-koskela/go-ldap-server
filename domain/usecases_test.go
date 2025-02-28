package domain

import (
	"errors"
	"testing"

	"github.com/ville-koskela/go-ldap-server/test"
)

/**
 * Tests for domain
 */

func TestAuthenticateUser(t *testing.T) {
	mockDB := NewMockDatabase()
	mockPasswordTool := NewMockPasswordTool()
	useCases := NewUseCases(mockDB, mockPasswordTool)

	mockDB.setFindUserByNameResult(User{Username: "testuser", Password: "drowssap"}, nil)

	authenticated := useCases.AuthenticateUser("testuser", "password")
	test.Assert(t, true, authenticated, "User should be authenticated")

	mockDB.setFindUserByNameResult(User{}, errors.New("user not found"))

	authenticated = useCases.AuthenticateUser("nonexistinguser", "password")
	test.Assert(t, false, authenticated, "User should not be authenticated")
}

func TestAddUser(t *testing.T) {
	mockDB := NewMockDatabase()
	mockPasswordTool := NewMockPasswordTool()
	useCases := NewUseCases(mockDB, mockPasswordTool)

	mockDB.setAddUserError(nil)

	err := useCases.AddUser(User{Username: "newuser", Password: "password"})
	test.Assert(t, nil, err, "User should be added")

	mockDB.setAddUserError(errors.New("error adding user"))

	err = useCases.AddUser(User{Username: "newuser", Password: "password"})
	test.Assert(t, errors.New("error adding user"), err, "User should not be added")
}

/**
 * Mock the database for domain
 */
type MockDatabase struct {
	addUserError   error
	findUserResult User
	findUserError  error
}

func (m *MockDatabase) setAddUserError(err error) {
	m.addUserError = err
}

func (m *MockDatabase) AddUser(user User) error {
	return m.addUserError
}

func (m *MockDatabase) setFindUserByNameResult(user User, err error) {
	m.findUserResult = user
	m.findUserError = err
}

func (m *MockDatabase) FindUserByName(username string) (User, error) {
	return m.findUserResult, m.findUserError
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		addUserError:   nil,
		findUserResult: User{},
		findUserError:  nil,
	}
}

type MockPasswordTool struct{}

func (m *MockPasswordTool) HashPassword(password string) (string, error) {
	hashedPassword := ""
	for _, char := range password {
		hashedPassword = string(char) + hashedPassword
	}
	return hashedPassword, nil
}

func (m *MockPasswordTool) ComparePassword(hashedPassword string, password string) bool {
	comparablePassword, err := m.HashPassword(password)

	if err != nil {
		return false
	}

	return hashedPassword == comparablePassword
}

func NewMockPasswordTool() *MockPasswordTool {
	return &MockPasswordTool{}
}
