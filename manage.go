package gocron

var crons map[string]*Cron

func init() {
	crons = make(map[string]*Cron)
}

func NewCron(tags ...string) *Cron {
	c := newCron(tags...)
	crons[c.Tag] = c
	return c
}

func GetCron(tag string) *Cron {
	if len(tag) == 0 {
		return nil // TODO:return default
	} else if c, ok := crons[tag]; ok {
		return c
	}
	return nil
}
