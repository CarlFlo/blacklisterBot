package bot

import "github.com/bwmarrin/discordgo"

func joinedGuildHandler(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	//event.Guild.ID
	//Add the guildID to the database

}
