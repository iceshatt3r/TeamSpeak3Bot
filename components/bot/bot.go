package bot

import (
	"log"
	"time"

	"github.com/Overflow3D/ts3Bot_v2/components/dispatcher"
	"github.com/Overflow3D/ts3Bot_v2/components/query"
)

// {
//   "ServerID": "4",
//   "HeadAdminCliDB": "1",
//   "Spacer" : {
//     "test": "595",
//     "S1": "19",
//     "S2": "20",
//     "S3": "21",
//     "S4": "28"
//   },
//   "ChannelAdmin": "18",
//   "BotMainChannel": "595",
//   "GuestRoom": "644",
//   "TempGroup": "62",
//   "PermGroup": "63",
//   "PunishRoom": "595",
//   "OmittedRang": ["Head Admin", "Administrator", "Vouched"]
// }

const executor, listener = "SkyNet", "SkyNetEyes"

type TeamSpeakBots struct {
	Bots  map[string]*Bot
	Await chan struct{}
}

//Bot ,stores all information about Bot struct
type Bot struct {
	onlineSince     int64
	query           *query.TelNet
	shutdownProcess *ShutDown
}

type Config struct {
	Address  string
	ServerID string
	Login    string
	Password string
	BotNames []string
}

//ShutDown ,contains of information about shutting down bot
type ShutDown struct {
	stopSchelduler chan struct{}
	isForced       bool
}

func SetConfig(addr, login, password, id string, names []string) *Config {
	return &Config{Address: addr, Login: login, Password: password, ServerID: id, BotNames: names}
}

func New(config *Config) (*TeamSpeakBots, error) {
	bots := &TeamSpeakBots{Bots: make(map[string]*Bot, 2)}
	for indexName := 0; indexName < len(config.BotNames); indexName++ {
		newBot, err := bots.setUpBot(config, indexName)
		if err != nil {
			return nil, err
		}
		bots.Bots[config.BotNames[indexName]] = newBot
	}
	bots.Await = make(chan struct{})
	return bots, nil
}

func (t *TeamSpeakBots) setUpBot(config *Config, indexName int) (*Bot, error) {
	bot := new(Bot)
	var err error

	bot.query, err = query.NewServerQuery(config.Address, isListener(config.BotNames[indexName]))
	if err != nil {
		return nil, err
	}
	var commands []*query.Command
	commands = startParameters(config.ServerID, config.BotNames[indexName], config.Login, config.Password)
	if isListener(config.BotNames[indexName]) {
		go t.notifyRegister(bot)
		commands = append(commands, query.RegisterListener()...)
	}
	bot.startTime()
	bot.query.ExecMultiple(commands, false)

	return bot, nil
}

func (t *TeamSpeakBots) notifyRegister(b *Bot) {
	for {
		notifications := <-b.query.Notify
		notifyEvent := query.FormatResponse(notifications, "notify")
		dispatcher.Dispatch(notifyEvent)
	}
}

func (b *Bot) startTime() {
	b.onlineSince = time.Now().Unix()
}

func (b *Bot) scheduler() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			log.Println("tick tack")
		case <-b.shutdownProcess.stopSchelduler:
			ticker.Stop()
			return
		}
	}

}

func startParameters(sid, name, login, pass string) []*query.Command {
	return []*query.Command{query.UseServer(sid), query.LogIn(login, pass), query.Nickname(name)}
}

func isListener(bot string) bool {
	return bot == listener
}
