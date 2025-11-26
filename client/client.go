package main

import (
	"client/session"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	fmt.Println("Hello world!")
	session.HandleSession()
}
