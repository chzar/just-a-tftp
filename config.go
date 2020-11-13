package main

import (
	"flag"
)

type config struct {
	Directory        string
	Readonly         bool
	ConnectionString string
}

func argparse() *config {
	readonly := flag.Bool("ro", false, "runs the server in readonly mode")
	directory := flag.String("dir", "./", "the directory to serve files from")
	connectionString := flag.String("conns", ":69", "server connection string")
	flag.Parse()

	c := &config{
		Directory:        *directory,
		Readonly:         *readonly,
		ConnectionString: *connectionString,
	}
	return c
}
