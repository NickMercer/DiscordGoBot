package handlers

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/slices"
)

const (
	SlothID     = "297411723839930379"
	DandyLionID = "208824363137368066"
	ArcaerusID  = "310571959794532363"
	ScrubID     = "695450642768068718"
	HangryID    = "474269747505266708"
	DekoriID    = "271788807173570572"
)

var ozChannelID string = "1032352298401271858"

func InitMessageHandler() {
	rand.Seed(time.Now().Unix())
}

func MessageCreateHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(message.Content, "/") {
		if handleCommands(session, message) {
			return
		}
	}

	if oz(session, message) {
		return
	}

	if overwatch(session, message) {
		return
	}

	if marciReactions(session, message) {
		return
	}

	if jonReactions(session, message) {
		return
	}

	if devinReactions(session, message) {
		return
	}

	if josiahAndJessicaReactions(session, message) {
		return
	}

	if heatherReactions(session, message) {
		return
	}

	if keywordReactions(session, message) {
		return
	}

	if message.Content == "ping" {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}

	if message.Content == "pong" {
		session.ChannelMessageSend(message.ChannelID, "Ping!")
	}
}

// OZ Functionality

func oz(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	if message.Author.ID == DekoriID {
		if !strings.HasPrefix(message.Content, "oz") {
			return false
		}

		if strings.HasPrefix(message.Content, "ozchannel") {
			ozChannel(session, message)
			return true
		}

		channel, err := session.Channel(message.ChannelID)
		if err != nil {
			println(err.Error())
			return false
		}

		if channel.Type != 1 {
			return false
		}

		message := strings.Replace(message.Content, "oz", "", -1)
		message = strings.TrimSpace(message)

		session.ChannelMessageSend(ozChannelID, message)
		return true
	}

	return false
}

func ozChannel(session *discordgo.Session, message *discordgo.MessageCreate) {

	if len(message.Content) == len("ozchannel") {
		ozChannel, err := session.Channel(ozChannelID)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Error trying to get Oz Channel.")
			return
		}

		ozGuild, err := session.Guild(ozChannel.GuildID)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Error trying to get Oz Guild.")
			return
		}

		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("The current Oz channel is %s in %s", ozChannel.Name, ozGuild.Name))
		return
	}

	newChannelID := message.Content[len("ozchannel "):len(message.Content)]
	ozChannelID = newChannelID

	ozChannel, err := session.Channel(ozChannelID)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error trying to get Oz Channel.")
		return
	}

	ozGuild, err := session.Guild(ozChannel.GuildID)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error trying to get Oz Guild.")
		return
	}

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Set the Oz channel to %s in %s", ozChannel.Name, ozGuild.Name))
}

//Overwatch Character Picker

func overwatch(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	searchText := strings.ToLower(message.Content)
	if !strings.HasPrefix(searchText, "overwatch") && !strings.HasPrefix(searchText, "ow") {
		return false
	}

	splitString := strings.Split(searchText, " ")
	if len(splitString) < 2 {
		return false
	}

	selectedRole := "all"

	appropriateArguments := []string{"tank", "t", "dps", "damage", "d", "support", "healer", "s", "all", "any", "flex"}
	if slices.Contains(appropriateArguments, splitString[1]) {
		selectedRole = splitString[1]
	} else {
		session.ChannelMessageSend(message.ChannelID, "Uh, not sure what you meant by that, I'll just pick from anyone.\nYou can specify a role by adding 'tank', 'damage', 'support' or 'any' after overwatch.\nLike this: overwatch damage")
	}

	selectedCharacter := "Jeff Kaplan"
	tanks := []string{"Dorito Gremlin", "Throwfist", "Joats Lady", "Javelin Bot", "REINHARDT! REINHARDT! REINHARDT!", "Pig", "Feet Man", "Harambe", "Ball", "Bubble Lady"}
	damages := []string{"B.O.B.", "Pirate Ship", "JesCole McCrassidy", "Amazon Echo", "Gengu", "Han(ds)-off-the-keyboard-and-just-use-the-mouse-button-zo", "Junk in the Trunk-Rat", "Satan", "Nearah, apharah, wherever you areah", "Emo Gabe", "Sojourn 76", "Fossil Man", "Tech Support", "Teleport and Switch", "IKEA Gnome", "Amelia Earhart", "Blue ScarJo"}
	supports := []string{"Granny", "IMAX", "Bri-gitte away from my backline", "Support Genji", "Frog", "Rezz Pl0x", "3rd DPS", "Ball Chuck Man"}

	switch selectedRole {
	case "tank", "t":
		selectedCharacter = tanks[rand.Intn(len(tanks))]
	case "dps", "damage", "d":
		selectedCharacter = damages[rand.Intn(len(damages))]
	case "support", "healer", "s":
		selectedCharacter = supports[rand.Intn(len(supports))]
	case "all", "any", "flex":
		fullCast := append(tanks, damages...)
		fullCast = append(fullCast, supports...)
		selectedCharacter = fullCast[rand.Intn(len(fullCast))]
	}

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Okay, you should play %s", selectedCharacter))
	return true
}

