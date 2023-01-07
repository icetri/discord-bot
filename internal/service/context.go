package service

import (
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Commands CmdMap
	Channel  *discordgo.Channel
	Discord  *discordgo.Session
	Guild    *discordgo.Guild
	User     *discordgo.User
	Name     string
	Args     []string
}

func NewContext(
	discord *discordgo.Session,
	guild *discordgo.Guild,
	channel *discordgo.Channel,
	user *discordgo.User,
	commands CmdMap,
	name string,
	args []string,
) *Context {
	return &Context{
		Commands: commands,
		Channel:  channel,
		Discord:  discord,
		Guild:    guild,
		User:     user,
		Name:     name,
		Args:     args,
	}
}

func (ctx Context) GetVoiceChannel() *discordgo.Channel {
	for _, state := range ctx.Guild.VoiceStates {
		if state.UserID != ctx.User.ID {
			continue
		}

		// TODO err logger?
		vc, err := ctx.Discord.State.Channel(state.ChannelID)
		if err != nil {
			return nil
		}

		return vc
	}

	return nil
}

// TODO err logger
func (ctx Context) Reply(content string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.Channel.ID, content)
	if err != nil {
		return nil
	}

	return msg
}

// TODO err logger
func (ctx Context) JoinVoiceChannel(guildID, voiceChannelID string) (*discordgo.VoiceConnection, error) {
	voiceSession, err := ctx.Discord.ChannelVoiceJoin(guildID, voiceChannelID, false, true)
	if err != nil {
		return nil, err
	}

	return voiceSession, nil
}
