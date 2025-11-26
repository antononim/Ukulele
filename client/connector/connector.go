package connector

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	addr *string = flag.String("addr", "localhost", "http address")
	port *string = flag.String("port", "8000", "http port")
)

func Connect() {
	flag.Parse()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Client app starting up")

	host := *addr + ":" + *port
	u := url.URL{Scheme: "ws", Host: host, Path: "/ws"}
	conn, resp := newConn(u.String(), nil)
	logrus.Debug("conn resp: ", resp)
	defer conn.Close()

	go readServerMessages(conn)

	// go writeServerMessages(conn)

	for {
		// var messageArray []string
		var msg string
		_, err := fmt.Scan(&msg)
		if err != nil {
			logrus.WithError(err).Error("Unable to read console input")
		}
		logrus.WithField("msg", msg).Debug("Message Entered")

		if msg == "q" {
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logrus.WithError(err).Fatal("Write Close")
			}
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			logrus.WithError(err).Fatal("WriteMessage")
		}
	}

}

func newConn(urlStr string, rHead http.Header) (*websocket.Conn, *http.Response) {
	conn, resp, err := websocket.DefaultDialer.Dial(urlStr, rHead)
	if err != nil {
		logrus.WithError(err).Fatal("Dial")
	}
	return conn, resp
}

func readServerMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.WithError(err).Fatal("Read Unexpected Close")
			} else {
				logrus.WithError(err).Error("Read")
			}
			break
		}

		logrus.WithField("msg", string(message)).Debug("Message recieved")
	}
}

// func writeServerMessages(conn *websocket.Conn) {
// 	for {
// 		var msg string
// 		fmt.Scanln(&msg)

// 		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))d
// 		if err != nil {
// 			log.Println("write:", err)
// 			return
// 		}
// 	}
// }
