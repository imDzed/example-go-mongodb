package handler

type UserReq struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

type UserRes struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}
