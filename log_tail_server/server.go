package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"bufio"

	"github.com/gorilla/websocket"
)

const (
	filePath string = "logFile.txt"
	one_kb   int    = int(1000)
	one_mb   int    = 1000 * one_kb
)

func getLatest10Lines() ([]string, error) {
	// open file
	logFile, err := os.Open(filePath)
	if err != nil {
		return []string{}, err
	}
	defer logFile.Close()

	fileInfo, err := logFile.Stat()
	if err != nil {
		return []string{}, err
	}
	fileSize := fileInfo.Size()
	buffData := make([]byte, one_mb)
	if fileSize > int64(one_mb) {
		// get last one mb chunk
		n, err := logFile.ReadAt(buffData, int64(fileSize-int64(one_mb)))
		if n == 0 {
			fmt.Println("no of bytes read is zero")
		}
		if err == io.EOF {
			fmt.Println("reached end of file")
		}
		if err != nil {
			return []string{}, err
		}
	} else {
		buffReder := bufio.NewReader(logFile)
		n, err := buffReder.Read(buffData)
		if n == 0 {
			fmt.Println("no of bytes read is zero")
		}
		if err == io.EOF {
			fmt.Println("reached end of file")
		}
		if err != nil {
			return []string{}, err
		}
	}

	// fmt.Println("=========== read chunk ================")
	// log.Println(string(buffData))

	var dataList []string
	data := string(buffData)
	dataList = strings.Split(data, "\n")
	if len(dataList) > 10 {
		dataList = dataList[len(dataList)-10:]
	}

	for _, lineData := range dataList {
		fmt.Println("line : ", lineData)
	}

	return dataList, nil
}

var webSocketConnections map[*websocket.Conn]bool

type SocketHandler struct {
}

// now here upgrade http to socket connection
func (sh SocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	newConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error while upgrading http to websocket", err)
		return
	}
	if _, ok := webSocketConnections[newConn]; ok {
		fmt.Println("already connected")
	} else {
		fmt.Println("socket connection made and saved")
		webSocketConnections[newConn] = true

		// TODO : share latest 10 lines
		var tenLines []string

		tenLines, err = getLatest10Lines()
		if err != nil {
			fmt.Println("err : ", err)
		}

		for _, data := range tenLines {
			err = newConn.WriteMessage(websocket.TextMessage, []byte(data))
			if err != nil {
				fmt.Println("erro while sending intial data", err)
			}
		}
	}
}

func main() {

	// getLatest10Lines()
	// 1. how many ways to create http enpoints
	var socketHandler SocketHandler
	http.Handle("/socket/logs", socketHandler)

	webSocketConnections = make(map[*websocket.Conn]bool)

	fmt.Println("listening at 8080")
	err := http.ListenAndServe(":8080", socketHandler)
	if err != nil {
		fmt.Println("err : ", err)
	}
}
