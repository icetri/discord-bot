package commands

import (
	"bytes"
	"fmt"

	"github.com/discord-bot/internal/platform/discord"
	"github.com/discord-bot/internal/service"
)

func Help(ctx service.Context) {
	buffer := bytes.NewBufferString("Commands supported: \n")

	for cmdName, cmdStruct := range ctx.Commands {
		buffer.WriteString(fmt.Sprintf("\t %s%s - %s\n", discord.PREFIX, cmdName, cmdStruct.GetHelp()))
	}

	ctx.Reply(buffer.String())
}
