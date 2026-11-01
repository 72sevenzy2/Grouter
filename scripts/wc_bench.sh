#!/bin/bash

cd .. || exit 1
cd benchmarks || exit 1

go test -bench="^(BenchmarkGrouterWildcardRoute|BenchmarkGolangStdlibRouterWildCardRoute)$" -run=^$ -benchmem -cpuprofile=profiles/5cpu.prof -memprofile=profiles/5mem.prof -count=10 > results/wildcardRoutesResults.txt
benchstat results/wildcardRoutesResults.txt

go tool pprof profiles/5mem.prof
