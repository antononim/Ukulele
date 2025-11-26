package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	addr *string = flag.String("addr", "localhost", "http address")
	port *string = flag.String("port", "8000", "http port")
)

func main() {
	http.HandleFunc("/ws", newConnHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This part is not ready yet, please use ws://ip/ws on this server.")
	})
	var host string = *addr + ":" + *port
	logrus.Info("http server starting on ", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		logrus.WithError(err).Fatal("ListenAndServe")
	}
}

func newConnHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	go handleConnection(conn)
}

func handleConnection(conn *websocket.Conn) {
	defer conn.Close()
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		// just print out for now
		logrus.WithFields(logrus.Fields{
			"type": messageType,
			"msg":  string(msg),
		}).Info("Message received")

		// send echo resp
		if err := conn.WriteMessage(messageType, msg); err != nil {
			logrus.WithError(err).Error("Error at writing a message")
			continue
		}
	}
}
