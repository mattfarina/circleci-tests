package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("/tmp/workspace", 0744)
	if err != nil {
		log.Fatal(err)
	}
	message := []byte("Hello Helm")
	err = ioutil.WriteFile("/tmp/workspace/test", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
