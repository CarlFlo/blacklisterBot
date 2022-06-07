package bot

import (
	"strings"

	"github.com/CarlFlo/blacklisterBot/src/config"
	"github.com/CarlFlo/blacklisterBot/src/utils"
	"github.com/CarlFlo/malm"
	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore messages from bots
	if config.CONFIG.IgnoreBotMessages && m.Author.Bot {
		return
	}

	// Is the message a command and is the user authorized to use the bot?
	if strings.HasPrefix(m.Content, config.CONFIG.BotPrefix) && utils.IsAuthorized(m.Author.ID) {
		handleCommand(s, m)
		return
	}

	// Check the message for blacklisted content
	checkAttachments(s, m)
}

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

}

func checkAttachments(s *discordgo.Session, m *discordgo.MessageCreate) {

	for _, att := range m.Message.Attachments {

		switch att.ContentType {
		case "image/png", "image/jpeg":
			img, err := handleImage(&att.URL)
			if err != nil {
				malm.Error("%s", err)
			}

			if banned := checkImage(img); banned {
				malm.Info("Blacklisted image posted by %s", m.Author.Username)
				removeMessage(s, m)
			}

		default:
			malm.Debug("Unknown content type: %s", att.ContentType)
		}

	}
}

func removeMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
		malm.Error("Could not delete the message: %s", err)
	}
}
