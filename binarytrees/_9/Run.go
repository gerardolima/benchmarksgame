// The Computer Language Benchmarks Game
// http://benchmarksgame.alioth.debian.org/
//
// Go implementation of binary-trees, based on the reference implementation
// gcc #3, on Go #8 (which is based on Rust #4) as the following links, below:
// - https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-gcc-3.html
// - https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-go-8.html
// - https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-rust-4.html
//
// Comments aim to be analogous as those in the reference implementation and are
// intentionally verbose, to help programmers unexperienced in GO to understand
// the implementation.
//
// The following alternative implementations were considered before submitting
// this code. All of them had worse readability and didn't yield better results
// on my local machine:
//
// 0. general:
// 0.1 using uint32, instead of int;
//
// 1. func Count:
// 1.1 using a local stack, instead of using a recursive implementation; the
//     performance degraded, even using a pre-allocated slice as stack and
//     manually handling its size;
// 1.2 assigning Left and Right to nil after counting nodes; the idea to remove
//     references to instances no longer needed was to make GC easier, but this
//     did not work as intended;
// 1.3 using a walker and channel, sending 1 on each node; although this looked
//     idiomatic to GO, the performance suffered a lot;
// 2. func NewTree:
// 2.1 allocating all tree nodes on a tree slice upfront and making references
//     to those instances, instead of allocating two sub-trees on each call;
//     this did not improve performance;
//
// Contributed by Gerardo Lima (https://github.com/gerardolima)
// Based on previous work from Adam Shaver, Isaac Gouy, Marcel Ibes Jeremy,
//  Zerfas, Jon Harrop, Alex Mizrahi, Bruno Coutinho, ...
//

package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
)

type Tree struct {
	Left  *Tree
	Right *Tree
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(Tree)
	},
}

func NewNode() *Tree {
	return pool.Get().(*Tree)
}

func FreeNode(t *Tree) {
	pool.Put(t)
}

// Count the nodes in the given complete binary tree.
func CountNodes(t *Tree) int {
	count := 1

	// Only test the Left node (this binary tree is expected to be complete).
	if t.Left != nil {
		count += CountNodes(t.Right) + CountNodes(t.Left)
	}
	FreeNode(t)
	return count
}

// Create a complete binary tree of `depth` and return it as a pointer.
func NewTree(depth int) *Tree {
	t := NewNode()
	if depth > 0 {
		t.Left = NewTree(depth - 1)
		t.Right = NewTree(depth - 1)
	}
	return t
}

func Run(maxDepth int) {

	var wg sync.WaitGroup

	// Set minDepth to 4 and maxDepth to the maximum of maxDepth and minDepth +2.
	const minDepth = 4
	if maxDepth < minDepth+2 {
		maxDepth = minDepth + 2
	}

	// Create an indexed string buffer for outputing the result in order.
	outCurr := 0
	outSize := 3 + (maxDepth-minDepth)/2
	outBuff := make([]string, outSize)

	// Create binary tree of depth maxDepth+1, compute its Count and set the
	// first position of the outputBuffer with its statistics.
	wg.Add(1)
	go func() {
		tree := NewTree(maxDepth + 1)
		msg := fmt.Sprintf("stretch tree of depth %d\t check: %d",
			maxDepth+1,
			CountNodes(tree))

		outBuff[0] = msg
		wg.Done()
	}()

	// Create a long-lived binary tree of depth maxDepth. Its statistics will be
	// handled later.
	var longLivedTree *Tree
	wg.Add(1)
	go func() {
		longLivedTree = NewTree(maxDepth)
		wg.Done()
	}()

	// Create a lot of binary trees, of depths ranging from minDepth to maxDepth,
	// compute and tally up all their Count and record the statistics.
	for depth := minDepth; depth <= maxDepth; depth += 2 {
		iterations := 1 << (maxDepth - depth + minDepth)
		outCurr++

		wg.Add(1)
		go func(depth, iterations, index int) {

			ch := make(chan int, iterations)
			for i := 0; i < iterations; i++ {
				// Create a binary tree of depth and accumulate total counter with its
				// node count.
				a := NewTree(depth)
				ch <- CountNodes(a)
			}
			close(ch)

			acc := 0
			for cc := range ch {
				acc += cc
			}

			msg := fmt.Sprintf("%d\t trees of depth %d\t check: %d",
				iterations,
				depth,
				acc)

			outBuff[index] = msg
			wg.Done()
		}(depth, iterations, outCurr)
	}

	wg.Wait()

	// Compute the checksum of the long-lived binary tree that we created
	// earlier and store its statistics.
	msg := fmt.Sprintf("long lived tree of depth %d\t check: %d",
		maxDepth,
		CountNodes(longLivedTree))
	outBuff[outSize-1] = msg

	// Print the statistics for all of the various tree depths.
	for _, m := range outBuff {
		fmt.Println(m)
	}
}

func main() {
	n := 0
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	Run(n)
}
