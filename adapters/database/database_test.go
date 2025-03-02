package database

import (
	"testing"

	"github.com/ville-koskela/go-ldap-server/domain"
	"github.com/ville-koskela/go-ldap-server/test"
)

var dbTypes = []string{"inmemory"}

func TestDatabase_AddUser(t *testing.T) {
	for _, dbType := range dbTypes {
		t.Run(dbType, func(t *testing.T) {
			db, _ := InitializeDatabase(dbType)
			user, err := db.AddUser(domain.User{
				Username: "testuser",
				Password: "password",
				Email:    "",
				FullName: "",
				UID:      0,
				GID:      0,
			})

			test.Assert(t, nil, err, "User should be added")
			test.Assert(t, "testuser", user.Username, "Username should be testuser")
		})
	}
}
