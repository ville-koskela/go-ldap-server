package ldaphandle

import (
	"log"
	"strconv"
	"time"

	"github.com/ville-koskela/go-ldap-server/domain"

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

			entry.AddAttribute("objectClass", "posixAccount", "shadowAccount", "inetOrgPerson")
			entry.AddAttribute("cn", ldapMessage.AttributeValue(user.Username))
			entry.AddAttribute("uid", ldapMessage.AttributeValue(user.Username))
			entry.AddAttribute("sn", ldapMessage.AttributeValue(user.FullName))
			entry.AddAttribute("mail", ldapMessage.AttributeValue(user.Email))
			entry.AddAttribute("uidNumber", ldapMessage.AttributeValue(strconv.Itoa(user.UID)))
			entry.AddAttribute("gidNumber", ldapMessage.AttributeValue(strconv.Itoa(user.GID)))
			entry.AddAttribute("homeDirectory", ldapMessage.AttributeValue("/home/"+user.Username))
			entry.AddAttribute("loginShell", "/bin/bash")

			// @TODO: We probably want to keep track of password age?
			shadowLastChange := int(time.Now().Unix() / 86400)
			entry.AddAttribute("shadowLastChange", ldapMessage.AttributeValue(strconv.Itoa(shadowLastChange)))
			entry.AddAttribute("shadowMax", ldapMessage.AttributeValue("99999"))
			entry.AddAttribute("shadowWarning", ldapMessage.AttributeValue("7"))

			w.Write(entry)
		}

		res := ldap.NewSearchResultDoneResponse(ldap.LDAPResultSuccess)

		log.Printf("Search Base=%s, Filter=%s", r.BaseObject(), r.Filter())

		w.Write(res)
	}
}
