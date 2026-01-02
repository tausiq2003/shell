package main

import (
	"fmt"
	"strings"
	//"strings"
)

func Echo(cmdList []string) (string, error) {

	content := fmt.Sprint(cmdList[1:])
	realContent := content[1 : len(content)-1]
	//handle single quote and double quotes, ig they work same,else write

	if (realContent[0] == '\'' && realContent[len(realContent)-1] == '\'') || (realContent[0] == '"' && realContent[len(realContent)-1] == '"') {
		// just going with the tests, there is another case, if you type one " or '  where > prompts
		trimmedContent := realContent[1 : len(realContent)-1]
		trimmedContent = strings.ReplaceAll(trimmedContent, "'", "")
		trimmedContent = strings.ReplaceAll(trimmedContent, `"`, "")
		//proceed

		return trimmedContent, nil

	}

	realContent = strings.ReplaceAll(realContent, "'", "")
	realContent = strings.ReplaceAll(realContent, `"`, "")

	realContentWithoutQuotes := strings.Join(strings.Fields(realContent), " ")

	return realContentWithoutQuotes, nil

}
