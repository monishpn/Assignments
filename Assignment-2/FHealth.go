package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go FILE")
		os.Exit(1)
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File Not Present")
		os.Exit(1)
	}

	info := 0
	er := 0
	war := 0
	//
	var strArr string

	defer file.Close()

	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {
		strArr = Scanner.Text()
		switch {
		case strings.HasPrefix(strArr, "[INFO]"):
			info++
		case strings.HasPrefix(strArr, "[WARNING]"):
			war++
		case strings.HasPrefix(strArr, "[ERROR]"):
			er++
		}

	}

	//Scanner.Split(bufio.ScanWords)
	//for Scanner.Scan() {
	//	strArr = append(strArr, Scanner.Text())
	//}
	//
	//for _, word := range strArr {
	//	switch word {
	//	case "[INFO]":
	//		info += 1
	//	case "[WARNING]":
	//
	//		war += 1
	//	case "[ERROR]":
	//		er += 1
	//
	//	}
	//}
	//
	fmt.Printf("Log Analysis of File : %s \n\n INFO : %d\n WARNING : %d\n ERROR : %d\n\n  ", fileName, info, war, er)

}
