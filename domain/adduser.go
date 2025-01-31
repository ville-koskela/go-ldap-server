package domain

func (uc *UseCases) AddUser(user User) error {
	return uc.db.AddUser(user)
}
