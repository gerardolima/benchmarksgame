#!/usr/local/bin/bash

function BenchmarkRun21 () {
  go test -benchmem -run=^$ -bench '^(BenchmarkRun21)$' . | grep '^ok'
}

for (( c=0; c<=5; c++ )); do
  BenchmarkRun21
done
