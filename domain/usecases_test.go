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
	mockUserRepository := NewMockUserRepository()
	mockPasswordManager := NewMockPasswordManager()
	useCases := NewUseCases(mockUserRepository, mockPasswordManager)

	mockUserRepository.setFindUserByUsernameResult(User{Username: "testuser", Password: "drowssap"}, nil)

	authenticated := useCases.AuthenticateUser("testuser", "password")
	test.Assert(t, true, authenticated, "User should be authenticated")

	mockUserRepository.setFindUserByUsernameResult(User{}, errors.New("user not found"))

	authenticated = useCases.AuthenticateUser("nonexistinguser", "password")
	test.Assert(t, false, authenticated, "User should not be authenticated")
}

func TestAddUser(t *testing.T) {
	mockUserRepository := NewMockUserRepository()
	mockPasswordManager := NewMockPasswordManager()
	useCases := NewUseCases(mockUserRepository, mockPasswordManager)

	mockUserRepository.setAddUserError(nil)

	err := useCases.AddUser(User{Username: "newuser", Password: "password"})
	test.Assert(t, nil, err, "User should be added")

	mockUserRepository.setAddUserError(errors.New("error adding user"))

	err = useCases.AddUser(User{Username: "newuser", Password: "password"})
	test.Assert(t, errors.New("error adding user"), err, "User should not be added")
}

/**
 * Mock the database for domain
 */
type MockUserRepository struct {
	addUserError   error
	findUserResult User
	findUserError  error
}

func (m *MockUserRepository) setAddUserError(err error) {
	m.addUserError = err
}

func (m *MockUserRepository) AddUser(user User) error {
	return m.addUserError
}

func (m *MockUserRepository) setFindUserByUsernameResult(user User, err error) {
	m.findUserResult = user
	m.findUserError = err
}

func (m *MockUserRepository) FindUserByUsername(username string) (User, error) {
	return m.findUserResult, m.findUserError
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		addUserError:   nil,
		findUserResult: User{},
		findUserError:  nil,
	}
}

type MockPasswordManager struct{}

func (m *MockPasswordManager) HashPassword(password string) (string, error) {
	hashedPassword := ""
	for _, char := range password {
		hashedPassword = string(char) + hashedPassword
	}
	return hashedPassword, nil
}

func (m *MockPasswordManager) ComparePassword(hashedPassword string, password string) bool {
	comparablePassword, err := m.HashPassword(password)

	if err != nil {
		return false
	}

	return hashedPassword == comparablePassword
}

func NewMockPasswordManager() *MockPasswordManager {
	return &MockPasswordManager{}
}
