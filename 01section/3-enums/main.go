package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type loglevel int

const (
	LogInfo loglevel = iota
	LogDebug
	LogWarn
	LogError
	LogFatal
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
