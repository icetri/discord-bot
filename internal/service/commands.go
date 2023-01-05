package service

type (
	Commands struct {
		cmds CmdMap
	}

	CmdMap map[string]CommandStruct

	CommandStruct struct {
		command Command
		help    string
	}

	Command func(ctx Context)
)

func NewCommands() *Commands {
	return &Commands{
		cmds: make(CmdMap),
	}
}

func (c Commands) GetCommands() CmdMap {
	return c.cmds
}

// TODO pointer?
func (c Commands) GetCommand(name string) (*Command, bool) {
	cmd, found := c.cmds[name]
	return &cmd.command, found
}

func (c Commands) Register(name string, command Command, helpMessage string) {
	cmd := CommandStruct{command: command, help: helpMessage}
	c.cmds[name] = cmd
}

func (cs *CommandStruct) GetHelp() string {
	return cs.help
}
