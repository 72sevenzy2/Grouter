#!/bin/bash

cd .. || exit 1
cd benchmarks || exit 1
go test -bench="^(BenchmarkGrouterStaticRoutes|BenchmarkGolangStdlibStaticRoutes)$" -run=^$ -benchmem -cpuprofile=profiles/2cpu.prof -memprofile=profiles/2mem.prof -count=10 > results/staticBenchmarks.txt
benchstat staticBenchmarks.txt

go tool pprof profiles/2mem.prof
