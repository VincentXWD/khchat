package main

import (
	"bytes"
	"encoding/base64"
)

func encode(raw []byte) []byte {
	var encoded bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &encoded)
	encoder.Write(raw)
	encoder.Close()
	return encoded.Bytes()
}

func decode(raw []byte) []byte {
	var decoded bytes.Buffer
	decoder := base64.NewDecoder(base64.StdEncoding, &decoded)
	decoder.Read(raw)
	return decoded.Bytes()
}