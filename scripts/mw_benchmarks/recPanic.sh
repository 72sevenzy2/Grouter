#!/bin/bash

cd .. || exit 1
cd .. || exit 1
cd benchmarks || exit 1
cd mw_benchmarks || exit 1

go test -bench="^(BenchmarkRecoveryMiddlewarePanic)$" -run=^$ -count=20 -benchmem -cpuprofile=profiles/recoveryPanicCpu.prof -memprofile=profiles/recoveryPanicMem.prof > ../mw_results/recovererPanicResults.txt
benchstat ../mw_results/recovererPanicResults.txt

go tool pprof profiles/recoveryPanicMem.prof
