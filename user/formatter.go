package user

func FormatterOutput(user User, token string) UserOutput {
	userOutput := UserOutput{
		Id:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Token:    token,
	}

	return userOutput
}
