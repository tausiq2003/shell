package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")
	reader := bufio.NewScanner(os.Stdin)
	reader.Split(bufio.ScanBytes)
	var typedCmd string
	for reader.Scan() {
		tokens := reader.Text()
		if tokens != "\n" {
			typedCmd += tokens

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
	fmt.Printf("%v: command not found", strings.Split(typedCmd, " ")[0])

}
