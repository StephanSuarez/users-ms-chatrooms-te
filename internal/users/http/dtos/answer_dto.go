package dtos

type AnswerDTO struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Err        error       `json:"error"`
	Data       interface{} `json:"data"`
}
