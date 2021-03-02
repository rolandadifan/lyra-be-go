package user

type Register struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type TokenInput struct {
	UserID int    `json:"user_id" binding:"required"`
	Token  string `json:"token" binding:"required"`
}
