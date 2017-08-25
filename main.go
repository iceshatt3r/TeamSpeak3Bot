package main

import (
	"fmt"
	"os"

	"github.com/Overflow3D/ts3Bot_v2/components/bot"
	_ "github.com/joho/godotenv/autoload"
)

var (
	addr     = os.Getenv("TELNET_ADDRESS")
	login    = os.Getenv("TELNET_LOGIN")
	password = os.Getenv("TELNET_PASSWORD")
	serverID = os.Getenv("TELNET_SERVER_ID")
)

func main() {
	config := bot.SetConfig(addr, login, password, serverID, []string{"SkyNetEyes", "SkyNet"})
	fmt.Println("Starting bot")
	bots, err := bot.New(config)
	if err != nil {
		panic(err)
	}
	<-bots.Await
}
