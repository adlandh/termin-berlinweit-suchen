package model

type Config struct {
	MainUrl     string
	Verbose     bool
	CheckPeriod int
}

func (c Config) GetConfig() Config {
	return c
}

func (c Config) GetMainUrl() string {
	return c.MainUrl
}

func (c Config) GetVerbose() bool {
	return c.Verbose
}

func (c Config) GetCheckPeriod() int {
	return c.CheckPeriod
}
