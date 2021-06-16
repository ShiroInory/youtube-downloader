package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	var url string
	var format string
	var out bytes.Buffer

	fmt.Printf("Input URL:")
	fmt.Scanln(&url)
	cmd := exec.Command("youtube-dl", "-F", url)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
	fmt.Printf("Input format:")
	fmt.Scanln(&format)
	cmd = exec.Command("youtube-dl", "-f", format, "--write-thumbnail", url)
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}
