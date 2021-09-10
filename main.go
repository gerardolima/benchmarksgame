// The Computer Language Benchmarks Game
// http://benchmarksgame.alioth.debian.org/
//
// Go adaptation of binary-trees Rust #4 program
// Use semaphores to match the number of workers with the CPU count
//
// contributed by Marcel Ibes
// modified by Isaac Gouy
// modified by Adam Shaver(use the struct constructor for bottomUpTree)
//
//

package main

import (
	"benchmarksgame/binarytrees/_8"
	"flag"
	"strconv"
)

func main() {
	n := 0
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	_8.Run(uint32(n))
}
