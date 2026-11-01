#!/bin/bash

cd .. || exit 1
cd benchmarks || exit 1
go test -bench="^(BenchmarkGrouterDynamicRoutes|BenchmarkGolangStdlibRouterDynamicRoutes)$" -run=^$ -benchmem -cpuprofile=profiles/DynamicRoutesCpu.prof -memprofile=profiles/DynamicRoutesMem.prof -count=10 > results/dynamicBenchmarks.txt
benchstat results/dynamicBenchmarks.txt

go tool pprof profiles/DynamicRoutesMem.prof
