package gonetstat

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Getnetstat() {

	cmd := exec.Command("netstat", "-p", "tcp")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
