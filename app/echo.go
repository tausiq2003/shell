package main

import (
	"fmt"
	"strings"
)

func Echo(cmdList []string) (string, error) {
	content := fmt.Sprint(cmdList[1:])
	realContent := content[1 : len(content)-1]

	opened := false
	// doubleOpened := false
	var result string
	var quotesContent string
	var withoutQuotesContent string
	for _, v := range realContent {
		if v == '\'' {
			//fmt.Println("bar")
			opened = !opened

		}
		//		if v == 39 {
		//			doubleOpened = !doubleOpened
		//		}
		if opened && v != '\'' {
			if withoutQuotesContent != "" {
				result += strings.Join(strings.Fields(withoutQuotesContent), " ")
				result += " "
				withoutQuotesContent = ""
			}
			quotesContent += string(rune(v))
		} else if !opened {

			if v == '\'' {
				// then it is just closed
				result += quotesContent
				quotesContent = ""
				//				fmt.Printf("--->%v\n", result)
				//				fmt.Println("baz")

			} else {
				//treat it as normal thing
				withoutQuotesContent += string(rune(v))
				//				fmt.Println("foo")
				// fmt.Println(withoutQuotesContent)
			}

		}
	}
	//	fmt.Println(quotesContent)
	//	fmt.Printf("%q\n", withoutQuotesContent)

	result += strings.TrimSpace(strings.Join(strings.Fields(withoutQuotesContent), " "))

	result = strings.TrimRight(result, " ")

	return result, nil

}
