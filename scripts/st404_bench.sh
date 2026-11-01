#!/bin/bash

cd .. || exit 1
cd benchmarks || exit 1
go test -bench="^(BenchmarkGrouterStaticRoutes404|BenchmarkGolangStdlibStaticRoutes404)$" -run=^$ -benchmem -cpuprofile=profiles/StaticRoutes404Cpu.prof -memprofile=profiles/StaticRoutes404Mem.prof -count=10 > results/staticBenchmarks404.txt
benchstat results/staticBenchmarks404.txt

go tool pprof profiles/StaticRoutes404Mem.prof
