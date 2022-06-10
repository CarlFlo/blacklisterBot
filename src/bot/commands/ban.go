package commands

import (
	"bytes"
	"fmt"
	"image"
	"net/url"
	"strings"

	"github.com/CarlFlo/blacklisterBot/src/database"
	"github.com/CarlFlo/blacklisterBot/src/utils"
	"github.com/CarlFlo/malm"
	"github.com/bwmarrin/discordgo"
)

func Ban(s *discordgo.Session, m *discordgo.MessageCreate, args *[]string) {

	for _, att := range m.Attachments {
		*args = append(*args, att.URL)
	}

	if len(*args) == 1 {
		// No URL provided. Ban the last URL found in the n previous messages message.
		utils.SendMessageFailure(s, m, "No URL provided")
		return
	}

	// itterate over all the potential URLs provided
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

	if err := blacklist.Save(); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			go utils.SendMessageNeutral(s, m, "Already banned")
		} else {
			utils.SendDirectMessage(s, m, fmt.Sprintf("Unhandled DB error: '%s'", err))
			malm.Error("Unhandled DB error: '%s'", err)
			return
		}
	} else {
		go utils.SendMessageSuccess(s, m, "Ban successful")
	}

	utils.RemoveMessage(s, m)
}
