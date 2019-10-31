package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sim"
	"strconv"
	"strings"
)

func main() {
	var debug bool
	var stat bool
	args := os.Args[1:]
	if len(args) < 5 {
		fmt.Println("usage: ./SIM <CACHE_SIZE> <ASSOC> <REPLACEMENT> <WB> <TRACE_FILE>")
		return
	}

	size, err := strconv.ParseUint(args[0], 10, 64)
	assoc, err := strconv.ParseUint(args[1], 10, 64)
	repl, err := strconv.ParseInt(args[2], 10, 64)
	wb, err := strconv.ParseInt(args[3], 10, 64)
	traceFile := args[4]
	if len(args) > 5 {
		debug, err = strconv.ParseBool(args[5])
		if err != nil {
			debug = false
		}
	}

	if len(args) > 6 {
		stat, err = strconv.ParseBool(args[6])
		if err != nil {
			stat = false
		}
	}

	file, err := os.Open(traceFile)
	if err != nil {
		fmt.Printf("sim: unable to open trace file - %v\n", traceFile)
		return
	}
	defer file.Close()

	// Options
	options := &sim.Options{
		Size:        uint64(size), // Bytes
		Assoc:       int(assoc),   // n-way
		BlockSize:   64,           // Bytes
		Replacement: int(repl),    // LRU
		WriteBack:   int(wb),      // Write Through/Write Back
		Debug:       debug,        // Debug Tracing

	}

	// Starat Cache
	cache := &sim.Cache{Options: options}
	cache.Init()

	// Execute loop
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var op int
		line := scanner.Text()
		lines := strings.Split(line, " ")
		if lines[0] == "W" {
			op = 0
		} else if lines[0] == "R" {
			op = 1
		} else {
			log.Fatal(fmt.Sprintf("Incorrect op: %v\n", line))
		}
		err := cache.Execute(op, lines[1])
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
	}

	if stat {
		fmt.Printf(
			"%v,%v,%v\n",
			cache.Stats.MissRatio(),
			cache.Stats.Writes,
			cache.Stats.Reads,
		)
	} else {
		fmt.Printf(
			"Miss Ratio: %v\nWrites: %v\nReads: %v\n",
			cache.Stats.MissRatio(),
			cache.Stats.Writes,
			cache.Stats.Reads,
		)
	}

}
