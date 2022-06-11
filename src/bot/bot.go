package bot

import (
	"github.com/CarlFlo/blacklisterBot/src/config"
	"github.com/CarlFlo/malm"
	"github.com/bwmarrin/discordgo"
)

// Discord session
var Session *discordgo.Session

func StartBot() {

	variableCheck()

	// Creates the session
	var err error
	Session, err = discordgo.New("Bot " + config.CONFIG.Token)
	if err != nil {
		malm.Fatal("Error creating Discord session: %s", err)
	}

	// Adds message handler (https://github.com/bwmarrin/discordgo/blob/37088aefec2241139e59b9b804f193b539be25d6/eventhandlers.go#L937)
	Session.AddHandler(messageHandler)
	Session.AddHandler(readyHandler)

	// Attempts to open connection
	err = Session.Open()
	if err != nil {
		malm.Fatal("%s", err)
	}

}

func StopBot() {
	Session.Close()
}

func variableCheck() {

	// This function checks if some important variables are set in the config file
	problem := false

	if len(config.CONFIG.Token) == 0 {
		malm.Error("No bot Token provided in the config file!")
		problem = true
	}

	if len(config.CONFIG.BotInfo.AppID) == 0 {
		malm.Error("No AppID provided in the config file! (The bot's Discord ID)")
		problem = true
	}

	if len(config.CONFIG.TrustedUsersIDs) == 0 {
		malm.Error("No TrustedUsersIDs provided in the config file! (This should be your Discord ID along with other people that are allowed to use the bots commands)")
		problem = true
	}

	if problem {
		malm.Fatal("There are at least one variable missing in the configuration file. Please fix the above errors!")
	}

}
