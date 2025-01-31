package database

import (
	"github.com/ville-koskela/go-ldap-server/domain"
)

type Database interface {
	domain.Database
	Close() error
}

func InitializeDatabase() (Database, error) {
	return NewInMemoryDatabase(), nil
}
