package main

import (
        "log"
        "os"

	"k0s.io/pkg/cli/caddy"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Fatalln(caddy.Run(os.Args[1:]))
}
