package cron

import "errors"

var (
	ERR_INVALID_SECOND = errors.New("Invalid second")
	ERR_INVALID_MINUTE = errors.New("Invalid minute")
	ERR_INVALID_HOUR   = errors.New("Invalid hour")
	ERR_INVALID_DAY    = errors.New("Invalid day")
)
