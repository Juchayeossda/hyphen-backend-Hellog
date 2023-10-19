package model

type SISS struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
	Data    struct {
		ID string `json:"id" `
	} `json:"data"`
}
