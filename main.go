package main

import (
	"os"

	"github.com/Overflow3D/ts3Bot_v2/components/bot"
)

var (
	addr     = os.Getenv("TELNET_ADDRESS")
	login    = os.Getenv("TELNET_LOGIN")
	password = os.Getenv("TELNET_PASSWORD")
	serverID = os.Getenv("TELNET_SERVER_ID")
)

func main() {
	config := bot.SetConfig(addr, login, password, serverID, []string{"SkyNet", "SkyNetEyes"})
	bot.New(config)
}
