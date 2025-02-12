package domain

type Database interface {
	AddUser(user User) error
	FindUserByName(username string) (User, error)
}

type PasswordTool interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) bool
}

type UseCases struct {
	db Database
	pw PasswordTool
}

func NewUseCases(db Database, pw PasswordTool) *UseCases {
	return &UseCases{db: db, pw: pw}
}
