package domain

type UserReader interface {
	FindUserByUsername(username string) (User, error)
}

type UserWriter interface {
	AddUser(user User) (User, error)
}

type UserRepository interface {
	UserReader
	UserWriter
	ListUsers() ([]User, error)
}

type PasswordHasher interface {
	HashPassword(password string) (string, error)
}

type PasswordVerifier interface {
	ComparePassword(hashedPassword string, password string) bool
}

type PasswordManager interface {
	PasswordHasher
	PasswordVerifier
}

type UseCases struct {
	repo      UserRepository
	pwManager PasswordManager
}

func NewUseCases(repo UserRepository, pwManager PasswordManager) *UseCases {
	return &UseCases{
		repo:      repo,
		pwManager: pwManager,
	}
}
