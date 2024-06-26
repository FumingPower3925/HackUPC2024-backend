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
var MAX = 3
var MIN = 1

var path = "pkg/static/"

func GetCommandVoice(command string) (string, error) {
	switch command {
	case "FORWARD":
		num := strconv.FormatInt(int64(rand.Intn(NUM_FORWARD-MIN)+MIN), 10)
		return path + "Forward" + num + ".wav", nil
	case "CONTINUE FORWARD":
		num := strconv.FormatInt(int64(rand.Intn(NUM_FORWARD-MIN)+MIN), 10)
		return path + "ContinueForward" + num + ".wav", nil
	case "TURN LEFT":
		num := strconv.FormatInt(int64(rand.Intn(NUM_FORWARD-MIN)+MIN), 10)
		return path + "Left" + num + ".wav", nil
	case "TURN RIGHT":
		num := strconv.FormatInt(int64(rand.Intn(NUM_FORWARD-MIN)+MIN), 10)
		return path + "Right" + num + ".wav", nil
	case "TURN AROUND":
		num := strconv.FormatInt(int64(rand.Intn(NUM_FORWARD-MIN)+MIN), 10)
		return path + "TurnAround" + num + ".wav", nil
	case "ARRIVED":
		return path + "Arrived" + ".wav", nil
	default:
		return "", errors.New("command not supported")
	}
}

func GetAudio(response string) (string, error) {
	switch response {
	case "REST":
		num := strconv.FormatInt(int64(rand.Intn(MAX-MIN)+MIN), 10)
		return path + "Rest" + num + ".wav", nil
	case "BATHROOM":
		num := strconv.FormatInt(int64(rand.Intn(MAX-MIN)+MIN), 10)
		return path + "Bathroom" + num + ".wav", nil
	case "FLIGHT":
		num := strconv.FormatInt(int64(rand.Intn(MAX-MIN)+MIN), 10)
		return path + "Flight" + num + ".wav", nil
	default:
		return "", errors.New("command not supported")
	}
}
