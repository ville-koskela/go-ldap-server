package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ville-koskela/go-ldap-server/adapters/database"
	"github.com/ville-koskela/go-ldap-server/adapters/password"
	"github.com/ville-koskela/go-ldap-server/domain"
	"github.com/ville-koskela/go-ldap-server/ldaphandle"
	ldap "github.com/vjeantet/ldapserver"
)

/**
 * Main function is more or less of a copy & paste
 * from the example in the ldapserver library.
 */
func main() {

	db, _ := database.InitializeDatabase("sqlite3")
	pw := password.PasswordTool
	uc := domain.NewUseCases(db, pw)

	// add default user
	uc.AddUser(domain.User{
		Username: "test",
		Password: "test",
		Email:    "default@example.com",
		FullName: "Default User",
		UID:      1000,
		GID:      1000,
	})

	//ldap logger
	ldap.Logger = log.New(os.Stdout, "[server] ", log.LstdFlags)

	//Create a new LDAP Server
	server := ldap.NewServer()

	routes := ldap.NewRouteMux()
	routes.Bind(ldaphandle.HandleBind(uc.AuthenticateUser))
	server.Handle(routes)

	// listen on 10389
	go server.ListenAndServe("127.0.0.1:10389")

	// When CTRL+C, SIGINT and SIGTERM signal occurs
	// Then stop server gracefully
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	close(ch)

	server.Stop()
}
