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

const CurrentVersion = "2022-06-11"

func init() {

	utils.Clear()
	malm.SetLogVerboseBitmask(35) // Debug, Info & Warning

	config.Load()
	database.Load()
	utils.CheckVersion(CurrentVersion)
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

// https://discordapp.com/oauth2/authorize?&client_id=984944576294944768&scope=bot&permissions=142342
