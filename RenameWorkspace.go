package main

import (
	"fmt"
	"os"
	"strings"

	"go.i3wm.org/i3/v4"

	"i3-scripts-go/utils"
)

func replaceSpacesWithUnderscore(s string) string {
	return strings.ReplaceAll(strings.TrimSpace(s), " ", "_")
}

func renamei3Ws(wsIndex int64, newName string) {
	var cmd string

	if newName == "" {
		cmd = fmt.Sprintf("rename workspace to %d", wsIndex)
	} else {
		cmd = fmt.Sprintf("rename workspace to %d:%s", wsIndex, replaceSpacesWithUnderscore(newName))
	}

	i3.RunCommand(cmd)
}

func main() {
	focusedWS, err := utils.GetFocusedWorkspace()
	if err != nil {
		utils.NotifySend(2, fmt.Sprintf("RenameWorkspace: %s", err.Error()))
		os.Exit(1)
	}

	userPromptResult := utils.Runi3Input("Append title: ", 0)

	renamei3Ws(focusedWS.Num, userPromptResult)
}
