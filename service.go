package main

import (
	"errors"
	"strings"
)

// StringService interface
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// implementation
type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when the input string is empty
var ErrEmpty = errors.New("Empty String")
