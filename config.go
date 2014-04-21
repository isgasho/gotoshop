package main

import (
	"flag"
	"fmt"
)

// Server config
type Config struct {
	Addr string
}

// Parses command line args and returns new Config
func NewConfig() Config {
	portPtr := flag.Int("port", 3000, "port")

	flag.Parse()

	return Config{
		Addr: fmt.Sprintf(":%d", *portPtr),
	}
}
