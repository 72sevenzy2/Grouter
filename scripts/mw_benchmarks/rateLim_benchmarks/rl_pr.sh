#!/bin/bash

# navigating to root from current dir
cd .. || exit 1
cd .. || exit 1
cd .. || exit 1

# navigating to needed benchmarking dir.
cd benchmarks || exit 1
cd mw_benchmarks || exit 1
cd rateLimit_benchmarks || exit 1

go test -bench=BenchmarkRateLimit_parallelReqs -count=20 -benchmem -cpuprofile=profiles/parallelReqsCpu.prof -memprofile=profiles/parallelReqsMem.prof > ../../mw_results/rateLimiter_results/parallelReqsResults.txt
benchstat ../../mw_results/rateLimiter_results/parallelReqsResults.txt

go tool pprof profiles/parallelReqsMem.prof
