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
	if config.CONFIG.Settings.IgnoreBotMessages && m.Author.Bot {
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
		check(s, m, &url)
	}

	for _, att := range m.Message.Attachments {
		switch att.ContentType {
		case "image/png", "image/jpeg":
			check(s, m, &att.URL)

		default:
			malm.Debug("Unknown content type: %s", att.ContentType)
		}
	}
}

func check(s *discordgo.Session, m *discordgo.MessageCreate, link *string) {
	img, err := handleImage(link)
	if err != nil {
		malm.Error("%s", err)
	}

	if banned, method := checkImage(img); banned {
		if config.CONFIG.Settings.LogRemovalInConsole {
			malm.Info("Blacklisted image posted by %s [Method: %s]", m.Author.Username, method)
		}
		utils.RemoveMessage(s, m)
	}
}

func findURLInMessage(m *discordgo.MessageCreate) []string {

	// regex - Detects an url containing an image
	r := regexp.MustCompile(`(http(s?):)([/|.|\w|\s|-])*\.(?:jpg|jpeg|png)`)
	matches := r.FindAllString(m.Message.Content, -1)
	return matches
}
