package utils

import "github.com/CarlFlo/blacklisterBot/src/config"

func IsAuthorized(discordID string) bool {

	for _, id := range config.CONFIG.TrustedUsersIDs {
		if id == discordID {
			return true
		}
	}
	return false
}
