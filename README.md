## Setup

Easiest way to setup your dev-tools is by using [asdf](https://asdf-vm.com/guide/getting-started.html). Once setup and [go-plugin](https://github.com/asdf-community/asdf-golang) installed, you are good to go. Go to
project root-folder and run command:

```
asdf install
```

## Githooks

You can add `.githooks` directory to your hooks path by running command:

```
git config core.hooksPath .githooks
```

## Formatting the code

To format the code after making changes, run `gofmt -w .` in the project root.

## Running tests

To run tests, simply run:

```
go test ./...
```

in the root of the project. You may add `-v` flag to get more verbose output from the tests if you wish
or `-cover` to get coverage.

## Running app

Build the app with `go build .` and run the app with `./go-ldap-server`.

Test ldapsearch with:

```
ldapsearch -x -H ldap://127.0.0.1:10389 -D "test" -w "test2"
```

You can also try searching without binding:

```
ldapsearch -H ldap://127.0.0.1:10389 -x -b "dc=example,dc=org" "(objectClass=*)"
```

Bind to the LDAP server with:

```
ldapwhoami -x -H ldap://127.0.0.1:10389 -D "test" -w "test2"
```

## Configuring the app

App can be configured using env-variables. Currently supported variables are:

| Variable name | Default   | Variable description                                                     |
| ------------- | --------- | ------------------------------------------------------------------------ |
| LDAP_PORT     | `10389`   | Defines the port the ldap server binds to and starts listening requests. |
| DB_TYPE       | `sqlite3` | Defines the database type used for storing data.                         |
