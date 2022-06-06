package bot

import (
	"strings"

	"github.com/CarlFlo/blacklisterBot/src/config"
	"github.com/CarlFlo/blacklisterBot/src/utils"
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
	interceptMessage(s, m)
}

func handleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

}

func interceptMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

}
