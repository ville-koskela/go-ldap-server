package domain

func (uc *UseCases) ListUsers() ([]User, error) {
	return uc.repo.ListUsers()
}
