package environment

import (
	"os"
	"strconv"
)

const (
	DEFAULT_LDAP_PORT = 10389
	DEFAULT_DB_TYPE   = "sqlite3"
)

type Env struct {
	LdapPort int
	DbType   string
}

func NewEnv() *Env {
	ldapPort, err := strconv.Atoi(os.Getenv("LDAP_PORT"))

	if err != nil {
		ldapPort = DEFAULT_LDAP_PORT
	}

	return &Env{
		LdapPort: ldapPort,
		DbType:   os.Getenv("DB_TYPE"),
	}
}

func (e *Env) GetLdapPort() int {
	if e.LdapPort == 0 {
		return DEFAULT_LDAP_PORT
	}
	return e.LdapPort
}

func (e *Env) GetDbType() string {
	if e.DbType == "" {
		return DEFAULT_DB_TYPE
	}
	return e.DbType
}
