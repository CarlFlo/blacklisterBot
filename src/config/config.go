package config

import (
	"encoding/json"
	"io/ioutil"
)

var CONFIG *configStruct

type configStruct struct {
	Token             string   `json:"token"`
	BotPrefix         string   `json:"botPrefix"`
	TrustedUsersIDs   []string `json:"trustedUsersIDs"`
	IgnoreBotMessages bool     `json:"ignoreBotMessages"`
	BotInfo           botInfo  `json:"botInfo"`
	Database          database `json:"database"`
}

type botInfo struct {
	AppID      string `json:"appID"`
	Permission uint64 `json:"permission"`
	VersionURL string `json:"versionURL"`
}

type database struct {
	FileName string `json:"filename"`
}

// ReloadConfig is a wrapper function for reloading the config. For clarity
func ReloadConfig() error {
	return readConfig()
}

// readConfig will read the config file
func readConfig() error {

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return err
	}

	return json.Unmarshal(file, &CONFIG)
}

func createConfig() error {

	// Default config settings
	configStruct := configStruct{
		Token:             "",
		BotPrefix:         ",",
		TrustedUsersIDs:   []string{},
		IgnoreBotMessages: true,
		BotInfo: botInfo{
			AppID:      "",
			Permission: 207878, // https://discordapi.com/permissions.html#207878
			VersionURL: "https://raw.githubusercontent.com/CarlFlo/blacklisterBot/main/main.go",
		},
		Database: database{
			FileName: "blacklist.db",
		},
	}

	jsonData, _ := json.MarshalIndent(configStruct, "", "   ")
	err := ioutil.WriteFile("config.json", jsonData, 0644)

	return err
}

func LoadConfiguration() error {

	if err := readConfig(); err != nil {
		if err = createConfig(); err != nil {
			return err
		}
		if err = readConfig(); err != nil {
			return err
		}
	}
	return nil
}
