package service

func Play(ctx Context) {
	if _, err := ctx.Discord.ChannelMessageSend(ctx.Channel.ID, "play :)"); err != nil {
		return
	}
}
