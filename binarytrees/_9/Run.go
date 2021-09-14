// The Computer Language Benchmarks Game
// http://benchmarksgame.alioth.debian.org/
//
// Go implementation of binary-trees, based on the reference implementation
// gcc #3, on Go #8 (which is based on Rust #4) as the following links, below:
// - https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-gcc-3.html
// - https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-go-8.html
// - https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-rust-4.html
//
// Comments on code aimed to be analogous as those in the reference implementation.
//
// Contributed by Gerardo Lima (https://github.com/gerardolima)
// Based on previous work from Adam Shaver, Isaac Gouy, Marcel Ibes, Jeremy
//  Zerfas, Jon Harrop, Alex Mizrahi, Bruno Coutinho, Volodymyr M. Lisivka, Tom
//  Kaitchuck, Matt Brubeck, Cristi Cobzarenco, ...
//

package _9

import (
	"fmt"
	"sync"
)

type Tree struct {
	Left  *Tree
	Right *Tree
}

// Count the nodes in the given complete binary tree `tree`.
func (t *Tree) Count() int {

	// As `tree` is expected to be complete, only test it's left side to check
	// whether is a leaf-node. Current implementation is recursive = current node
	// + count of its left side + count of its right side.
	if t.Left != nil {
		return 1 + t.Right.Count() + t.Left.Count()
	}
	return 1
}

// Create a complete binary tree of `depth` and return it as a pointer.
func NewTree(depth int) *Tree {
	if depth > 0 {
		return &Tree{Left: NewTree(depth - 1), Right: NewTree(depth - 1)}
	} else {
		return &Tree{}
	}
}

func Run(maxDepth int) {

	var wg sync.WaitGroup

	// Set minDepth to 4 and maxDepth to the maximum of maxDepth and minDepth +2.
	const minDepth = 4
	if maxDepth < minDepth+2 {
		maxDepth = minDepth + 2
	}

	produced := 0
	forests := 3 + (maxDepth-minDepth)/2
	outputBuffer := make([]string, forests)

	// Create binary tree of depth maxDepth+1, compute its Count and set the
	// first position of the outputBuffer with its statistics.
	wg.Add(1)
	go func() {
		tree := NewTree(maxDepth + 1)
		m := fmt.Sprintf("stretch tree of depth %d\t check: %d", maxDepth+1, tree.Count())

		outputBuffer[0] = m
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
		produced++

		wg.Add(1)
		go func(depth, iterations, index int) {
			chk := 0
			for i := 0; i < iterations; i++ {
				// Create a binary tree of depth and accumulate total counter with its
				// node count.
				a := NewTree(depth)
				chk += a.Count()
			}
			m := fmt.Sprintf("%d\t trees of depth %d\t check: %d", iterations, depth, chk)

			outputBuffer[index] = m
			wg.Done()
		}(depth, iterations, produced)
	}

	wg.Wait()

	// Compute the checksum of the long-lived binary tree that we created
	// earlier and store its statistics.
	m := fmt.Sprintf("long lived tree of depth %d\t check: %d", maxDepth, longLivedTree.Count())
	outputBuffer[forests-1] = m

	// Print the statistics for all of the various tree depths.
	for _, m := range outputBuffer {
		fmt.Println(m)
	}
}

/*
func main() {
	n := 0
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	Run(n)
}
*/

// notes, command-line, and program output
// NOTES:
// 64-bit Ubuntu quad core
// go version go1.17 linux/amd64

// Mon, 16 Aug 2021 22:01:10 GMT

// MAKE:
// /opt/src/go1.17/go/bin/go build -o binarytrees.go-8.go_run binarytrees.go-8.go

// 3.72s to complete and log all make actions

// COMMAND LINE:
// ./binarytrees.go-8.go_run 21

// PROGRAM OUTPUT:
// stretch tree of depth 22	 check: 8388607
// 2097152	 trees of depth 4	 check: 65011712
// 524288	 trees of depth 6	 check: 66584576
// 131072	 trees of depth 8	 check: 66977792
// 32768	 trees of depth 10	 check: 67076096
// 8192	 trees of depth 12	 check: 67100672
// 2048	 trees of depth 14	 check: 67106816
// 512	 trees of depth 16	 check: 67108352
// 128	 trees of depth 18	 check: 67108736
// 32	 trees of depth 20	 check: 67108832
// long lived tree of depth 21	 check: 4194303
