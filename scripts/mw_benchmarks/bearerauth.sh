#!/bin/bash

cd .. || exit 1
cd .. || exit 1
cd benchmarks || exit 1
cd mw_benchmarks || exit 1

go test -bench=BenchmarkBearerAuthMiddleware -count=20 -benchmem -cpuprofile=profiles/bearerAuthCpu.prof -memprofile=profiles/bearerAuthMem.prof > ../mw_results/bearerAuthResults.txt
benchstat ../mw_results/bearerAuthResults.txt

go tool pprof profiles/bearerAuthMem.prof
