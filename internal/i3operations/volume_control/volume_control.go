package volume_control

import (
	"fmt"
	"log"
	"math"
	"os/exec"
	"strconv"
	"strings"

	common "github.com/modernpacifist/i3-scripts-go/internal/i3operations"
)

func getCurrentVolume() float64 {
	var out []byte
	var err error

	cmd := `amixer -D pulse sget Master | grep 'Left:' | awk -F'[][]' '{ print $2 }' | tr -d '%'`
	if out, err = exec.Command("bash", "-c", cmd).Output(); err != nil {
		log.Fatal(err)
	}

	var num float64
	if num, err = strconv.ParseFloat(strings.TrimSuffix(string(out), "\n"), 64); err != nil {
		log.Fatal(err)
	}

	return num
}

func ToggleVolume() {
	cmd := `pactl set-sink-mute @DEFAULT_SINK@ toggle`
	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		log.Fatal(err)
	}

	common.NotifySend(1.5, "VolumeControl: toggled")
}

func RoundVolume() {
	currentVolume := getCurrentVolume()
	roundedVolume := math.Round(currentVolume/5) * 5

	cmd := fmt.Sprintf("pactl set-sink-volume @DEFAULT_SINK@ %.f%%", roundedVolume)
	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		log.Fatal(err)
	}

	common.NotifySend(1.5, fmt.Sprintf("VolumeControl: rounded to %.f%%", roundedVolume))
}

func AdjustVolume(changeValue string, maxVolume float64) {
	currentVolume := getCurrentVolume()
	change, err := strconv.ParseFloat(changeValue, 64)
	if err != nil {
		log.Fatal(err)
	}

	newVolume := currentVolume + change
	if newVolume > maxVolume {
		common.NotifySend(1.5, fmt.Sprintf("VolumeControl: cannot adjust volume above %.f%%", maxVolume))
		return
	}

	cmd := fmt.Sprintf("pactl set-sink-volume @DEFAULT_SINK@ %s%%", changeValue)
	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		log.Fatalf("%s: pactl is not installed on this system", err)
	}

	common.NotifySend(1.5, fmt.Sprintf("VolumeControl: %s%%", changeValue))
}
