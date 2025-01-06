package config

type DiscordConfig struct {
	BotToken             string `env:"BOT_TOKEN"`
	TestGuildID          string `env:"TEST_GUILD_ID"`
	RemoveCommandsOnExit bool   `env:"REMOVE_COMMANDS_ON_EXIT"`
}
