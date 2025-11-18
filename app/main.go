package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	// TODO: Uncomment the code below to pass the first stage
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// seg: input starts here

		reader := bufio.NewScanner(os.Stdin)
		reader.Split(bufio.ScanBytes)
		var cmd string
		for reader.Scan() {
			tokens := reader.Text()
			if tokens != "\n" {
				cmd += tokens

			}
			if tokens == "\n" {
				break
			}
			//fmt.Println(tokens)
		}
		if err := reader.Err(); err != nil {
			// Your code here
			panic(err)
		}
		//handling commands here
		cmd = strings.TrimSpace(cmd)

		if cmd == "" {
			continue
		}
		if strings.Split(cmd, " ")[0] == "exit" {
			exitCode := strings.Split(cmd, " ")[1]
			intexitCode, err := strconv.Atoi(exitCode)
			if err != nil {
				// Your code here
				panic(err)
			}
			if intexitCode < 0 || intexitCode > 125 {
				fmt.Println("Invalid exit code")
				os.Exit(1)
			}
			os.Exit(intexitCode)

		}
		if strings.Split(cmd, " ")[0] == "echo" {
			fmt.Println(cmd[5:])
			continue

		}
		fmt.Printf("%v: command not found\n", strings.Split(cmd, " ")[0])

	}

}
