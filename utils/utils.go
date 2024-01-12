package utils

import (
	"fmt"
	"log"
	"os/exec"
	"errors"

	"go.i3wm.org/i3/v4"
)

func NotifySend(seconds float32, msg string) {
	_, err := exec.Command("bash", "-c", fmt.Sprintf("notify-send --expire-time=%.f \"%s\"", seconds*1000, msg)).Output()
	// TODO: probably should catch this error via defer <04-12-23, modernpacifist> //
	if err != nil {
		log.Println(err)
	}
}

func Runi3Input() string {
	output, err := exec.Command("bash", "-c", "i3-input -P \"Append title: \" | grep -oP \"output = \\K.*\" | tr -d '\n'").Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(output)
}

func GetI3Tree() i3.Tree {
	tree, err := i3.GetTree()
	if err != nil {
		log.Fatal(err)
	}
	return tree
}

func GetWorkspaceNodes() *i3.Node {
	i3Tree := GetI3Tree()

	node := i3Tree.Root.FindFocused(func(n *i3.Node) bool {
		return n.Focused == true
	})

	if node == nil {
		log.Fatal(errors.New("Could not get focused node"))
	}

	return node
}

func GetFocusedNode() *i3.Node {
	i3Tree := GetI3Tree()

	node := i3Tree.Root.FindFocused(func(n *i3.Node) bool {
		return n.Focused == true
	})

	if node == nil {
		log.Fatal(errors.New("Could not get focused node"))
	}

	return node
}

func GetFocusedWorkspace() (i3.Workspace, error) {
	o, err := i3.GetWorkspaces()
	if err != nil {
		return i3.Workspace{}, err
	}

	for _, ws := range o {
		if ws.Focused == true {
			return ws, nil
		}
	}

	return i3.Workspace{}, errors.New("Could not get focused workspace")
}
