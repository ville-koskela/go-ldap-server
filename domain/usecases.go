package domain

type Database interface {
	AddUser(user User) error
	FindUserByName(username string) (User, error)
}

type UseCases struct {
	db Database
}

func NewUseCases(db Database) *UseCases {
	return &UseCases{db: db}
}
