package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	cmdUNK  = ""
	cmdFWD  = "forward"
	cmdDOWN = "down"
	cmdUP   = "up"
)

var (
	commands = []string{cmdFWD, cmdDOWN, cmdUP}
)

type subcmd struct {
	dir    string
	change int
}

func main() {
	input, err := ioutil.ReadFile("testdata/real.input")
	if err != nil {
		log.Fatalf("failed to open real.input: %s", err)
	}

	cmds, err := parse(input)
	if err != nil {
		log.Fatalf("failed to parse real.input: %s", err)
	}

	aim, x, y := simulate(0, 0, 0, cmds)
	log.Printf("simulated to (aim:%d, %d,%d)", aim, x, y)
	log.Printf("result = %d", x*y)
}

func simulate(aim, x, y int, cmds []subcmd) (int, int, int) {
	for i := range cmds {
		switch cmds[i].dir {
		case cmdFWD:
			x += cmds[i].change
			y += cmds[i].change * aim
		case cmdDOWN:
			aim += cmds[i].change
		case cmdUP:
			aim -= cmds[i].change
		}
	}

	return aim, x, y
}

func parse(buf []byte) ([]subcmd, error) {
	var (
		split = bytes.Split(buf, []byte{'\n'})
		ret   []subcmd
	)

	for i := range split {
		if len(split[i]) == 0 {
			continue
		}

		cmd, err := parsecmd(string(split[i]))
		if err != nil {
			return nil, err
		}

		ret = append(ret, cmd)
	}

	return ret, nil
}

func parsecmd(cmd string) (command subcmd, err error) {
	cmdline := strings.Split(cmd, " ")
	if len(cmdline) != 2 {
		return subcmd{}, fmt.Errorf("invalid command: %s", cmd)
	}

	command.dir = cmdUNK
	for i := range commands {
		if cmdline[0] != commands[i] {
			continue
		}

		command.dir = commands[i]
		break
	}

	n, err := strconv.Atoi(cmdline[1])
	if err != nil {
		return subcmd{}, fmt.Errorf("bad change: %s", cmdline[1])
	}

	command.change = n

	return
}
