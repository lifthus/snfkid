package main

import (
	"flag"
	"log"
)

func main() {
	port := flag.String("port", "", "port to listen on")
	machine := flag.Int("machine", -1, "machine ID")

	switch {
	case *port != "":
		// open snowflake server
	case *machine != -1:
		// just generate a snowflake for the given machine ID
	default:
		log.Fatal("either port or machine ID must be specified")
	}
}
