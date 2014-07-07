/*
                      adbs - adb with serial number

         Copyright (c) 2012 Soichiro Kashima. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	if !hasAdb() {
		fmt.Printf(`'adb' command not found.
The 'adbs' tool uses adb(Android Debug Bridge).
Please install the Android SDK and add
the directory which has adb command to
the environment variable 'PATH'.
`)
		os.Exit(1)
	}

	serial := flag.String("s", "", "Serial number(forward match)")
	flag.Parse()
	if *serial == "" {
		fmt.Printf(`Usage: adbs [-s SERIAL] ADB_COMMAND
  SERIAL      - Serial number of target device.
                You don't need to input complete serial number.
                Just part of it is okay. (forward match)
  ADB_COMMAND - command string to pass to the device.
`)
		os.Exit(1)
	}

	c := exec.Command("adb", "devices")
	stdout, err := c.StdoutPipe()
	if err != nil {
		fmt.Println("Failed to check devices")
	}
	c.Start()
	r := bufio.NewReader(stdout)
	var matched string
	err = nil
	for err == nil {
		var line []byte
		line, _, err = r.ReadLine()
		sline := string(line)
		if strings.HasPrefix(sline, "List of devices attached") {
			continue
		}
		if sline == "" {
			continue
		}
		regex := regexp.MustCompile("^" + *serial + "[0-9a-zA-Z]*")
		matched = regex.FindString(sline)
		if matched == "" {
			fmt.Println("Specified device not found\n")
			os.Exit(1)
		}
		fmt.Printf("adbs: serial: %s\n", matched)
		break
	}

	args := retrieveRestArgs()
	c = exec.Command("adb", "-s", matched, args)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()
}

func hasAdb() bool {
	path, err := exec.LookPath("adb")
	return path != "" && err == nil
}

func retrieveRestArgs() string {
	arrayArgs := flag.Args()
	allArgs := ""
	for i := range arrayArgs {
		allArgs = allArgs + arrayArgs[i]
	}
	return allArgs
}
