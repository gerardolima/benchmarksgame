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
	"sync"
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

func Run(maxDepth int) {

	var wg sync.WaitGroup

	messages := make(chan message, runtime.NumCPU())
	expected := 2 // initialize with the 2 summary messages

	const minDepth = 4
	if maxDepth < minDepth+2 {
		maxDepth = minDepth + 2
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		tree := NewTree(maxDepth + 1)
		m := fmt.Sprintf("stretch tree of depth %d\t check: %d", maxDepth+1, tree.Count())
		messages <- message{0, m}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		longLivedTree := NewTree(maxDepth)
		m := fmt.Sprintf("long lived tree of depth %d\t check: %d", maxDepth, longLivedTree.Count())
		messages <- message{math.MaxInt, m}
	}()

	for depth := minDepth; depth <= maxDepth; depth += 2 {
		iterations := 1 << (maxDepth - depth + minDepth)
		expected++

		wg.Add(1)
		go func(depth, iterations int) {
			defer wg.Done()
			chk := 0
			for i := 0; i < iterations; i++ {
				a := NewTree(depth)
				chk += a.Count()
			}
			m := fmt.Sprintf("%d\t trees of depth %d\t check: %d", iterations, depth, chk)

			messages <- message{depth, m}
		}(depth, iterations)
	}

	wg.Wait()
	close(messages)

	var sortedMsg []message
	for m := range messages {
		sortedMsg = append(sortedMsg, m)
		expected--
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
