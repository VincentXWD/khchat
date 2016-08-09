package main

import (
	"bufio"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/fatih/color"
	"net"
	"os"
	"strings"
	"time"
)

func Alert() {

}

func ClientHandle(server string) {
	var (
		name     string
		trimName string
		buf      []byte
	)
	conn, err := net.DialTimeout("tcp", server, time.Second*30)
	defer conn.Close()
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	log.Infof("Connect %s successfully.", server)
	cin := bufio.NewReader(os.Stdin)
	fmt.Print("Please input your name first : ")
	name, _ = cin.ReadString('\n')
	trimName = strings.Trim(name, "\r\n")
	_, err = conn.Write([]byte(trimName + " has connected.\n"))
	if err != nil {
		log.Errorln(err.Error())
	}
	// recieve
	go func() {
		for {
			buf = make([]byte, 512)
			_, err = conn.Read(buf)
			flag := CheckError(err)
			if flag == 0 {
				log.Fatalln("Server has disconnected.")
			}
			color.Green(string(Decode(buf)))
		}
	}()
	for {
		// send
		func() {
			input, err := cin.ReadString('\n')
			if err != nil {
				log.Errorln(err.Error())
			}
			trimInput := strings.Trim(input, "\r\n")
			_, err = conn.Write(Encode([]byte(trimName + ": " + trimInput)))
			if err != nil {
				log.Errorln(err.Error())
			}
		}()
	}

}

func main() {
	ClientHandle(os.Args[1])
}
