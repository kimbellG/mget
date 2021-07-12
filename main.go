package main

import (
	"flag"
	"fmt"
	"log"
	"mget/client"
	"os"
)

var (
	filename = flag.String("f", "", "output file for result")
)

func main() {
	flag.Parse()

	url, err := getURL()
	if err != nil {
		log.Fatalf("mget: %v", err)
	}

	dst, err := getDst()
	if err != nil {
		log.Fatalf("mget: create file: %v", err)
	}
	defer dst.Close()

	if err := client.Download(url, dst); err != nil {
		log.Fatalf("mget: %v", err)
	}
}

func getURL() (string, error) {
	urls := flag.Args()
	if len(urls) != 1 {
		return "", fmt.Errorf("incorrect usage: required only one link")
	}

	return urls[0], nil
}

func getDst() (*os.File, error) {
	if *filename == "" {
		return os.Stdout, nil
	}

	return os.Create(*filename)
}
