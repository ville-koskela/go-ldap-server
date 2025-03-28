package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/ville-koskela/go-ldap-server/adapters/database"
	environment "github.com/ville-koskela/go-ldap-server/adapters/env"
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

	env := environment.NewEnv()
	db, _ := database.InitializeDatabase(env.GetDbType())
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
	routes.Search(ldaphandle.HandleSearch(uc.ListUsers))
	routes.Extended(ldaphandle.HandleWhoami()).RequestName(ldap.NoticeOfWhoAmI).Label("Ext - Who Am I")
	server.Handle(routes)

	go server.ListenAndServe("127.0.0.1:" + strconv.Itoa(env.GetLdapPort()))

	// When CTRL+C, SIGINT and SIGTERM signal occurs
	// Then stop server gracefully
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	close(ch)

	db.Close()
	server.Stop()
}
