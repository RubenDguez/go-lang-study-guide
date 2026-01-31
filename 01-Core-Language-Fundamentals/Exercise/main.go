package main

import "fmt"

type LogLevel int

const (
	Trace LogLevel = iota
	Debug 
	Info 
	Warn 
	Error 
)

var levelNames = []string{"Trace", "Debug", "Info", "Warn", "Error"}

func (level LogLevel) ToString() string {
	if level < Trace || level > Error {
		return "unknown"
	}

	return levelNames[level]
}

func printLogLevel (level LogLevel) {
	fmt.Printf("%d [ %s ]\n", level, level.ToString())
}

func main() {
	printLogLevel(Trace)
	printLogLevel(Debug)
	printLogLevel(Info)
	printLogLevel(Warn)
	printLogLevel(Error)
	printLogLevel(10)
}
