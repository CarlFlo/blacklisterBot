package bot

import (
	"fmt"

	"github.com/CarlFlo/malm"
	"github.com/bwmarrin/discordgo"
)

func readyHandler(s *discordgo.Session, ready *discordgo.Ready) {

	serverOrServers := "server"
	if len(s.State.Guilds) > 1 {
		serverOrServers += "s"
	}

	malm.Info("Bot is connected and present on %d %s", len(s.State.Guilds), serverOrServers)

	statusMessage := fmt.Sprintf("on %d %s", len(s.State.Guilds), serverOrServers)

	// Shows up like the bot is streaming. Allows us to have a link.
	s.UpdateStreamingStatus(0, statusMessage, "https://www.youtube.com/watch?v=yjdG80Rs8Zo")
	// Normal message
	//s.UpdateGameStatus(0, statusMessage)
}
