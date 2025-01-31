package domain

type Database interface {
	AddUser(user User) error
	FindUserByName(username string) (User, error)
}

type Password interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) bool
}

type UseCases struct {
	db Database
	pw Password
}

func NewUseCases(db Database, pw Password) *UseCases {
	return &UseCases{db: db, pw: pw}
}
