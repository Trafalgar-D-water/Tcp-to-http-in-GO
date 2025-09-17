package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println("Current working dir:", dir)
	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	str := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}
		str = ""
		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != 0 {
			str += string(data[:i])
			data = data[i+1:]
			fmt.Printf("read: %s\n", str)
			str = ""
		}
	}

}

// elts checl
