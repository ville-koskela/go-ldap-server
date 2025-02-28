package domain

type UserRepository interface {
	FindUserByName(username string) (User, error)
	AddUser(user User) error
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
