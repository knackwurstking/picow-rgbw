package main

import "fmt"

const VERSION = "0.0.1"

func echoVersion() {
	fmt.Printf("%s %s\n", applicationName, VERSION)
}
