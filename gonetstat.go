package gonetstat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

type nodeconn struct {
	line        string
	proto       string
	local       string
	localport   int
	foreign     string
	foreignport int
	state       string
}

func Getnetstat() {

	var line string

	ips := hostip()

	fmt.Println(ips)

	cmd := exec.Command("netstat", "-np", "tcp")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		line = scanner.Text()
		processline(line)
		//fmt.Println(line) // Println will add back the final '\n'
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}

func processline(l string) (n *nodeconn) {

	protos := map[string]bool{
		"tcp4": true,
		"tcp6": true,
	}

	lparts := strings.Split(l, " ")

	if protos[lparts[0]] {

		n = new(nodeconn)

		c := 0

		for _, v := range lparts {

			if v != "" {

				switch c {

				case 0:
					// proto
					n.proto = v
				case 1, 2:
					// nothing to do
				case 3:
					// local
					n.local = v
				case 4:
					// foreign
					n.foreign = v
				case 5:
					n.state = v
				}
				c = c + 1
			}
		}
		fmt.Println(n)
	}

	return
}

func hostip() (addrs []string) {

	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	addrs, err = net.LookupHost(name)

	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}

	return

}
