package utils

import (
	"time"

	"github.com/CarlFlo/blacklisterBot/src/config"
	"github.com/CarlFlo/malm"
	"github.com/bwmarrin/discordgo"
)

func SendDirectMessage(s *discordgo.Session, m *discordgo.MessageCreate, content string) (*discordgo.Message, error) {
	ch, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		return nil, err
	}
	return s.ChannelMessageSend(ch.ID, content)
}

func SendMessageSuccess(s *discordgo.Session, m *discordgo.MessageCreate, content string) (*discordgo.Message, error) {
	return sendMessageEmbed(s, m, content, 1673044)
}

func SendMessageFailure(s *discordgo.Session, m *discordgo.MessageCreate, content string) (*discordgo.Message, error) {
	return sendMessageEmbed(s, m, content, 15282218)
}

func SendMessageNeutral(s *discordgo.Session, m *discordgo.MessageCreate, content string) (*discordgo.Message, error) {
	return sendMessageEmbed(s, m, content, 28368)
}

func sendMessageEmbed(s *discordgo.Session, m *discordgo.MessageCreate, content string, color int) (*discordgo.Message, error) {
	return s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: content,
		Color:       color,
	})
}

func RemoveMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
		malm.Error("Could not delete the message: %s", err)
	}
}

// Removes a message after n seconds.
// If RemoveBotMessageAfter in the config is -1 then the message wont be deleted
func RemoveMessageAfter(s *discordgo.Session, channelID, messageID string) {

	if config.CONFIG.Settings.RemoveBotMessageAfter < 0 {
		return
	}

	time.Sleep(config.CONFIG.Settings.RemoveBotMessageAfter * time.Second)

	if err := s.ChannelMessageDelete(channelID, messageID); err != nil {
		malm.Error("Could not delete the message: %s", err)
	}
}
