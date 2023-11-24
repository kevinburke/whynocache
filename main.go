package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var outputRx = regexp.MustCompile(`ok  \\t([a-zA-Z][\S]+)\\t(\(cached\)|[0-9s.]+)`)

func main() {
	ctx := context.Background()
	flag.Parse()
	args := flag.Args()
	gover, err := exec.Command("go", "version").CombinedOutput()
	if err != nil {
		log.Fatalf("error retrieving go version: %v", err)
	}
	goVersion := strings.TrimSpace(strings.TrimPrefix(string(gover), "go version"))
	fmt.Println("goVersion", goVersion)
	if len(args) > 0 && args[0] == "make" {
		// run1
		cmd := exec.CommandContext(ctx, os.Args[1], os.Args[2:]...)
		cmd.Env = append(os.Environ(), "GODEBUG=gocachehash=1", "GOFLAGS='-json'")
		var buf1 bytes.Buffer
		var testbuf1 bytes.Buffer
		cmd.Stdout = &testbuf1
		cmd.Stderr = &buf1
		if err := cmd.Run(); err != nil {
			io.Copy(os.Stderr, &buf1)
			log.Fatal(err)
		}
		// run2
		cmd2 := exec.CommandContext(ctx, os.Args[1], os.Args[2:]...)
		cmd2.Env = append(os.Environ(), "GODEBUG=gocachehash=1", "GOFLAGS='-json'")
		var buf2 bytes.Buffer
		var testbuf2 bytes.Buffer
		cmd2.Stdout = &testbuf2
		cmd2.Stderr = &buf2
		if err := cmd2.Run(); err != nil {
			io.Copy(os.Stderr, &buf2)
			log.Fatal(err)
		}
		// io.Copy(os.Stdout, &buf1)
		// io.Copy(os.Stdout, &buf2)
		scanner := bufio.NewScanner(&testbuf2)
		cached := make(map[string]struct{})
		uncached := make(map[string]struct{})
		for scanner.Scan() {
			line := scanner.Text()
			matches := outputRx.FindStringSubmatch(line)
			if matches != nil {
				if strings.Contains(line, "(cached)") {
					cached[matches[1]] = struct{}{}
				} else {
					uncached[matches[1]] = struct{}{}
				}
			}
		}
		scanner = bufio.NewScanner(&buf1)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "HASH[testInputs]: ") {
				rest := line[len("HASH[testInputs]: "):]
				if strings.HasPrefix(goVersion, strings.Trim(rest, `"`)) {
					continue
				}
				fmt.Println("rest", rest)
			}
		}
	} else {
	}
	fmt.Println("args", args)
}
