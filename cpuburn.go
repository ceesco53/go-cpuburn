package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

var (
	numBurn        int
	updateInterval int
)

func cpuBurn() {
	for {
	}
}

func init() {
	flag.IntVar(&numBurn, "n", 0, "number of cores to burn (0 = all)")
	flag.IntVar(&updateInterval, "u", 10, "seconds between updates (0 = don't update)")
	flag.Parse()
	if numBurn <= 0 {
		numBurn = runtime.NumCPU()
	}
}

func main() {
	runtime.GOMAXPROCS(numBurn)
	fmt.Printf("Burning %d CPUs/cores\n", numBurn)
	for i := 0; i < numBurn; i++ {
		go cpuBurn()
	}
	t := time.Tick(time.Duration(updateInterval) * time.Second)
	for secs := updateInterval; ; secs += updateInterval {
		<-t
		fmt.Printf("%d seconds\n", secs)
	}
}