package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

// attaching a method to a datatype or class.
func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}
	return levelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log Level : %d, %s\n", level, level.String())
}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
	printLogLevel(10)
}
