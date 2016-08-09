package gocron

import "errors"

var (
	ErrSecond        error = errors.New("gocron: Invalid second")
	ErrMinute              = errors.New("gocron: Invalid minute")
	ErrHour                = errors.New("gocron: Invalid hour")
	ErrDay                 = errors.New("gocron: Invalid day")
	ErrValueStart          = errors.New("gocron: Invalid start value")
	ErrValueInterval       = errors.New("gocron: Invalid interval value")
)
