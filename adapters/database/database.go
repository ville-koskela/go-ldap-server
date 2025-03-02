package database

import (
	"github.com/ville-koskela/go-ldap-server/domain"
)

type Database interface {
	domain.UserRepository
	Close() error
}

func InitializeDatabase(dbType string) (Database, error) {
	switch dbType {
	case "inmemory":
		return NewInMemoryDatabase(), nil
	default:
		panic("Unknown database type")
	}
}
