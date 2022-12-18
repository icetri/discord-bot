package service

import (
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Command      *Command
	Channel      *discordgo.Channel
	Discord      *discordgo.Session
	VoiceChannel *discordgo.Channel
	Guild        *discordgo.Guild
	User         *discordgo.User
	Name         string
	Args         []string
}

func NewContext(
	discord *discordgo.Session,
	guild *discordgo.Guild,
	channel *discordgo.Channel,
	user *discordgo.User,
	command *Command,
	name string,
	args []string,
) *Context {
	return &Context{
		Command: command,
		Channel: channel,
		Discord: discord,
		Guild:   guild,
		User:    user,
		Name:    name,
		Args:    args,
	}
}

func (ctx Context) GetVoiceChannel() *discordgo.Channel {
	if ctx.VoiceChannel != nil {
		return ctx.VoiceChannel
	}

	for _, state := range ctx.Guild.VoiceStates {
		if state.UserID != ctx.User.ID {
			continue
		}

		// TODO err logger?
		ctx.VoiceChannel, _ = ctx.Discord.State.Channel(state.ChannelID)

		return ctx.VoiceChannel
	}

	return nil
}
