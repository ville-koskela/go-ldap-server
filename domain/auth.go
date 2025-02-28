package domain

func (uc *UseCases) AuthenticateUser(username string, password string) bool {
	user, err := uc.repo.FindUserByName(username)
	if err != nil {
		return false
	}

	return uc.pwManager.ComparePassword(user.Password, password)
}
