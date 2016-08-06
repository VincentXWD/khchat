package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net"
	"os"
	"strings"
	"time"
)

func Encode(raw []byte) []byte {
	var encoded bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &encoded)
	encoder.Write(raw)
	encoder.Close()
	return encoded.Bytes()
}

func Alert() {

}

func ConnectServer(server string) {
	var (
		name     string
		trimName string
	)
	conn, err := net.DialTimeout("tcp", server, time.Second*30)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	defer conn.Close()
	log.Infof("Connect %s successfully.", server)
	cin := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your name first : ")
	name, _ = cin.ReadString('\n')
	trimName = strings.Trim(name, "\r\n")
	conn.Write([]byte(trimName + " has connected.\n"))
	for {
		fmt.Print("You: ")
		input, err := cin.ReadString('\n')
		if err != nil {
			log.Errorln(err.Error())
		}
		trimInput := strings.Trim(input, "\r\n")
		_, err = conn.Write(Encode([]byte(trimName + ": " + trimInput)))
		if err != nil {
			log.Errorln(err.Error())
		}
	}
}

func main() {
	ConnectServer(os.Args[1])
}
