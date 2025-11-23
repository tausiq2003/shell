package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)

func typeCheck(cmd string) (string, bool) { // only the command

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
	sbmap := map[string]struct{}{"echo": struct{}{}, "[": struct{}{}, "type": struct{}{}, "exit": struct{}{}, "pwd": struct{}{}}
	_, exists := sbmap[cmd]
	if exists {
		return fmt.Sprintf("%v is a shell builtin\n", cmd), true
	} else {

		pathStr := os.Getenv("PATH")
		pathList := strings.Split(pathStr, ":")
		// now we have path and cmd, now we have to search for it
		for _, path := range pathList {

			dir, _ := os.Open(path)
			defer dir.Close()
			files, _ := dir.Readdir(-1)
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
						return fmt.Sprintf("%v is %v/%v\n", cmd, dir.Name(), cmd), true
					}
				}

			}
		}

	}
	return fmt.Sprintf("%v not found\n", cmd), false
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
		cmdList := strings.Split(cmd, " ")
		if cmdList[0] == "exit" {
			if len(cmdList) == 1 {
				os.Exit(0)
			}
			exitCode := cmdList[1]
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
		if cmdList[0] == "echo" {
			if cmd == "echo" {
				os.Exit(0)
			}
			fmt.Println(cmd[5:])
			continue

		}
		if cmdList[0] == "type" {
			data, _ := typeCheck(cmdList[1])
			fmt.Print(data)
			continue

		}
		if cmdList[0] == "pwd" {
			if len(cmdList) > 1 {
				log.Fatal("pwd: too many arguments")
			}
			dir, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(dir)
			continue
		}
		if cmdList[0] == "cd" {
			// lets see it in future
			//			if len(cmdList) == 1 {
			//				os.Chdir("~")
			//				continue
			//			}
			if err := os.Chdir(cmdList[1]); err != nil {
				log.Fatalf("cd: %v: No such file or directory", cmdList[1])
			}
			continue
		}
		if len(cmdList) > 0 {

			_, exists := typeCheck(cmdList[0])
			if exists {
				execute := exec.Command(cmdList[0], cmdList[1:]...)
				execute.Stdout = os.Stdout
				execute.Stderr = os.Stderr
				if err := execute.Run(); err != nil {
					log.Fatal(err)
				}
				continue
			}
		}

		fmt.Printf("%v: command not found\n", strings.Split(cmd, " ")[0])

	}

}
