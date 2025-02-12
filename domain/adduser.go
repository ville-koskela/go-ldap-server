package domain

func (uc *UseCases) AddUser(user User) error {
	hash, err := uc.pw.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hash

	return uc.db.AddUser(user)
}
