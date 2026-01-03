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
				continue
			}

			data, err := Echo(cmdList)
			if err != nil {
				log.Fatal("Error", err)
			}
			fmt.Println(data)
			continue

		}
		//		if cmdList[0] == "cat" {
		//			if cmd == "cat" {
		//				continue
		//			}
		//		}
		if cmdList[0] == "type" {
			data, _ := TypeCheck(cmdList[1])
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
			if len(cmdList) == 1 {
				os.Chdir(os.Getenv("HOME"))
				continue
			}
			if cmdList[1] == "~" {
				cmdList[1] = os.Getenv("HOME")
			}
			if err := os.Chdir(cmdList[1]); err != nil {
				fmt.Printf("cd: %v: No such file or directory\n", cmdList[1])
				// 	os.Exit(1) should i do it?
			}
			continue
		}
		if len(cmdList) > 0 {

			_, exists := TypeCheck(cmdList[0])
			if exists {
				data, err := Echo(cmdList)
				if err != nil {
					log.Fatal(err)
				}
				execute := exec.Command(cmdList[0], data)
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
