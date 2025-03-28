package ldaphandle

import (
	ldap "github.com/vjeantet/ldapserver"

)

func HandleWhoami() ldap.HandlerFunc {

	return func(w ldap.ResponseWriter, m *ldap.Message) {
		// WhoAmI essentially does nothing
		res := ldap.NewExtendedResponse(ldap.LDAPResultSuccess)
		w.Write(res)
	}
}
