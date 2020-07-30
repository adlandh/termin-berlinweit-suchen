package model

type Config struct {
	MainURL     string
	Verbose     bool
	CheckPeriod int
}

func (c Config) GetConfig() Config {
	return c
}

func (c Config) GetMainURL() string {
	return c.MainURL
}

func (c Config) GetVerbose() bool {
	return c.Verbose
}

func (c Config) GetCheckPeriod() int {
	return c.CheckPeriod
}