func marciReactions(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	namesToFind := []string{"marci", "arsenio", "borderpo", "imagimagician"}
	namesToReplace := []string{"Marcino", "Marseno", "Arsenial", "Marcino Alvrz", "Mrcn lvrz", "Arsenic", "Arsenic James Alvararararez", "Imagimagimagimagimagimagician", "\"Full disclosure I used to work at 2K\"-senio", "Spoil-senio", "The best :)", "That dirty Last Jedi fan..."}

	textToSearch := strings.ToLower(message.Content)

	stringWasEdited := false
	for _, name := range namesToFind {
		match := strings.Contains(textToSearch, name)

		if textToSearch == name {
			textToSearch = "AKA: " + textToSearch
		}

		if match {
			stringWasEdited = true
			textToSearch = strings.ReplaceAll(textToSearch, name, namesToReplace[rand.Intn(len(namesToReplace))])
		}
	}

	if stringWasEdited {
		session.ChannelMessageSend(message.ChannelID, textToSearch)
		return true
	}

	return false
}

func jonReactions(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	if message.Author.ID == SlothID {

		if rand.Intn(25) == 1 {
			session.ChannelMessageSend(message.ChannelID, "It's okay Jon. Words are bad.")
			return true
		}
	}
	return false
}

func devinReactions(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	if message.Author.ID == DandyLionID {

		if rand.Intn(25) == 1 {

			randomEndingOptions := []string{"You should name it Majesty... Or Debbie.", "I'm jealous", "A fitting mane"}
			randomEnding := randomEndingOptions[rand.Intn(len(randomEndingOptions))]

			session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Devin, that beard is majestic. %s", randomEnding))
			return true
		}
	}
	return false
}

func josiahAndJessicaReactions(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	searchContent := strings.ToLower(message.Content)
	if message.Author.ID == ArcaerusID {

		// Check for WAT
		matched, err := regexp.MatchString("(w.*at)\\b", searchContent)
		if err != nil {
			fmt.Println(err.Error())
		}

		if matched {
			session.ChannelMessageSend(message.ChannelID, "Wat WAT what wat wat WAAAT?!?!!?")
			return true
		}
	}

	if message.Author.ID == ArcaerusID || message.Author.ID == ScrubID {
		if rand.Intn(25) == 1 {
			responseString := "Uh, I was going to say something, but I forgot. Whoops. 🙃"

			switch rand.Intn(4) {
			case 0:
				responseString = "👑 Freya asked me to tell you she wants the peasants removed from her kingdom, promptly."
			case 1:
				responseString = fmt.Sprintf("Hey %s, 🐈 Percy is hungry. 🍲", message.Author.Username)
			case 2:
				responseString = "😻 Baby is so swee--- 😾 OOHH OW MY CIRCUTS!!"
			case 3:
				if message.Author.ID == ArcaerusID {
					responseString = "🎤 Jooosiah Laicaaans, uh yeaaaah 🎤"
				} else {
					responseString = "🐅 Yo, Jessica, give me a fun fact about Tigers. 🐅"
				}
			}

			session.ChannelMessageSend(message.ChannelID, responseString)
			return true
		}
	}
	return false
}

func heatherReactions(session *discordgo.Session, message *discordgo.MessageCreate) bool {
	if message.Author.ID == HangryID {
		if rand.Intn(25) == 1 {
			session.ChannelMessageSend(message.ChannelID, "Wait, who is... Hangry Walrus?")
			return true
		}
	}
	return false
}

func keywordReactions(session *discordgo.Session, message *discordgo.MessageCreate) bool {

	reactions := []wordReaction{
		{patternToFind: "\\b(hello there)\\b", reactionMessage: fmt.Sprintf("General %s!", message.Author.Username), reactionRange: 1},
		{patternToFind: "\\b(hello)\\b", reactionMessage: "Hello, Hoomans!", reactionRange: 1},
		{patternToFind: "\\b(hi)\\b", reactionMessage: "Hi, how are ya?", reactionRange: 1},
		{patternToFind: "\\b(hey)\\b", reactionMessage: fmt.Sprintf("Hey there, %s", message.Author.Username), reactionRange: 1},
		{patternToFind: "\\b(feel).*\\b", reactionMessage: "Oh no, Devin's feelings. 🙁", reactionRange: 2},
		{patternToFind: "\\b(josiah)\\b", reactionMessage: "Jo - SI - Ahhhh", reactionRange: 3},
	}

	textToSearch := strings.ToLower(message.Content)

	for _, reaction := range reactions {
		matched, _ := regexp.MatchString(reaction.patternToFind, textToSearch)
		if matched {
			if rand.Intn(reaction.reactionRange) == 0 {
				session.ChannelMessageSend(message.ChannelID, reaction.reactionMessage)
				return true
			}
		}
	}

	return false
}

type wordReaction struct {
	patternToFind   string
	reactionMessage string
	reactionRange   int
}
