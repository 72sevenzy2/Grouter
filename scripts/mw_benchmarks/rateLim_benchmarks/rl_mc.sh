#!/bin/bash

# navigating to root from current dir
cd .. || exit 1
cd .. || exit 1
cd .. || exit 1

# navigating to needed benchmarking dir.
cd benchmarks || exit 1
cd mw_benchmarks || exit 1
cd rateLimit_benchmarks || exit 1

go test -bench=BenchmarkRateLimit_multiClients -count=20 -benchmem -cpuprofile=profiles/multiClientsCpu.prof -memprofile=profiles/multiClientsMem.prof > ../../mw_results/rateLimiter_results/multiClientResults.txt
benchstat ../../mw_results/rateLimiter_results/multiClientResults.txt

go tool pprof profiles/multiClientsMem.prof
