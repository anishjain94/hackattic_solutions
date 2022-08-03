package common

type AckDto struct {
	Success bool    `json:"success"`
	Message *string `json:"message"`
}

type SuccessDto struct {
	Meta AckDto      `json:"meta"`
	Data interface{} `json:"data"`
}
