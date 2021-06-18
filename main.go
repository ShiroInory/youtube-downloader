package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	var url string
	var format string

	fmt.Printf("Input URL:")
	fmt.Scanln(&url)
	cmd := exec.Command("youtube-dl", "-F", url)
	showMeg(cmd)

	fmt.Printf("\nInput format:")
	fmt.Scanln(&format)
	cmd = exec.Command("youtube-dl", "-f", format, "--write-thumbnail", url)
	showMeg(cmd)

	pause()
}

func showMeg(cmd *exec.Cmd) {
	var stdoutBuf, stderrBuf bytes.Buffer
	var errStdout, errStderr error

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	Stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	Stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		_, errStdout = io.Copy(Stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(Stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
}

func pause() {
	fmt.Printf("按任意鍵退出...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
