package gocron

import "errors"

var (
	ErrSecond        error = errors.New("Invalid second")
	ErrMinute              = errors.New("Invalid minute")
	ErrHour                = errors.New("Invalid hour")
	ErrDay                 = errors.New("Invalid day")
	ErrValueStart          = errors.New("Invalid start value")
	ErrValueInterval       = errors.New("Invalid interval value")
)
