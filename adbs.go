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
	// Check adb's availability
	if !hasAdb() {
		fmt.Printf(`'adb' command not found.
The 'adbs' tool uses adb(Android Debug Bridge).
Please install the Android SDK and add
the directory which has adb command to
the environment variable 'PATH'.
`)
		os.Exit(1)
	}

	// Get args and show usage
	serial := flag.String("s", "", "Serial number(forward match)")
	needHelp := flag.Bool("h", false, "Show this message")
	flag.Parse()
	if *needHelp {
		fmt.Printf(`Usage: adbs [[OPTIONS] ADB_COMMAND|-h]
  -h            - Show this help
  OPTIONS:
    -s SERIAL   - Serial number of target device.
                  You don't need to input complete serial number.
                  Just part of it is okay. (forward match)
  ADB_COMMAND   - command string to pass to the device.
`)
		os.Exit(1)
	}

	var matched string

	// Get device serial from input
	if *serial == "" {
		c := exec.Command("adb", "devices")
		stdout, err := c.StdoutPipe()
		if err != nil {
			fmt.Println("Failed to check devices")
		}
		c.Start()
		r := bufio.NewReader(stdout)
		candidates := []string{}
		err = nil
		for i := 1; err == nil; {
			var line []byte
			line, _, err = r.ReadLine()
			sline := string(line)
			if strings.HasPrefix(sline, "List of devices attached") {
				continue
			}
			if sline == "" {
				continue
			}
			candidateSerial := regexp.MustCompile("^[0-9a-zA-Z\\.:]+").FindString(sline)
			candidates = append(candidates, candidateSerial)
			i++
		}
		if len(candidates) == 0 {
			fmt.Println("No device attached")
			os.Exit(1)
		} else if len(candidates) == 1 {
			// This is the only device attached
			matched = candidates[0]
		} else {
			for i := range candidates {
				fmt.Printf("[%d] %s\n", i+1, candidates[i])
			}
			var input int
			fmt.Printf("Device to execute command: ")
			fmt.Scanf("%d", &input)
			if 1 <= input && input <= len(candidates) {
				matched = candidates[input-1]
				fmt.Printf("Specified: %s\n", matched)
			} else {
				fmt.Printf("Invalid number: %d\n", input)
				os.Exit(1)
			}
		}
		if matched == "" {
			fmt.Println("Serial not specified")
			os.Exit(1)
		}
	} else {
		// Find specified device
		c := exec.Command("adb", "devices")
		stdout, err := c.StdoutPipe()
		if err != nil {
			fmt.Println("Failed to check devices")
		}
		c.Start()
		r := bufio.NewReader(stdout)
		err = nil
		candidates := []string{}
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
			regex := regexp.MustCompile("^" + *serial + "[0-9a-zA-Z\\.:]*")
			candidateMatched := regex.FindString(sline)
			if candidateMatched != "" {
				candidates = append(candidates, candidateMatched)
			}
		}
		if len(candidates) == 0 {
			fmt.Println("Specified device not found\n")
			os.Exit(1)
		} else if 1 < len(candidates) {
			fmt.Println("Multiple candidate devices found:")
			for i := range candidates {
				fmt.Printf("[%d] %s\n", i+1, candidates[i])
			}
			os.Exit(1)
		}
		matched = candidates[0]
		fmt.Printf("adbs: serial: %s\n", matched)
	}

	// Give adb command to device
	c := exec.Command("adb", append([]string{"-s", matched}, flag.Args()...)...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()
}

func hasAdb() bool {
	path, err := exec.LookPath("adb")
	return path != "" && err == nil
}
