package main

import (
	"Caesar/logic/coder"
	"flag"
	"fmt"
	"strings"
)

// Цезарь это структура, содержит внутри алфавит, для использования де/энкрипта
// coder is a struct, holding alphabet being used, alongside with Encrypt and Decrypt functions

func main() {
	const EnglishAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var text = flag.String("encrypt", "", "Target value")
	var shifts = flag.Uint("shift", 0, "Target value")
	var text1 = flag.String("decrypt", "", "Target value")

	caesar := coder.NewCaesar(EnglishAlphabet)
	flag.Parse()
	fmt.Println("original text:", *text)
	fmt.Println("shift :", *shifts)
	fmt.Println("original text:", *text1)
	if valid(*text1, EnglishAlphabet) == false {
		fmt.Println("ошибка ")
	} else {
		fmt.Println(caesar.Encrypt(*text, *shifts))
		fmt.Println(caesar.Decrypt(*text1, *shifts))
	}
	if valid(*text, EnglishAlphabet) == false {
		fmt.Println("ошибка ")
	} else {
		fmt.Println(caesar.Encrypt(*text, *shifts))
		fmt.Println(caesar.Decrypt(*text1, *shifts))
	}
	//fmt.Println(caesar.Encrypt(*text, *shifts))
	//fmt.Println(caesar.Decrypt(*text1, *shifts))
}
func valid(data string, alphabet string) bool {
	dataList := strings.Split(data, "")
	alphabetList := strings.Split(alphabet, "")
	for _, value1 := range dataList {
		for _, value2 := range alphabetList {
			//log.Printf("value:%v, value2:%v\n", value1, value2)
			if value1 == value2 {

				return true
			}
		}

	}
	return false
}
