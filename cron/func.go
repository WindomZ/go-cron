package cron

func range0To59(n int) bool {
	return (n >= 0 && n <= 59)
}

func ValidSecond(n int) error {
	if !range0To59(n) {
		return ErrSecond
	}
	return nil
}

func ValidMinute(n int) error {
	if !range0To59(n) {
		return ErrMinute
	}
	return nil
}
