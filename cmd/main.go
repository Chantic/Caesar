package cmd

import (
	"Caesar/logic/coder"
	"flag"
	"fmt"
)

// Цезарь это структура, содержит внутри алфавит, для использования де/энкрипта
// coder is a struct, holding alphabet being used, alongside with Encrypt and Decrypt functions

func main() {
	const EnglishAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var text = flag.String("encrypt", "-", "Target value")
	var shifts = flag.Uint("shift", 0, "Target value")
	var text1 = flag.String("decrypt", "-", "Target value")
	//var shifts1 = flag.Uint("shift", 0, "Target value")
	caesar := coder.NewCaesar(EnglishAlphabet)
	flag.Parse()
	fmt.Println("original text:", *text)
	fmt.Println("shift :", *shifts)
	fmt.Println("original text:", *text1)
	fmt.Println(caesar.Encrypt(*text, *shifts))
	fmt.Println(caesar.Decrypt(*text1, *shifts))
}
