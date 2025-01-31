package ldaphandle

import (
	"log"

	ldap "github.com/vjeantet/ldapserver"
)

type UseCases interface {
	AuthenticateUser(username, password string) (bool, error)
}

func HandleBind(usecases UseCases) ldap.HandlerFunc {

	return func(w ldap.ResponseWriter, m *ldap.Message) {
		r := m.GetBindRequest()
		res := ldap.NewBindResponse(ldap.LDAPResultSuccess)

		log.Printf("Bind User=%s, Pass=%s", string(r.Name()), string(r.AuthenticationSimple()))

		success, err := usecases.AuthenticateUser(string(r.Name()), string(r.AuthenticationSimple()))

		if err != nil || !success {
			log.Printf("Bind failed User=%s", string(r.Name()))
			res.SetResultCode(ldap.LDAPResultInvalidCredentials)
			res.SetDiagnosticMessage("invalid credentials")
			w.Write(res)
		}

		w.Write(res)

	}
}
