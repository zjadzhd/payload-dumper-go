package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var (
		list            bool
	
	)
	flag.BoolVar(&list, "l", false, "Show list of partitions in payload.bin (shorthand)")
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}
	filename := flag.Arg(0)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s\n", filename)
	}

	payloadBin := filename
	payload := NewPayload(payloadBin)
	if err := payload.Open(); err != nil {
		log.Fatal(err)
	}
	payload.Init()

	if list {
		return
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] [inputfile]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}