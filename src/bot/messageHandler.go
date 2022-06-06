package bot

import (
	"strings"

	"github.com/CarlFlo/blacklisterBot/src/config"
	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore messages from bots
	if config.CONFIG.IgnoreBotMessages && m.Author.Bot {
		return
	}

	// Is the message a command and is the user authorized to use the bot?
	if !strings.HasPrefix(m.Content, config.CONFIG.BotPrefix) && !isAuthorized(m.Author.ID) {
		return
	}

}

func isAuthorized(discordID string) bool {

	for _, id := range config.CONFIG.TrustedUsersIDs {
		if id == discordID {
			return true
		}
	}
	return false
}
