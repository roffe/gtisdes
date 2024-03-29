package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"os"
)

const (
	//Key = "aced00057372001e636f6d2e73756e2e63727970746f2e70726f76696465722e4445534b65796b349c35da1568980200015b00036b65797400025b427870757200025b42acf317f8060854e0020000787000000008cd0d02d951ec8cef"
	Key      = "cd0d02d951ec8cef"
	HelpText = "Usage: gtisde <encrypt|decrypt> <filename>"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal(HelpText)
	}
	filename := os.Args[2]

	data, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	key, err := hex.DecodeString(Key)
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "encrypt":
		if bytes.HasPrefix(data, []byte{10, 15, 15, 14}) {
			log.Println("File is already encrypted")
			return
		}
		encrypt(data, key, filename)
	case "decrypt":
		if !bytes.HasPrefix(data, []byte{10, 15, 15, 14}) {
			log.Println("File is not encrypted")
			return
		}
		data = data[4:]
		decrypt(data, key, filename)
	default:
		log.Fatal(HelpText)
	}

}

func readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func decrypt(data, key []byte, filename string) {
	b, err := DesDecrypt(data, key)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(filename, b, 0644); err != nil {
		log.Fatal(err)
	}
}

func encrypt(data, key []byte, filename string) {
	enc, err := DesEncrypt(data, key)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(filename, append([]byte{10, 15, 15, 14}, enc...), 0644); err != nil {
		log.Fatal(err)
	}
}
