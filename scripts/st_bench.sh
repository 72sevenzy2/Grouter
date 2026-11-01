#!/bin/bash

cd .. || exit 1
cd benchmarks || exit 1
go test -bench="^(BenchmarkGrouterStaticRoutes|BenchmarkGolangStdlibStaticRoutes)$" -run=^$ -benchmem -cpuprofile=profiles/StaticRoutesCpu.prof -memprofile=profiles/StaticRoutesMem.prof -count=10 > results/staticBenchmarks.txt
benchstat results/staticBenchmarks.txt

go tool pprof profiles/StaticRoutesMem.prof
