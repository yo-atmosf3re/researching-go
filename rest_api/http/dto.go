package http

import (
	"encoding/json"
	"errors"
	"time"
)

// DTO == data transfer object, e.g. for receiving data from request
type TaskDTO struct {
	Title       string
	Description string
}

func (t TaskDTO) ValidateForCreate() error { // very simple handling error - create to method for struct DTO which will be to validate fields this DTO
	if t.Title == "" {
		return errors.New("title is empty")
	}
	if t.Description == "" {
		return errors.New("description is empty")
	}
	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func CreateErrorDTO(message string) ErrorDTO {
	return ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}
}
