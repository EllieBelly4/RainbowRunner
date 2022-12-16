package main

import (
	"RainbowRunner/internal/objects/actions"
	"os"
	"path/filepath"
	"strings"
	template2 "text/template"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	for i := 0; i <= 0xFF; i++ {
		behav := actions.BehaviourAction(i)

		GenerateFor(behav, cwd)
	}
}

func GenerateFor(behav actions.BehaviourAction, cwd string) {
	actionName := strings.Split(behav.String(), "BehaviourAction")[1]

	filePath := filepath.Join(cwd, "action_"+strings.ToLower(actionName)+".go")

	if _, err := os.Stat(filePath); err == nil {
		return
	}

	file, err := os.Create(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	template := template2.New("ActionTemplate")

	template, err = template.Parse(ActionTemplate)

	if err != nil {
		panic(err)
	}

	err = template.Execute(file, struct {
		ActionName string
	}{
		ActionName: actionName,
	})
}
