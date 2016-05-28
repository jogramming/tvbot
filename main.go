package main

import (
	"flag"
	"fmt"
	"github.com/jonas747/discordgo"
	"os"
)

var (
	flagConfigPath = flag.String("c", "config.json", "The path to the config file")
	config         *Config
	session        *discordgo.Session
)

func checkErr(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	var err error
	config, err = LoadConfig(*flagConfigPath)
	checkErr("Failed loading config", err)

	session, err = discordgo.New(config.Token)
	checkErr("Failed creating discordgo session", err)

	addHandlers()

	err = session.Open()
	checkErr("Failed opening gateway connection", err)
	select {}
}

func addHandlers() {
	session.AddHandler(DiscordReady)
	session.AddHandler(DiscordMessageCreate)
}
