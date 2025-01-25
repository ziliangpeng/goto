#!/bin/bash

assert() {
    if [ "$1" == "$2" ]; then
        echo "Correct"
    else
        echo "Incorrect"
    fi
}

# find a way to time it
result=$(go run 69.go prime.go | tail -n 1)
expected=$(cat 69.ans)
assert $expected $result
