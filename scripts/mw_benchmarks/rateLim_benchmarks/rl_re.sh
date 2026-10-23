#!/bin/bash

# navigating to root from current dir
cd .. || exit 1
cd .. || exit 1
cd .. || exit 1

# navigating to needed benchmarking dir.
cd benchmarks || exit 1
cd mw_benchmarks || exit 1
cd rateLimit_benchmarks || exit 1

go test -bench=BenchmarkRateLimit_ReqsExceeded -count=20 -benchtime=1x -benchmem -cpuprofile=profiles/reqsExceedCpu.prof -memprofile=profiles/reqsExceedMem.prof > ../../mw_results/rateLimiter_results/reqsExceedResults.txt
benchstat ../../mw_results/rateLimiter_results/reqsExceedResults.txt

go tool pprof profiles/reqsExceedMem.prof
