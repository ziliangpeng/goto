#!/bin/bash

# may be a good idea to refer to this file as correct answer to check against
# https://raw.githubusercontent.com/lucky-bai/projecteuler-solutions/refs/heads/master/Solutions.md

get_answer() {
    PE_FNAME="/tmp/project_euler.ans"
    if [ ! -f "$PE_FNAME" ]; then
        curl -s https://raw.githubusercontent.com/lucky-bai/projecteuler-solutions/refs/heads/master/Solutions.md > $PE_FNAME
    fi
    cat $PE_FNAME | grep "^$1\." | cut -d ' ' -f 2 | tr -cd '[:print:]'
}

check() {
    if [ "$1" == "$2" ]; then
        printf "\033[32mCorrect\033[0m\n"
    else
        printf "\033[31mIncorrect\033[0m\n"
    fi
}

for file in *.go; do
    if [[ $file =~ ^([0-9]+)\.go$ ]]; then
        number=${BASH_REMATCH[1]}
        echo "Problem $number:"
    else continue
    fi
    # find a way to time it
    result=$(go run $number.go prime.go | tail -n 1)
    expected=$(get_answer $number)
    check $expected $result
done
