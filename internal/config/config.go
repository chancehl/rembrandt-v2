package config

type Config struct {
	TestGuildID          string
	RemoveCommandsOnExit bool
}

func NewConfig(testGuildID string, removeCommandsOnExit bool) *Config {
	return &Config{
		TestGuildID:          testGuildID,
		RemoveCommandsOnExit: removeCommandsOnExit,
	}
}
