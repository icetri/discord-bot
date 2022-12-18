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
	cmds := &Commands{
		cmds: make(CmdMap),
	}

	cmds.registerCommands()

	return cmds
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

func (c Commands) registerCommands() {
	//	c.Register("help", HelpCommand, "TODO")
	c.Register("play", Play, "TODO")
	//	c.Register("stop", StopCommand, "TODO")
}
