package domain

func (uc *UseCases) AuthenticateUser(username string, password string) bool {
	user, err := uc.repo.FindUserByUsername(username)
	if err != nil {
		return false
	}

	return uc.pwManager.ComparePassword(user.Password, password)
}
