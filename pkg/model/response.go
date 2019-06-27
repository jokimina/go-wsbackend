package model

type Response struct {
	Status int64       `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrResponse struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
}
