package main

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)

import (
	"fmt"
	"os"
	"strings"
)

func TypeCheck(cmd string) (string, bool) { // only the command

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
