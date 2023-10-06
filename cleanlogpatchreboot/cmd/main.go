package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var falsePositives = []string{"SELinux", "timed out", "Free disk space"}

func main() {
	loadLog()
}

func loadLog() {
	f, err := os.Open("input.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// if line begins with "ok" then print it
		if strings.HasPrefix(line, "ok") {
			//exctract only what is between [ and ]
			s := strings.Split(line, "[")
			s = strings.Split(s[1], "]")
			fmt.Println(s[0])
		}
		// if line begins with "msg" then print it
		if strings.Contains(line, "\"msg\"") {
			// only print what is after "msg"
			s := strings.Split(line, "\"msg\"")
			fp := false
			for _, word := range falsePositives {
				if strings.Contains(s[1], word) {
					fp = true
					continue
				}
			}
			if fp {
				fmt.Printf("\033[31m%s\033[0m\n", s[1])
			} else {
				fmt.Printf("\033[32m%s\033[0m\n", s[1])
			}
		}
	}
}
