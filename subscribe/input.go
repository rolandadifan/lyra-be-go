package subscribe

type SubscribeInput struct {
	Email string `json:"email" binding:"required,email"`
}

type SubscribeDetail struct {
	ID int `uri:"id" binding:"required"`
}
