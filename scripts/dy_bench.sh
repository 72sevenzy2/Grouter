#!/bin/bash

cd .. || exit 1
cd benchmarks || exit 1
go test -bench="^(BenchmarkGrouterDynamicRoutes|BenchmarkGolangStdlibRouterDynamicRoutes)$" -run=^$ -benchmem -cpuprofile=profiles/1cpu.prof -memprofile=profiles/1mem.prof -count=10 > results/dynamicBenchmarks.txt
benchstat results/dynamicBenchmarks.txt

go tool pprof profiles/1mem.prof
