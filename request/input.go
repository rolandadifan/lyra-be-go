package request

type RequestInput struct {
	Name    string `json:"name" binding:"required"`
	Request string `json:"request" binding:"required"`
}

type RequestDetail struct {
	ID int `uri:"id" binding:"required"`
}
