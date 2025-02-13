package main

import (
	"log"
	"os"

	"k0s.io/pkg/cli/chassis"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	chassis.Run(os.Args[1:])
}
