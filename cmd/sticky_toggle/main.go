package main

import (
	"log"

	stickyToggle "github.com/modernpacifist/i3-scripts-go/internal/i3operations/sticky_toggle"
)

func main() {
	if err := stickyToggle.Execute(); err != nil {
		log.Fatal(err)
	}
}
