package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken string
	guildId  string
	fp       string
)

func main() {
	flag.StringVar(&botToken, "t", "", "Token for the Discord bot")
	flag.StringVar(&guildId, "g", "", "Guild ID of the Discord server")
	flag.StringVar(&fp, "o", ".", "Filepath to where to save the CSV files")
	flag.Parse()

	if botToken == "" || guildId == "" {
		flag.Usage()
		os.Exit(1)
	}

	bot, err := discordgo.New("Bot " + botToken)
	if err != nil {
		panic(err)
	}

	err = bot.Open()
	if err != nil {
		panic(err)
	}
	defer bot.Close()

	channels, err := bot.GuildChannels(guildId)
	if err != nil {
		panic(err)
	}

	for _, channel := range channels {
		// For some reason, categories are also considered channels.
		// And so are voice chats, etc.
		// Skip everything that's not an actual guild text chat.
		if channel.Type != discordgo.ChannelTypeGuildText {
			continue
		}

		var (
			lastId string = ""
			count  int    = 0
		)

		csvHandler, err := NewCSVHandler(fmt.Sprintf("%s/%s.csv", fp, channel.ID), []string{"timestamp", "guild_id", "channel_id", "username", "content"})
		if err != nil {
			panic(err)
		}

		for {
			messages, err := StreamChannelMessages(bot, guildId, channel.ID, lastId)
			if err != nil {
				fmt.Println(err)
				break
			}
			if len(messages) == 0 {
				fmt.Printf("\033[2K\r ✅ DONE: %s count=%d\n", channel.Name, count)
				break
			}

			csvHandler.SaveMessagesToFile(messages)

			lastId = messages[len(messages)-1].MessageId
			count += len(messages)

			fmt.Printf("\033[2K\r ⟲ Processing: %s count=%d", channel.Name, count)
		}

		csvHandler.Close()
	}
	fmt.Println("done")
}
