package ai

import (
	"errors"
	"math/rand"
	"strconv"
)

var NUM_FORWARD = 3
var NUM_CONT_FORWARD = 3
var LEFT = 5
var RIGHT = 5
var AROUND = 3
var MIN = 1

var path = "pkg/static/"

func GetCommandVoice(command string) (string, error) {
	switch command {
	case "FORWARD":
		num := strconv.FormatInt(int64(rand.Intn(NUM_FORWARD-MIN)+MIN), 10)
		return path + "Forward" + num, nil
	case "CONTINUE FORWARD":
		num := strconv.FormatInt(int64(rand.Intn(NUM_FORWARD-MIN)+MIN), 10)
		return path + "ContinueForward" + num, nil
	case "TURN LEFT":
		num := strconv.FormatInt(int64(rand.Intn(NUM_FORWARD-MIN)+MIN), 10)
		return path + "Left" + num, nil
	case "TURN RIGHT":
		num := strconv.FormatInt(int64(rand.Intn(NUM_FORWARD-MIN)+MIN), 10)
		return path + "Right" + num, nil
	case "TURN AROUND":
		num := strconv.FormatInt(int64(rand.Intn(NUM_FORWARD-MIN)+MIN), 10)
		return path + "TurnAround" + num, nil
	default:
		return "", errors.New("command not supported")
	}
}
