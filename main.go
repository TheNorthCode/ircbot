package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mpgelliston/ircbot/actions"
	"github.com/mpgelliston/ircbot/bot"
)

var portFlag = flag.Int("P", 6667, "The port on which to connect the IRC server")

var verboseFlag = flag.Bool("v", false, "Verbose logging")
var debugFlag = flag.Bool("d", false, "Debug logging, if this is set the -v flag is ignored")

var nickFlag = flag.String("n", "NorthBot", "The NICK which the bot will assume")
var serverFlag = flag.String("s", "irc.libera.chat", "The IRC server URL")
var passFlag = flag.String("p", "", "The user password of the bot")
var chanFlag = flag.String("c", "", "The channel to connect to on the server")

func main() {
	flag.Parse()

	fmt.Println("Configuring IRC Bot...")
	options := bot.BotOptions{
		Nick:     *nickFlag,
		User:     "NorthBot",
		Name:     "NorthBot",
		Password: *passFlag,
		Server:   *serverFlag,
		Port:     *portFlag,
		Channel:  *chanFlag,
		Admins:   map[string]bool{"matt1982": true, "lux0r": true, "mrbalihai": true},
		Debug:    *debugFlag,
		Verbose:  *verboseFlag,
	}

	fmt.Println("Starting IRC Bot")

	b, err := bot.NewBot(options)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Attaching Actions")
	b.AddAction(bot.BotAction{
		Name:   "Join",
		Action: actions.JoinAction,
	})
	b.AddAction(bot.BotAction{
		Name:   "Welcome",
		Action: actions.WelcomeAction,
	})
	b.AddAction(bot.BotAction{
		Name:   "Sexy",
		Action: actions.SexyAction,
	})
	b.AddAction(bot.BotAction{
		Name:   "List Actions",
		Action: actions.ListActionsAction,
	})
	b.AddAction(bot.BotAction{
		Name:   "Uptime",
		Action: actions.UptimeAction,
	})
	b.AddAction(bot.BotAction{
		Name:   "Crypto",
		Action: actions.CryptoPriceAction,
	})

	// Create the client
	b.Connect()
}
