package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ranjeetsinghbnl/go-learning/greetings"
)

func main() {
	flag.Parse()
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	messages, err := greetings.Hellos(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
