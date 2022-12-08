package chatcommander

import (
	commands2 "RainbowRunner/internal/game/chatcommander/commands"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/objects"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

var commands = map[string]func(player *objects.RRPlayer, args []string){
	"exec": commands2.ExecuteLua,
}

var commandSplitRegex = regexp.MustCompile(`(?:^@|)(?:(".*"|\S+)(?: |$))+?`)

type ChatCommander struct {
}

func (c ChatCommander) Execute(player *objects.RRPlayer, msg string) {
	if !strings.HasPrefix(msg, "@") {
		log.Errorf("chat command does not start with a @ and is not valid: %s", msg)
		SendErrorMessageResponse(player, fmt.Sprintf("chat command does not start with a @ and is not valid: %s", msg))
		return
	}

	splitCmd := commandSplitRegex.FindAll([]byte(msg), -1)

	if len(splitCmd) == 0 {
		log.Errorf("unable to parse chat command: %s", msg)
		SendErrorMessageResponse(player, fmt.Sprintf("unable to parse chat command: %s", msg))
		return
	}

	commandName := strings.Trim(string(splitCmd[0][1:]), " ")

	cmd, ok := commands[commandName]

	if !ok {
		SendErrorMessageResponse(player, fmt.Sprintf("could not find command: %s", commandName))
		log.Errorf("could not find command: %s", commandName)
		return
	}

	args := make([]string, 0, len(splitCmd)-1)

	if len(splitCmd) > 1 {
		for i := 1; i < len(splitCmd); i++ {
			argString := strings.Trim(string(splitCmd[i]), ` "`)

			args = append(args, argString)
		}
	}

	cmd(player, args)
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
