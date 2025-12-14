package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var name string

func PrintHello(name string) {
	fmt.Println("Hello, " + name)
}

func main() {
	flag.StringVar(&name, "name", "", "Your name")
	flag.Parse()
	if name == "" {
		log.Println("Please provide your name with -name flag")
		flag.Usage()
		os.Exit(2)
	}
	fmt.Println("Hello, " + name)
}
