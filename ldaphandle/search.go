package ldaphandle

import (
	"log"

	"github.com/ville-koskela/go-ldap-server/domain"

	ldap "github.com/vjeantet/ldapserver"
)

type ListUsersFunc func() ([]domain.User, error)

func HandleSearch(listUsers ListUsersFunc) ldap.HandlerFunc {
	return func(w ldap.ResponseWriter, m *ldap.Message) {
		r := m.GetSearchRequest()
		res := ldap.NewSearchResultDoneResponse(ldap.LDAPResultSuccess)

		log.Printf("Search Base=%s, Filter=%s", r.BaseObject(), r.Filter())

		w.Write(res)
	}
}
