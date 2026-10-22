#!/bin/bash

cd .. || exit 1
cd .. || exit 1
cd benchmarks || exit 1
cd mw_benchmarks || exit 1

go test -bench=BenchmarkLoggerMiddleware -count=20 -benchmem -cpuprofile=profiles/loggerCpu.prof -memprofile=profiles/loggerMem.prof > ../mw_results/loggerResults.txt
benchstat ../mw_results/loggerResults.txt

go tool pprof profiles/loggerMem.prof
