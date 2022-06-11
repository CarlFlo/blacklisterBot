package commands

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"image"
	"image/jpeg"
	"net/url"

	"github.com/CarlFlo/blacklisterBot/src/database"
	"github.com/CarlFlo/blacklisterBot/src/utils"
	"github.com/CarlFlo/malm"
	"github.com/bwmarrin/discordgo"
)

func Unban(s *discordgo.Session, m *discordgo.MessageCreate, args *[]string) {

	for _, att := range m.Attachments {
		*args = append(*args, att.URL)
	}

	if len(*args) == 1 {
		// No URL provided. Ban the last URL found in the n previous messages message.
		utils.SendMessageFailure(s, m, "No image URL provided")
		return
	}

	for _, link := range (*args)[1:] {
		_, err := url.Parse(link)
		if err != nil {
			utils.SendMessageFailure(s, m, "Invalid URL")
			return
		}

		b, err := utils.FetchFromURL(&link)
		if err != nil {
			utils.SendMessageFailure(s, m, "Internal error")
			malm.Error("%s", err)
			return
		}

		img, _, err := image.Decode(bytes.NewReader(b))
		if err != nil {
			utils.SendMessageFailure(s, m, "Internal error")
			malm.Error("%s", err)
			return
		}
		buf := new(bytes.Buffer)
		if err = jpeg.Encode(buf, img, nil); err != nil {
			utils.SendMessageFailure(s, m, "Internal error")
			malm.Error("%s", err)
			return
		}

		hasher := sha1.New()
		hasher.Write(buf.Bytes())
		SHA1 := fmt.Sprintf("%x", hasher.Sum(nil))

		malm.Info("Removing SHA1: %s", SHA1)

		if err := database.DB.Delete(&database.Blacklist{}, "sha1 = ?", SHA1).Error; err != nil {
			malm.Error("Could not delete: '%s'", err)
		}
	}

	utils.RemoveMessage(s, m)
}
