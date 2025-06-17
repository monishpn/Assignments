package main

import "fmt"

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(msg string) {
	fmt.Println("\n\nCONSOLE: ", msg)
}

type FileLogger struct{}

func (FileLogger) Log(msg string) {
	var s []string
	s = append(s, msg)
	fmt.Println("FILE: ", s)
}

type RemoteLogger struct{}

func (RemoteLogger) Log(msg string) {
	fmt.Printf("REMOTE: %s\n", msg)
}

func LogAll(log []Logger, message string) {

	for _, l := range log {
		switch l.(type) {
		case ConsoleLogger:
			l.Log(message)
		case FileLogger:
			l.Log(message)
		case RemoteLogger:
			l.Log(message)
		}
	}
}

func main() {
	fmt.Println("Enter The Message")
	var msg string
	_, err := fmt.Scanln(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	LogAll([]Logger{ConsoleLogger{}, FileLogger{}, RemoteLogger{}}, msg)
}
