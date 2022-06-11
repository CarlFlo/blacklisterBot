package commands

import (
	"github.com/CarlFlo/blacklisterBot/src/config"
	"github.com/CarlFlo/blacklisterBot/src/utils"
	"github.com/CarlFlo/malm"
	"github.com/bwmarrin/discordgo"
)

func Reload(s *discordgo.Session, m *discordgo.MessageCreate) {

	var msg *discordgo.Message
	var err error

	utils.RemoveMessageAfter(s, m.ChannelID, m.ID)

	if err := config.ReloadConfig(); err != nil {
		malm.Error("Could not reload the config: %s", err)
		msg, err = utils.SendMessageFailure(s, m, "Failed to reload config. Check the console")

		if err == nil {
			utils.RemoveMessageAfter(s, msg.ChannelID, msg.ID)
		}
		return
	}

	msg, err = utils.SendMessageSuccess(s, m, "Config reloaded")
	if err == nil {
		utils.RemoveMessageAfter(s, msg.ChannelID, msg.ID)
	}

}
