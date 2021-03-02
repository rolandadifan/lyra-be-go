package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserFormat(user User) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return formatter
}

type TokenFormatter struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

func TokenFormate(token Token) TokenFormatter {
	formatter := TokenFormatter{
		ID:     token.ID,
		UserID: token.UserID,
		Token:  token.Token,
	}
	return formatter
}
