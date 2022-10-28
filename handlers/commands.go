package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleCommands(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	var contentToSearch = strings.ToLower(message.Content)
	switch contentToSearch {
	case "/joke":
		fetchDadJoke(session, message)
		return true
	}

	return false
}

func fetchDadJoke(session *discordgo.Session, message *discordgo.MessageCreate) {
	response, err := http.Get("http://icanhazdadjoke.com/slack")
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Hmmm, I seem to be fresh out of dad jokes.")
		fmt.Print(err.Error())
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Hmmm, I seem to be fresh out of dad jokes.")
		fmt.Print(err.Error())
		return
	}

	var responseObject DadJokeResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Hmmm, I seem to be fresh out of dad jokes.")
		fmt.Println(err.Error())
		return
	}

	joke := responseObject.Attachments[0].Text
	session.ChannelMessageSend(message.ChannelID, joke)
}

type DadJokeResponse struct {
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Text string `json:"text"`
}
