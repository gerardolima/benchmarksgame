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

package _8

import "testing"

func ExampleRun10() {
	// expected output from
	// https://benchmarksgame-team.pages.debian.net/benchmarksgame/download/binarytrees-output.txt

	Run(10)
	// Output:
	// stretch tree of depth 11	 check: 4095
	// 1024	 trees of depth 4	 check: 31744
	// 256	 trees of depth 6	 check: 32512
	// 64	 trees of depth 8	 check: 32704
	// 16	 trees of depth 10	 check: 32752
	// long lived tree of depth 10	 check: 2047
}

func ExampleRun21(b *testing.B) {
	// expected output from
	// https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-go-8.html

	Run(21)
	// Output:
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
}
func BenchmarkRun21(b *testing.B) {
	// expected output from
	// https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-go-8.html

	Run(21)
	// Output:
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
}
