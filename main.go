package main

import (
	"discordgobot/handlers"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := getBotToken()
	if token == "" {
		return
	}

	discord, err := discordgo.New("Bot " + token)
	if checkError("error during auth", err) {
		return
	}

	err = discord.Open()
	if checkError("error opening connection,", err) {
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	registerHandlers(discord)

	//Wait until close request
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

func getBotToken() string {
	tokenBytes, err := os.ReadFile("token.txt")
	if checkError("error getting bot token", err) {
		return ""
	}

	token := string(tokenBytes)
	return token
}

func checkError(errorMessage string, err error) bool {
	if err != nil {
		fmt.Println(errorMessage, err)
		return true
	}

	return false
}

func registerHandlers(discord *discordgo.Session) {
	println("handlers registered")
	discord.AddHandler(handlers.MessageCreateHandler)
	handlers.InitMessageHandler()
}
