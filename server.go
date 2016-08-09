package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net"
	"os"
	"strings"
)

var connect map[net.Conn]net.Conn

func BroadCast(self net.Conn, buf []byte) {
	for conn := range connect {
		if conn == self {
			continue
		}
		conn.Write(buf)
	}
}

func ServerHandle(conn net.Conn) {
	var (
		name      string
		info, buf []byte
	)
	info = make([]byte, 512)
	_, err := conn.Read(info)
	name = strings.Split(string(info), " ")[0]
	log.Warnf("Client from %s. He(She) is [%s].", conn.RemoteAddr().String(), name)
	CheckError(err)
	for {
		buf = make([]byte, 512)
		_, err := conn.Read(buf)
		BroadCast(conn, buf)
		flag := CheckError(err)
		if flag == 0 {
			log.Warnf("%s has disconnected.", name)
			BroadCast(nil, []byte("%s has disconnected."))
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
		connect[conn] = conn
		CheckError(err)
		go ServerHandle(conn)
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
	connect = make(map[net.Conn]net.Conn)
	Server(GetLocalIp(), os.Args[1])
}
