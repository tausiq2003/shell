package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)

func typeCheck(cmd string) { // only the command

	//	commands := []string{"exit", "type", "echo"}
	//ignore
	//	command := strings.Split(cmd, " ")[1]
	//	if slices.Contains(commands, command) {
	//		fmt.Printf("%v is a shell builtin\n", command)
	//	} else {
	//		fmt.Printf("%v: not found\n", command)
	//	}
	// so first take the cmd, compute the path
	//no first check if the shell-builtin in hash set, O(1)
	sbmap := map[string]struct{}{"echo": struct{}{}, "[": struct{}{}, "type": struct{}{}, "exit": struct{}{}}
	_, exists := sbmap[cmd]
	if exists {
		fmt.Printf("%v is a shell builtin\n", cmd)
	} else {

		pathStr := os.Getenv("PATH")
		pathList := strings.Split(pathStr, ":")
		// now we have path and cmd, now we have to search for it
		flag := 0
		for _, path := range pathList {

			dir, _ := os.Open(path)
			defer dir.Close()
			files, err := dir.Readdir(-1)
			if err != nil {
				log.Fatal("Error: ", err)
			}
			for _, file := range files {
				if file.Name() == cmd {

					if file.Mode()&73 != 0 {
						/* very important thing, it checks whether the file is executable or not. how?
						first we have permission bits as follows
						111 111 111
						if we & with 73 which is 001001001
						it will cancel all the bits r & w and return only non zero number
						if its only 755 like here it would return 73 else 64, 8 or 1
						*/
						fmt.Printf("%v is %v\n", cmd, dir.Name())
						flag = 1
						break

					}
				}

			}
			if flag == 1 {
				break

			}
		}
		if flag == 0 {
			fmt.Printf("%v not found\n", cmd)

		}

	}
}

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
				log.Fatal("Error", err)
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
		if strings.Split(cmd, " ")[0] == "type" {
			typeCheck(strings.Split(cmd, " ")[1])
			continue

		}
		fmt.Printf("%v: command not found\n", strings.Split(cmd, " ")[0])

	}

}
