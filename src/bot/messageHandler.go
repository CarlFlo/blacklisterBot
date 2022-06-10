package bot

import (
	"regexp"
	"strings"

	"github.com/CarlFlo/blacklisterBot/src/bot/commands"
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

	input := strings.TrimPrefix(m.Message.Content, config.CONFIG.BotPrefix)

	input = strings.ToLower(input)

	args := strings.Split(input, " ")

	switch args[0] {
	case "ban":
		commands.Ban(s, m, &args)
	case "unban":
		commands.Unban(&args)
	default:
		malm.Debug("Unknown command: %s", args[0])
	}

}

func checkAttachments(s *discordgo.Session, m *discordgo.MessageCreate) {

	for _, url := range findURLInMessage(m) {
		img, err := handleImage(&url)
		if err != nil {
			malm.Error("%s", err)
		}

		if banned := checkImage(img); banned {
			malm.Info("Blacklisted image posted by %s", m.Author.Username)
			utils.RemoveMessage(s, m)
		}
	}

	for _, att := range m.Message.Attachments {
		switch att.ContentType {
		case "image/png", "image/jpeg":
			img, err := handleImage(&att.URL)
			if err != nil {
				malm.Error("%s", err)
			}

			if banned := checkImage(img); banned {
				malm.Info("Blacklisted image posted by %s", m.Author.Username)
				utils.RemoveMessage(s, m)
			}

		default:
			malm.Debug("Unknown content type: %s", att.ContentType)
		}
	}
}

func findURLInMessage(m *discordgo.MessageCreate) []string {

	// regex
	r := regexp.MustCompile(`(http(s?):)([/|.|\w|\s|-])*\.(?:jpg|jpeg|png)`)
	matches := r.FindAllString(m.Message.Content, -1)
	return matches
}
