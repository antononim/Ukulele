package session

import (
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	s    time.Duration = time.Second
	prod bool          = true
)

var (
	IsActive bool = false
)

func HandleSession() {
	// Entry point

	var (
		TS  time.Time = getTSFromBD()
		now time.Time
	)

	for {
		now = time.Now()

		if (TS.Sub(now) > 0) && (!IsActive) {
			// TS > now => got more time to play
			activate()
		} else if (TS.Sub(now) <= 0) && (IsActive) {
			// TS < now => game over
			deactivate()
		}
	}
}

func activate() {
	/*
		Activate session/PC
	*/

	logrus.Debug("Activating...")
	IsActive = true
	if !prod {
		return
	}

	out, err := exec.Command("explorer.exe").CombinedOutput()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"output": out,
			"Error":  err,
		}).Error("Unable to open explorer")
	}

}

func deactivate() {
	/*
		Deactivate session/PC
	*/

	logrus.Debug("Deactivating...")
	IsActive = false
	if !prod {
		return
	}

	out, err := exec.Command("C:\\Windows\\System32\\taskkill", "/f", "/im", "explorer.exe").CombinedOutput()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"output": out,
			"Error":  err,
		}).Error("Unable to kill explorer")
	}

}

func getTSFromBD() time.Time {
	/*
		Temporary function, should be replaced with backend request
	*/

	return time.Now().Add(2 * s)
}
