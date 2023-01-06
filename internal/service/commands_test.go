package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommandStruct_GetHelp(t *testing.T) {
	type fields struct {
		command Command
		help    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Success return help",
			fields: fields{
				help: "play",
			},
			want: "play",
		},
		{
			name: "Success return empty help if command not exist",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &CommandStruct{
				command: tt.fields.command,
				help:    tt.fields.help,
			}

			help := cs.GetHelp()
			require.Equalf(t, tt.want, help, "GetHelp = %v, want %v", help, tt.want)
		})
	}
}

func TestCommands_GetCommand(t *testing.T) {
	var defaultCommand Command

	defaultCmdMap := CmdMap{
		"play": CommandStruct{
			command: defaultCommand,
			help:    "play",
		},
	}

	emptyCommand := defaultCommand
	emptyCommand = nil

	emptyCmdMap := make(CmdMap)

	cmdName := "play"
	emptyCmdName := ""

	tests := []struct {
		name  string
		arg   string
		cmds  CmdMap
		want  *Command
		found bool
	}{
		{
			name:  "Success get command",
			arg:   cmdName,
			cmds:  defaultCmdMap,
			want:  &defaultCommand,
			found: true,
		},
		{
			name:  "Success if command not exist",
			arg:   emptyCmdName,
			cmds:  defaultCmdMap,
			want:  &emptyCommand,
			found: false,
		},
		{
			name:  "Success if empty map",
			arg:   emptyCmdName,
			cmds:  emptyCmdMap,
			want:  &emptyCommand,
			found: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Commands{
				cmds: tt.cmds,
			}

			command, found := c.GetCommand(tt.arg)
			require.Equalf(t, tt.found, found, "found = %v, want %v", found, tt.found)
			require.Equalf(t, tt.want, command, "GetCommand = %v, want %v", command, tt.want)
		})
	}
}

func TestCommands_GetCommands(t *testing.T) {
	var defaultCommand Command

	defaultCmds := CmdMap{
		"play": CommandStruct{
			command: defaultCommand,
			help:    "play",
		},
		"stop": CommandStruct{
			command: defaultCommand,
			help:    "stop",
		},
	}

	emptyCmds := make(CmdMap)

	tests := []struct {
		name string
		cmds CmdMap
		want CmdMap
	}{
		{
			name: "Success get commands",
			cmds: defaultCmds,
			want: defaultCmds,
		},
		{
			name: "Success get empty commands",
			cmds: emptyCmds,
			want: emptyCmds,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Commands{
				cmds: tt.cmds,
			}

			commands := c.GetCommands()
			require.Len(t, commands, len(tt.want))
			for k, v := range commands {
				w, ok := tt.want[k]
				require.True(t, ok)
				require.Equal(t, v, w)
			}
		})
	}
}

func TestCommands_Register(t *testing.T) {
	var defaultCommand Command

	defaultCmd := make(CmdMap)

	oneRegisterCommand := CmdMap{
		"play": CommandStruct{
			command: defaultCommand,
			help:    "play",
		},
	}

	twoRegisterCommand := CmdMap{
		"play": CommandStruct{
			command: defaultCommand,
			help:    "play",
		},
		"stop": CommandStruct{
			command: defaultCommand,
			help:    "stop",
		},
	}

	type args struct {
		name        string
		command     Command
		helpMessage string
	}

	oneRegisterCommandArgs := args{
		name:        "play",
		command:     defaultCommand,
		helpMessage: "play",
	}

	twoRegisterCommandArgs := args{
		name:        "stop",
		command:     defaultCommand,
		helpMessage: "stop",
	}

	tests := []struct {
		name string
		cmds CmdMap
		args args
		want CmdMap
	}{
		{
			name: "Success register command",
			cmds: defaultCmd,
			args: oneRegisterCommandArgs,
			want: oneRegisterCommand,
		},
		{
			name: "Success add new register command",
			cmds: oneRegisterCommand,
			args: twoRegisterCommandArgs,
			want: twoRegisterCommand,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Commands{
				cmds: tt.cmds,
			}

			c.Register(tt.args.name, tt.args.command, tt.args.helpMessage)
			require.Equal(t, tt.want, c.cmds)
		})
	}
}
