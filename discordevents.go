package main

import (
	"github.com/jonas747/discordgo"
	"log"
)

func DiscordReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Println("Received ready from discord!")
}

func DiscordMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println("Failed getting channel from state", err)
		return
	}

	if channel.GuildID != config.Guild {
		log.Println("guild not tracked ignoring...")
		return
	}

	diff := channel.Position - config.ReorderStartPos
	if diff <= config.ReorderDeadZone {
		return
	}

	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		log.Println("Failed getting guild from state", err)
		return
	}

	log.Println("Sending channel", channel.Name, "To top")
	oldPos := channel.Position
	channel.Position = config.ReorderStartPos + 1

	for _, c := range guild.Channels {
		if c.Position <= config.ReorderStartPos || c.ID == channel.ID || c.Position > oldPos {
			continue // Not one of the managed channels, out of the range of the new order chage or already reordered
		}
		c.Position++
	}
	err = s.GuildChannelsEdit(channel.GuildID, guild.Channels)
	if err != nil {
		log.Println("Failed reordering", err)
		return
	}
}
