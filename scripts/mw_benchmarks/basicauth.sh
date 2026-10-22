#!/bin/bash

cd .. || exit 1
cd .. || exit 1
cd benchmarks || exit 1
cd mw_benchmarks || exit 1

go test -bench=BenchmarkBasicAuthMiddleware -count=20 -benchmem -cpuprofile=profiles/basicAuthCpu.prof -memprofile=profiles/basicAuthMem.prof > ../mw_results/basicAuthResults.txt
benchstat ../mw_results/basicAuthResults.txt

go tool pprof profiles/basicAuthMem.prof
