package main

import (
	"log"
	"os"

	"k0s.io/pkg/cli/goproxy"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Fatalln(agent.Run(os.Args[1:]))
}
