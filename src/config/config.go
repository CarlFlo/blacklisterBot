package config

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/CarlFlo/malm"
)

var CONFIG *configStruct

type configStruct struct {
	Token           string     `json:"token"`
	BotPrefix       string     `json:"botPrefix"`
	TrustedUsersIDs []string   `json:"trustedUsersIDs"`
	Thresholds      thresholds `json:"thresholds"`
	BotInfo         botInfo    `json:"botInfo"`
	Database        database   `json:"database"`
	Settings        settings   `json:"settings"`
}

type botInfo struct {
	AppID      string `json:"appID"`
	Permission uint64 `json:"permission"`
	VersionURL string `json:"versionURL"`
	DepositURL string `json:"depositURL"`
}

type database struct {
	FileName string `json:"filename"`
}

type thresholds struct {
	Average    int `json:"average"`
	Difference int `json:"difference"`
	Perception int `json:"perception"`
}

type settings struct {
	LogRemovalInConsole   bool          `json:"logRemovalInConsole"`
	IgnoreBotMessages     bool          `json:"ignoreBotMessages"`
	RemoveBotMessageAfter time.Duration `json:"removeBotMessageAfter"`
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
		Token:           "",
		BotPrefix:       ",",
		TrustedUsersIDs: []string{},
		Thresholds:      thresholds{Average: 10, Difference: 10, Perception: 10},
		BotInfo: botInfo{
			AppID:      "",
			Permission: 142342, // https://discordapi.com/permissions.html#142342
			VersionURL: "https://raw.githubusercontent.com/CarlFlo/blacklisterBot/main/main.go",
			DepositURL: "https://github.com/CarlFlo/blacklisterBot",
		},
		Database: database{
			FileName: "blacklist.db",
		},
		Settings: settings{
			LogRemovalInConsole:   true,
			IgnoreBotMessages:     true,
			RemoveBotMessageAfter: 3,
		},
	}

	jsonData, _ := json.MarshalIndent(configStruct, "", "   ")
	err := ioutil.WriteFile("config.json", jsonData, 0644)

	return err
}

func loadConfiguration() error {

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

// Loads the configuration
// Any problems will be logged
func Load() {

	if err := loadConfiguration(); err != nil {
		malm.Fatal("Error loading configuration: %v", err)
	}

}
