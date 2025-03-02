package ldaphandle

import (
	"log"

	"github.com/ville-koskela/go-ldap-server/domain"

	"fmt"

	ldapMessage "github.com/lor00x/goldap/message"
	ldap "github.com/vjeantet/ldapserver"
)

type ListUsersFunc func() ([]domain.User, error)

func HandleSearch(listUsers ListUsersFunc) ldap.HandlerFunc {
	return func(w ldap.ResponseWriter, m *ldap.Message) {
		r := m.GetSearchRequest()

		users, err := listUsers()
		if err != nil {
			log.Printf("Error listing users: %s", err)
			res := ldap.NewSearchResultDoneResponse(ldap.LDAPResultOperationsError)
			w.Write(res)
			return
		}

		for _, user := range users {
			entry := ldap.NewSearchResultEntry("cn=" + user.Username + "," + string(r.BaseObject()))
			entry.AddAttribute("uid", ldapMessage.AttributeValue(user.Username))
			entry.AddAttribute("uidNumber", ldapMessage.AttributeValue(fmt.Sprint(user.UID)))
			//entry.AddAttribute("uid", user.Username)
			//entry.AddAttribute("uidNumber", user.UID)
			//entry.AddAttribute("gidNumber", user.GID)
			//entry.AddAttribute("homeDirectory", "/home/" + user.Username)
			//entry.AddAttribute("loginShell", "/bin/bash")
			w.Write(entry)
		}

		res := ldap.NewSearchResultDoneResponse(ldap.LDAPResultSuccess)

		log.Printf("Search Base=%s, Filter=%s", r.BaseObject(), r.Filter())

		w.Write(res)
	}
}
