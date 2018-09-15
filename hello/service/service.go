package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Service struct {
}

type serializableError struct{ error }

func (s *serializableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Error())
}

func newSerializableError(text string) error {
	return &serializableError{errors.New(text)}
}

func (s Service) Hello(ctx context.Context, firstName string, lastName string) (string, error) {
	firstName = strings.Trim(firstName, "\t\r\n ")
	lastName = strings.Trim(lastName, "\t\r\n ")

	if len(firstName) == 0 && len(lastName) == 0 {
		return "", newSerializableError("missing required name information")
	}
	if len(firstName) == 0 {
		return fmt.Sprintf(
			"Hello Mr./Ms. %s, nice to meet you. Do you have a first name?",
			lastName,
		), nil
	}
	if len(lastName) == 0 {
		return fmt.Sprintf(
			"Hello %s, nice to meet you. Do you have a last name?",
			firstName,
		), nil
	}
	return fmt.Sprintf(
		"Hello %s %s, nice to meet you.",
		firstName, lastName,
	), nil
}
