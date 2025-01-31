package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ville-koskela/go-ldap-server/adapters/database"
	"github.com/ville-koskela/go-ldap-server/domain"
	"github.com/ville-koskela/go-ldap-server/ldaphandle"
	ldap "github.com/vjeantet/ldapserver"
)

/**
 * Main function is more or less of a copy & paste
 * from the example in the ldapserver library.
 */
func main() {

	db, _ := database.InitializeDatabase()
	uc := domain.NewUseCases(db)

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
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	close(ch)

	server.Stop()
}
