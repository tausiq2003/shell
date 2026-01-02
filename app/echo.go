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
	emitted := false
	inArgument := false

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
				if emitted && !inArgument {
					result += " "
				}
				result += strings.Join(strings.Fields(withoutQuotesContent), " ")
				withoutQuotesContent = ""
				emitted = true
				inArgument = true
			}
			quotesContent += string(rune(v))
		} else if !opened {
			if v == '\'' {
				// then it is just closed
				if quotesContent != "" {
					if emitted && !inArgument {
						result += " "
					}
					result += quotesContent
					quotesContent = ""
					emitted = true
					inArgument = true
					//				fmt.Printf("--->%v\n", result)
					//				fmt.Println("baz")
				}
			} else {
				//treat it as normal thing
				if v == ' ' {
					inArgument = false
				}
				withoutQuotesContent += string(rune(v))
				//				fmt.Println("foo")
				// fmt.Println(withoutQuotesContent)
			}
		}
	}
	//	fmt.Println(quotesContent)
	//	fmt.Printf("%q\n", withoutQuotesContent)

	if withoutQuotesContent != "" {
		if emitted && !inArgument {
			result += " "
		}
		result += strings.Join(strings.Fields(withoutQuotesContent), " ")
	}

	return result, nil
}
