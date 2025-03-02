package database

import (
	"testing"

	"github.com/ville-koskela/go-ldap-server/domain"
	"github.com/ville-koskela/go-ldap-server/test"
)

var dbTypes = []string{"inmemory", "sqlite3"}

var testUser = domain.User{
	Username: "testuser",
	Password: "password",
	Email:    "",
	FullName: "",
	UID:      0,
	GID:      0,
}

func TestDatabase_AddUser(t *testing.T) {
	for _, dbType := range dbTypes {
		t.Run(dbType, func(t *testing.T) {
			db, _ := InitializeDatabase(dbType)
			user, err := db.AddUser(testUser)

			test.Assert(t, nil, err, "User should be added")
			test.Assert(t, "testuser", user.Username, "Username should be testuser")
		})
	}
}

func TestDatabase_FindUserByUsername(t *testing.T) {
	for _, dbType := range dbTypes {
		t.Run(dbType, func(t *testing.T) {
			db, _ := InitializeDatabase(dbType)
			db.AddUser(testUser)

			user, err := db.FindUserByUsername("testuser")

			test.Assert(t, nil, err, "User should be found")
			test.Assert(t, "testuser", user.Username, "Username should be testuser")
		})
	}
}
