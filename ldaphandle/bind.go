package ldaphandle

import (
	"log"

	ldap "github.com/vjeantet/ldapserver"
)

type AuthenticateUserFunc func(username string, password string) bool

func HandleBind(authFunc AuthenticateUserFunc) ldap.HandlerFunc {

	return func(w ldap.ResponseWriter, m *ldap.Message) {
		r := m.GetBindRequest()

		dn := string(r.Name())
		pw := string(r.AuthenticationSimple())

		// @TODO: Don't print password in logs
		log.Printf("Bind User=%s, Pass=%s", dn, pw)

		if dn == "" {
			log.Printf("Anonymous bind accepted")
			res := ldap.NewBindResponse(ldap.LDAPResultSuccess)
			w.Write(res)
			return
		}

		success := authFunc(dn, pw)

		if success {
			log.Printf("Authenticated bind for DN: %s", dn)
			res := ldap.NewBindResponse(ldap.LDAPResultSuccess)
			w.Write(res)
			return
		}

		log.Printf("Bind failed for DN: %s", dn)
		res := ldap.NewBindResponse(ldap.LDAPResultInvalidCredentials)
		res.SetDiagnosticMessage("invalid credentials")
		w.Write(res)

	}
}
