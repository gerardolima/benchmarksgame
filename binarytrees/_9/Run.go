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

package _9

import (
	"fmt"
	"math"
	"runtime"
	"sort"
)

type Tree struct {
	Left  *Tree
	Right *Tree
}

func (t *Tree) Count() int {
	if t.Left != nil {
		return 1 + t.Right.Count() + t.Left.Count()
	}
	return 1
}

func NewTree(depth int) *Tree {
	if depth > 0 {
		return &Tree{Left: NewTree(depth - 1), Right: NewTree(depth - 1)}
	} else {
		return &Tree{}
	}
}

type message struct {
	Pos  int
	Text string
}

type ByPos []message

func (m ByPos) Len() int           { return len(m) }
func (m ByPos) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ByPos) Less(i, j int) bool { return m[i].Pos < m[j].Pos }

const minDepth = 4

func Run(n int) {
	cpuCount := runtime.NumCPU()

	maxDepth := n
	if minDepth+2 > n {
		maxDepth = minDepth + 2
	}

	depth := maxDepth + 1

	messages := make(chan message, cpuCount)
	expected := 2 // initialize with the 2 summary messages

	go func() {
		// do stretch tree and longLivedTree

		go func() {
			tree := NewTree(depth)
			messages <- message{0,
				fmt.Sprintf("stretch tree of depth %d\t check: %d",
					depth, tree.Count())}
		}()

		go func() {
			longLivedTree := NewTree(maxDepth)
			messages <- message{math.MaxInt,
				fmt.Sprintf("long lived tree of depth %d\t check: %d",
					maxDepth, longLivedTree.Count())}
		}()

		for halfDepth := minDepth / 2; halfDepth < maxDepth/2+1; halfDepth++ {
			depth := halfDepth * 2
			iterations := 1 << (maxDepth - depth + minDepth)
			expected++

			go func(depth, iterations int) {
				chk := 0
				for i := 0; i < iterations; i++ {
					a := NewTree(depth)
					chk += a.Count()
				}
				m := fmt.Sprintf("%d\t trees of depth %d\t check: %d", iterations, depth, chk)

				messages <- message{depth, m}
			}(depth, iterations)

		}
	}()

	var sortedMsg []message
	for m := range messages {
		sortedMsg = append(sortedMsg, m)
		expected--
		if expected == 0 {
			close(messages)
		}
	}

	sort.Sort(ByPos(sortedMsg))
	for _, m := range sortedMsg {
		fmt.Println(m.Text)
	}
}

/*
func main() {
	n := 0
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}

	Run8(uint(n))
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
