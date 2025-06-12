package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File Not Present")
		os.Exit(1)
	}

	info := 0
	er := 0
	war := 0

	defer file.Close()

	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)
	//for Scanner.Scan() {
	//	fmt.Println(Scanner.Text())
	//}
	for Scanner.Scan() {
		switch Scanner.Text() {
		case "[INFO]":
			info += 1
		case "[WARNING]":

			war += 1
		case "[ERROR]":
			er += 1

		}
	}

	fmt.Printf("Log Analysis of File : %s \n\n INFO : %d\n WARNING : %d\n ERROR : %d\n\n  ", fileName, info, war, er)
}
