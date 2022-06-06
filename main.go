package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CarlFlo/blacklisterBot/src/bot"
	"github.com/CarlFlo/blacklisterBot/src/config"
	"github.com/CarlFlo/blacklisterBot/src/database"
	"github.com/CarlFlo/blacklisterBot/src/utils"
	"github.com/CarlFlo/malm"
)

const CurrentVersion = "2022-06-06"

func init() {

	utils.Clear()

	malm.SetLogVerboseBitmask(0)

	if err := config.LoadConfiguration(); err != nil {
		malm.Fatal("Error loading configuration: %v", err)
	}

	if err := database.SetupDatabase(); err != nil {
		malm.Fatal("Database initialization error: %s", err)
	}

	// Handles checking if there is an update available for the bot
	upToDate, githubVersion, err := utils.BotVersonHandler(CurrentVersion)
	if err != nil {
		malm.Error("%s", err)
	}

	if upToDate {
		malm.Debug("Version %s", CurrentVersion)
	} else {
		malm.Info("New version available! New version: '%s'; Your version: '%s'", githubVersion, CurrentVersion)
	}

}

func main() {

	bot.StartBot()

	time.Sleep(500 * time.Millisecond) // Added this sleep so the messages below will come last
	// Keeps bot from closing. Waits for CTRL-C
	malm.Info("Press CTRL-C to initiate shutdown")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	malm.Info("Shutting down!")

	// Run cleanup code here
	close(sc)
	bot.StopBot()
}
