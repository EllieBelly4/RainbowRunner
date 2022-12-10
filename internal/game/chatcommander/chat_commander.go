package chatcommander

import (
	commands2 "RainbowRunner/internal/game/chatcommander/commands"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/objects"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var commands = map[string]commands2.ChatCommandHandler{
	"exec": commands2.ExecuteLua,
	"z": commands2.AliasCustom(commands2.ExecuteLua, func(player *objects.RRPlayer, args []string) []string {
		return []string{"general.changeZone", args[0]}
	}),
}

var commandSplitRegex = regexp.MustCompile(`(?:^@|)(?:(".*"|\S+)(?: |$))+?`)

type ChatCommander struct {
}

func (c ChatCommander) Execute(player *objects.RRPlayer, msg string) error {
	if !strings.HasPrefix(msg, "@") {
		SendErrorMessageResponse(player, fmt.Sprintf("chat command does not start with a @ and is not valid: %s", msg))
		return errors.New(fmt.Sprintf("chat command does not start with a @ and is not valid: %s", msg))
	}

	splitCmd := commandSplitRegex.FindAll([]byte(msg), -1)

	if len(splitCmd) == 0 {
		SendErrorMessageResponse(player, fmt.Sprintf("unable to parse chat command: %s", msg))
		return errors.New(fmt.Sprintf("unable to parse chat command: %s", msg))
	}

	commandName := strings.Trim(string(splitCmd[0][1:]), " ")

	cmd, ok := commands[commandName]

	if !ok {
		SendErrorMessageResponse(player, fmt.Sprintf("could not find command: %s", commandName))
		return errors.New(fmt.Sprintf("could not find command: %s", commandName))
	}

	args := make([]string, 0, len(splitCmd)-1)

	if len(splitCmd) > 1 {
		for i := 1; i < len(splitCmd); i++ {
			argString := strings.Trim(string(splitCmd[i]), ` "`)

			args = append(args, argString)
		}
	}

	cmd(player, args)
	return nil
}

func SendErrorMessageResponse(player *objects.RRPlayer, s string) {
	response := messages.ChatMessage{
		Channel: messages.MessageChannelSourceNoob,
		Message: "[ERROR] " + s,
		Sender:  "The Commander",
	}

	player.Conn.SendMessage(response)
}

func NewChatCommander() *ChatCommander {
	return &ChatCommander{}
}
