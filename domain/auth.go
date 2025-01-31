package domain

func (uc *UseCases) AuthenticateUser(username string, password string) bool {
	user, err := uc.db.FindUserByName(username)
	if err != nil {
		return false
	}

	if user.Password == password {
		return true
	} else {
		return false
	}
}
