#!/bin/bash

cd .. || exit 1
cd .. || exit 1
cd benchmarks || exit 1
cd mw_benchmarks || exit 1

go test -bench="^(BenchmarkRecoveryMiddleware)$" -run=^$ -count=20 -benchmem -cpuprofile=profiles/recoveryCpu.prof -memprofile=profiles/recoveryMem.prof > ../mw_results/recovererResults.txt
benchstat ../mw_results/recovererResults.txt

go tool pprof profiles/recoveryMem.prof
