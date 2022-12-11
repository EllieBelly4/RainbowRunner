package commands

import (
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/lua"
	"RainbowRunner/internal/objects"
	"fmt"
	log "github.com/sirupsen/logrus"
	lua2 "github.com/yuin/gopher-lua"
	"strings"
)

func ExecuteLua(player *objects.RRPlayer, args []string) {
	if len(args) < 1 {
		SendLuaErrorMessageResponse(player, "You must provide a script name with function name. Ex: @exec general.changeZone town")
		return
	}

	splitScriptName := strings.Split(args[0], ".")
	scriptName := "scripts." + strings.Join(splitScriptName[:len(splitScriptName)-1], ".")
	functionName := splitScriptName[len(splitScriptName)-1]

	log.Infoln("Execute Lua " + args[0])

	script := lua.GetScript(scriptName)

	if script == nil {
		SendLuaErrorMessageResponse(player, fmt.Sprintf("could not find script with name: %s", scriptName))
		log.Infof("could not find script with name: %s", scriptName)
		return
	}

	state := lua2.NewState()
	defer state.Close()

	objects.RegisterLuaGlobals(state)
	interceptPrint(player, state)

	err := script.Execute(state)

	if err != nil {
		SendLuaErrorMessageResponse(player, fmt.Sprintf("failed to execute script %s: %s", err, err.Error()))
		log.Infof("failed to execute script %s: %s", err, err.Error())
		return
	}

	if functionName == "_" {
		return
	}

	funkyRaw := state.GetGlobal(functionName)

	if _, ok := funkyRaw.(*lua2.LNilType); ok {
		SendLuaErrorMessageResponse(player, fmt.Sprintf("could not find function %s in script %s ", functionName, scriptName))
		log.Infof("could not find function %s in script %s ", functionName, scriptName)
		return
	}

	funky, ok := funkyRaw.(*lua2.LFunction)

	if !ok {
		SendLuaErrorMessageResponse(player, fmt.Sprintf("requests function %s was not a function in script %s ", functionName, scriptName))
		log.Infof("requests function %s was not a function in script %s ", functionName, scriptName)
		return
	}

	state.Push(funky)
	state.Push(player.CurrentCharacter.ToLua(state))

	argLen := len(args) - 1

	for i := 1; i < len(args); i++ {
		state.Push(lua2.LString(args[i]))
	}

	state.Call(argLen+1, 0)
	return
}

func interceptPrint(player *objects.RRPlayer, state *lua2.LState) {
	state.SetGlobal("print", state.NewFunction(func(state *lua2.LState) int {
		luaCustomPrint(player, state)
		return 0
	}))
}

func SendLuaPrintMessageResponse(player *objects.RRPlayer, msg string) {
	response := messages.ChatMessage{
		Channel: messages.MessageChannelSourceNoob,
		Message: "[Exec] " + msg,
		Sender:  "Lady Lua",
	}

	player.Conn.SendMessage(response)
}

func luaCustomPrint(player *objects.RRPlayer, L *lua2.LState) int {
	top := L.GetTop()

	msg := ""

	for i := 1; i <= top; i++ {
		msg += L.ToStringMeta(L.Get(i)).String()
		fmt.Print(msg)

		if i != top {
			msg += " "
			fmt.Print("\t")
		}
	}

	SendLuaPrintMessageResponse(player, msg)
	fmt.Println("")
	return 0
}

func SendLuaErrorMessageResponse(player *objects.RRPlayer, s string) {
	response := messages.ChatMessage{
		Channel: messages.MessageChannelSourceNoob,
		Message: "[ERROR] " + s,
		Sender:  "Lady Lua",
	}

	player.Conn.SendMessage(response)
}
