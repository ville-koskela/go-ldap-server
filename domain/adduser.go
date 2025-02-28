package domain

func (uc *UseCases) AddUser(user User) error {
	hash, err := uc.pwManager.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash

	return uc.repo.AddUser(user)
}
