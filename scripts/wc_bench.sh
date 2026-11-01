#!/bin/bash

cd .. || exit 1
cd benchmarks || exit 1

go test -bench="^(BenchmarkGrouterWildcardRoute|BenchmarkGolangStdlibRouterWildCardRoute)$" -run=^$ -benchmem -cpuprofile=profiles/WildcardRoutesCpu.prof -memprofile=profiles/WildcardRoutesMem.prof -count=10 > results/wildcardRoutesResults.txt
benchstat results/wildcardRoutesResults.txt

go tool pprof profiles/WildcardRoutesMem.prof
