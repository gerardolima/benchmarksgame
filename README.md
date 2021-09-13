# // benchmarksgame

https://benchmarksgame-team.pages.debian.net/benchmarksgame/

## How to
- run official tests `go test -timeout 30s -run ^ExampleRun10$ benchmarksgame/binarytrees/_8`
- run all tests `go test benchmarksgame/binarytrees/_8`
- run performance tests `go test -benchmem -run=^$ -bench ^(BenchmarkRun21)$ benchmarksgame/binarytrees/_8`
- TDD `fswatch . | xargs -n 1 -I{} go test -benchmem -run ^ExampleRun10$ -bench '^(BenchmarkRun21)$'`

## More Info
Original source downloaded form:
- [binarytrees-go-8](https://benchmarksgame-team.pages.debian.net/benchmarksgame/program/binarytrees-go-8.html)
