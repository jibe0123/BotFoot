package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/jibe0123/discordBot/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

type DiscordConfig struct {
	Token string `env:"TOKEN_DISCORD"`
}

func main() {
	dsConfig := DiscordConfig{}
	if err := env.Parse(&dsConfig); err != nil {
		return
	}
	goBot := app.Start(dsConfig.Token)

	// Wait here until CTRL-C or other term signal is received.
	log.Print("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Println("Shutting down server...")

	if err := goBot.Close(); err != nil {
		log.Print(err)
		return
	}
	log.Println("Server exiting")
}
