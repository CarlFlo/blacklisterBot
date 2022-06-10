package commands

import (
	"bytes"
	"fmt"
	"image"
	"net/url"

	"github.com/CarlFlo/blacklisterBot/src/database"
	"github.com/CarlFlo/blacklisterBot/src/utils"
	"github.com/CarlFlo/malm"
	"github.com/bwmarrin/discordgo"
)

func Ban(s *discordgo.Session, m *discordgo.MessageCreate, args *[]string) {

	if len(*args) == 1 {
		// No URL provided. Ban the last URL found in the n previous messages message.
		utils.SendMessageFailure(s, m, "No URL provided")
		return
	}

	link := (*args)[1]

	// Parse the link
	_, err := url.Parse(link)
	if err != nil {
		utils.SendMessageFailure(s, m, "Invalid URL")
		return
	}

	b, err := utils.FetchFromURL(&link)
	if err != nil {
		malm.Error("%s", err)
		return
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		utils.SendDirectMessage(s, m, fmt.Sprintf("Ban failed: '%s'", err))
		return
	}

	var blacklist database.Blacklist

	if err := blacklist.New(img, link); err != nil {
		utils.SendDirectMessage(s, m, fmt.Sprintf("Ban failed: '%s'", err))
		return
	}

	blacklist.Save()

	go utils.RemoveMessage(s, m)
	go utils.SendMessageSuccess(s, m, "Ban successful")
}
