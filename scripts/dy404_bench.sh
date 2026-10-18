#!/bin/bash

cd .. || exit 1
cd benchmarks || exit 1
go test -bench="^(BenchmarkGrouterDynamicRoutes404|BenchmarkGolangStdlibRouterDynamicRoutes404)$" -run=^$ -benchmem -cpuprofile=3cpu.prof -memprofile=3mem.prof -count=10 > results/dynamicBenchmarks404.txt
benchstat dynamicBenchmarks404.txt

go tool pprof profiles/3mem.prof
