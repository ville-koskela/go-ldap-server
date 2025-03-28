package ldaphandle

import (
	"log"

	ldap "github.com/vjeantet/ldapserver"
)

type AuthenticateUserFunc func(username string, password string) bool

func HandleBind(authFunc AuthenticateUserFunc) ldap.HandlerFunc {

	return func(w ldap.ResponseWriter, m *ldap.Message) {
		r := m.GetBindRequest()
		res := ldap.NewBindResponse(ldap.LDAPResultSuccess)

		user := string(r.Name())
		pass := string(r.AuthenticationSimple())

		// @TODO: Don't print password in logs
		log.Printf("Bind User=%s, Pass=%s", user, pass)

		success := authFunc(user, pass)

		if success {
			w.Write(res)
			return
		}

		log.Printf("Bind failed User=%s", string(r.Name()))
		res.SetResultCode(ldap.LDAPResultInvalidCredentials)
		res.SetDiagnosticMessage("invalid credentials")
		w.Write(res)

	}
}
