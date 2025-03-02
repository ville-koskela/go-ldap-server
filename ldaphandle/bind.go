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

		log.Printf("Bind User=%s, Pass=%s", string(r.Name()), string(r.AuthenticationSimple()))

		success := authFunc(string(r.Name()), string(r.AuthenticationSimple()))

		if success {
			w.Write(res)
		}

		log.Printf("Bind failed User=%s", string(r.Name()))
		res.SetResultCode(ldap.LDAPResultInvalidCredentials)
		res.SetDiagnosticMessage("invalid credentials")
		w.Write(res)

	}
}
