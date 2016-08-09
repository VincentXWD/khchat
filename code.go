package main

import (
	"bytes"
	"encoding/base64"
	log "github.com/Sirupsen/logrus"
)

func Encode(raw []byte) []byte {
	var encoded bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &encoded)
	encoder.Write(raw)
	encoder.Close()
	return encoded.Bytes()
}

func Decode(raw []byte) []byte {
	var buf bytes.Buffer
	decoded := make([]byte, 215)
	buf.Write(raw)
	decoder := base64.NewDecoder(base64.StdEncoding, &buf)
	decoder.Read(decoded)
	return decoded
}

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