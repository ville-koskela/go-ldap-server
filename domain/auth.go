package domain

func (uc *UseCases) AuthenticateUser(username string, password string) (bool, error) {
	user, err := uc.db.FindUserByName(username)
	if err != nil {
		return false, err
	}

	if user.Password == password {
		return true, nil
	} else {
		return false, nil
	}
}
