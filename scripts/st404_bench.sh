#!/bin/bash

cd .. || exit 1
cd benchmarks || exit 1
go test -bench="^(BenchmarkGrouterStaticRoutes404|BenchmarkGolangStdlibStaticRoutes404)$" -run=^$ -benchmem -cpuprofile=4cpu.prof -memprofile=4mem.prof -count=10 > results/staticBenchmarks404.txt
benchstat staticBenchmarks404.txt

go tool pprof profiles/4mem.prof
