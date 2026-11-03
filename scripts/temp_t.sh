#!/bin/bash

cd .. || exit 1

go test templating_test.go -v
