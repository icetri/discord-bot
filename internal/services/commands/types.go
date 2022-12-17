package commands

type (
	Commands struct {
		cmds CmdMap
	}

	CmdMap map[string]CommandStruct

	CommandStruct struct {
		command Command
		help    string
	}

	Command func(interface{})
)
