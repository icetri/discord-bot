package service

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Context struct {
	logger   logrus.FieldLogger
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
	logger logrus.FieldLogger,
) *Context {
	return &Context{
		Commands: commands,
		Channel:  channel,
		Discord:  discord,
		Guild:    guild,
		User:     user,
		Name:     name,
		Args:     args,
		logger:   logger,
	}
}

func (ctx Context) GetVoiceChannel() *discordgo.Channel {
	for _, state := range ctx.Guild.VoiceStates {
		if state.UserID != ctx.User.ID {
			continue
		}

		voiceChannel, err := ctx.Discord.State.Channel(state.ChannelID)
		if err != nil {
			ctx.logger.Errorf("get channel: %v", err)
			return nil
		}

		return voiceChannel
	}

	return nil
}

func (ctx Context) Reply(content string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.Channel.ID, content)
	if err != nil {
		ctx.logger.Errorf("channel message send: %v", err)
		return nil
	}

	return msg
}

func (ctx Context) JoinVoiceChannel(guildID, voiceChannelID string) *discordgo.VoiceConnection {
	voiceConnection, err := ctx.Discord.ChannelVoiceJoin(guildID, voiceChannelID, false, true)
	if err != nil {
		ctx.logger.Errorf("channel voice join: %v", err)
		return nil
	}

	return voiceConnection
}
