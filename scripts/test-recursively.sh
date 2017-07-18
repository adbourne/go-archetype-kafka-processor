#!/usr/bin/env bash
##
# Runs test with coverage recursively
##
set -e
importCommonScripts() {
    DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
    source "$DIR/common.sh"
}

runTestsRecursively() {
    echo "" > coverage.txt

    for d in $(go list ./... | grep -v vendor); do
        go test -v -race -coverprofile=profile.out -covermode=atomic $d \
        | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' \
        | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''

        if [ -f profile.out ]; then
            cat profile.out >> coverage.txt
            rm profile.out
        fi
    done
}

##
# main
##
importCommonScripts

printBold
echo "Running tests..."
printNormal

runTestsRecursively
