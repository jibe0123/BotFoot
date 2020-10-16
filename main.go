package main

import (
	"fmt"
	"github.com/jibe0123/discordBot/app"
	"github.com/jibe0123/discordBot/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	goBot := app.Start()

	// Wait here until CTRL-C or other term signal is received.
	log.Print("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Println("Shutting down server...")

	err = goBot.Close()
	if err != nil {
		log.Print(err)
	}
	log.Println("Server exiting")
}
