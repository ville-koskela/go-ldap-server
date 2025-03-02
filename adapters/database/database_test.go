package database

import (
	"sort"
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

var testUser2 = domain.User{
	Username: "testuser2",
	Password: "password",
	Email:    "",
	FullName: "",
	UID:      1,
	GID:      1,
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

func TestDatabase_ListUsers(t *testing.T) {
	for _, dbType := range dbTypes {
		t.Run(dbType, func(t *testing.T) {
			db, _ := InitializeDatabase(dbType)
			db.AddUser(testUser)
			db.AddUser(testUser2)

			users, err := db.ListUsers()

			test.Assert(t, nil, err, "Users should be listed")
			test.Assert(t, 2, len(users), "There should be 2 user")
			// Sort users by username to ensure consistent order
			sort.Slice(users, func(i, j int) bool {
				return users[i].Username < users[j].Username
			})
			test.Assert(t, "testuser", users[0].Username, "Username should be testuser")
			test.Assert(t, "testuser2", users[1].Username, "Username should be testuser2")
		})
	}
}
