package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/chancehl/rembrandt-v2/commands"
	"github.com/chancehl/rembrandt-v2/config"
	"github.com/chancehl/rembrandt-v2/handlers"
	"github.com/joho/godotenv"
)

var botConfig *config.BotConfig

var session *discordgo.Session

var registeredCommands = []*discordgo.ApplicationCommand{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	botConfig = &config.BotConfig{
		TestGuildID:          os.Getenv("TEST_GUILD_ID"),
		RemoveCommandsOnExit: os.Getenv("REMOVE_COMMANDS_ON_EXIT") == "true",
	}
}

func init() {
	// create bot session
	var err error
	session, err = discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf("invalid bot parameters: %v", err)
	}

	// register handlers
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := handlers.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}
	})

	session.AddHandler(handlers.OnBotReadyHandler)
}

func main() {
	log.Printf("starting bot with config %+v\n", *botConfig)

	// start the bot
	if err := session.Open(); err != nil {
		log.Fatalf("cannot open the session: %v", err)
	}
	defer session.Close()

	// register commands
	log.Printf("registering %d bot command(s)...\n", len(commands.AllCommands))
	for _, botCommand := range commands.AllCommands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, botConfig.TestGuildID, botCommand)
		if err != nil {
			log.Fatalf("- cannot register command %s: %v", botCommand.Name, err)
		} else {
			log.Printf("- registered command `/%s` for guild %s\n", cmd.Name, cmd.GuildID)
		}
		registeredCommands = append(registeredCommands, cmd)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("bot has started (press ctrl+c to exit)")
	<-stop

	// cleanup
	if botConfig.RemoveCommandsOnExit {
		log.Printf("removing %d bot command(s)...\n", len(registeredCommands))
		for _, command := range registeredCommands {
			if err := session.ApplicationCommandDelete(session.State.User.ID, botConfig.TestGuildID, command.ID); err != nil {
				log.Fatalf("- failed to remove command `/%s` from guild %s: %v", command.Name, botConfig.TestGuildID, err)
			} else {
				log.Printf("- removed command `/%s` from guild %s\n", command.Name, botConfig.TestGuildID)
			}
		}
	}
	log.Println("bot exited gracefully")
}
