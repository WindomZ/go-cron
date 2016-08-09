package cron

func AddFunc(spec string, cmd func()) error {
	return getCron().AddFunc(spec, cmd)
}
