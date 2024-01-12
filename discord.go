package main

import "github.com/bwmarrin/discordgo"

// StreamChannelMessages uses the `discordgo.Session` to fetch 100 messages from their API for the given channel.
// `guildId` specifies the guild (server), which the channel belongs to. `channelId` is the channel's ID.
// `beforeId` allows specifying the beginning Message ID  to use. If the parameter is provided, the API will return messages, which are older than the given `beforeId`.
func StreamChannelMessages(s *discordgo.Session, guildId string, channelId string, beforeId string) ([]*NormalisedDiscordMessage, error) {
	m, err := s.ChannelMessages(channelId, 100, beforeId, "", "")
	if err != nil {
		return nil, err
	}

	var result = make([]*NormalisedDiscordMessage, 0, len(m))
	for _, msg := range m {
		result = append(result, &NormalisedDiscordMessage{
			GuildId:   guildId,
			ChannelId: channelId,
			MessageId: msg.Reference().MessageID,
			Username:  msg.Author.Username,
			Content:   msg.Content,
			Timestamp: msg.Timestamp,
		})
	}
	return result, nil
}
