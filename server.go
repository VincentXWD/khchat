package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net"
	"os"
	"strings"
)

func CheckError(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			return 0
		}
		log.Fatalln(err.Error())
		return -1
	}
	return 1
}

func Decode(raw []byte) []byte {
	var buf bytes.Buffer
	decoded := make([]byte, 215)
	buf.Write(raw)
	decoder := base64.NewDecoder(base64.StdEncoding, &buf)
	decoder.Read(decoded)
	return decoded
}

func GetMessage(conn net.Conn, from string) {
	var (
		name string
	)
	info := make([]byte, 512)
	_, err := conn.Read(info)
	name = strings.Split(string(info), " ")[0]
	log.Warnf("Client from %s. He(She) is [%s].", conn.RemoteAddr().String(), name)
	CheckError(err)
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		flag := CheckError(err)
		if flag == 0 {
			log.Warnf("%s has disconnected.", name)
			return
		}
		fmt.Println(string(Decode(buf)))
	}
}

func Server(ip string, port string) {
	log.Infof("Server address: %s. Establishing.", ip)
	listener, err := net.Listen("tcp", ip+":"+port)
	defer listener.Close()
	CheckError(err)
	log.Infoln("Server established.")
	for {
		conn, err := listener.Accept()
		defer conn.Close()
		CheckError(err)
		go GetMessage(conn, conn.RemoteAddr().String())
	}
}

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return strings.Split(addrs[1].String(), "/")[0]
}
func main() {
	Server(GetLocalIp(), os.Args[1])
}
