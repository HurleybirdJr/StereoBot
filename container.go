package main

import (
	"stereobot/internal/bot"

	"github.com/bwmarrin/discordgo"
)

type Container struct {
	Discord *discordgo.Session
	Bot     *bot.Bot
}
