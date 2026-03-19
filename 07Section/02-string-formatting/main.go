package main

import (
	"errors"
	"fmt"
)

type ConfigItem struct {
	key   string
	value interface{}
	IsSet bool
}

/*
	 %v - the default formatting
	%+v -
	%#v
	%T
	%s
	%d
	%f (%.2f)
	%t

	%q
	%%
*/

func main() {

	appName := "EnvParser"
	version := 1.2
	port := 8000
	isEnable := true

	status := fmt.Sprintf("application %s (version:%.1f) is running on Port: %d. Enabled: %t",
		appName, version, port, isEnable)

	fmt.Println(status)

	item2 := ConfigItem{key: "APP_URL", value: "http://localhost:8000/", IsSet: false}
	item1 := ConfigItem{key: "Timeout_MS", value: 4000, IsSet: true}

	//item3:= ConfigItem{key:"DEBUG_MODE", value:false, IsSet:false}

	fmt.Println(item2)
	fmt.Printf("item1(%%v) : %v\n", item1)
	fmt.Printf("item1(%%+v) : %+v\n", item1)
	fmt.Printf("item1(%%#v) : %#v\n", item1)

	err := errors.New("test error")
	err = fmt.Errorf("here is the error on port %d: %w", port, err)
	fmt.Println(err)

}
