package main

import "time"

// NormalisedDiscordMessage represents the Discord's message structure in a format
// agnostic to how discord manages messages.
// This is the structure that is used to manage the messages outside of Discord's fetch API.
type NormalisedDiscordMessage struct {
	// Guild ID used to identify which server the message belongs to.
	GuildId string

	// Channel ID used to identify which channel the message belongs to.
	ChannelId string

	// Message ID of the actual Discord message in Discord's API.
	MessageId string

	// Account username of the message's author. This is not the user's nickname in the server.
	Username string

	// The actual string content of the message.
	Content string

	// Timestamp of the creation of the message.
	Timestamp time.Time
}
