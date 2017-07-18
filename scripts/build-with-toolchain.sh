#!/usr/bin/env bash

importCommonScripts() {
    DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
    source "$DIR/common.sh"
}

checkForGoPath() {
    if [ -z "$GOPATH" ]; then
        printRed
        echo "Please set GOPATH to continue"
        printDefaultColour
        exit 1
    fi
}

getGoCyclo() {
   if [ ! -d "$GOPATH/src/github.com/fzipp/gocyclo" ]; then
       echo "Getting gocyclo..."
       local gocycloPath="github.com/fzipp/gocyclo"
       go get "$gocycloPath"
       go install "$gocycloPath"
   fi
}

getGoLint() {
    if [ ! -d "$GOPATH/src/github.com/golang/lint" ]; then
        echo "Getting golint..."
        local golintPath="github.com/golang/lint"
        go get "$golintPath"
        go install "$golintPath"
    fi
}

getIneffassign() {
    if [ ! -d "$GOPATH/src/github.com/gordonklaus/ineffassign" ]; then
        echo "Getting ineffassign..."
        local ineffassignPath="github.com/gordonklaus/ineffassign"
        go get "$ineffassignPath"
        go install "$ineffassignPath"
    fi
}

getMisspell() {
    if [ ! -d "$GOPATH/src/github.com/client9/misspell/" ]; then
        echo "Getting misspell..."
        local misspellPath="github.com/client9/misspell/cmd/misspell"
        go get "$misspellPath"
        go install "$misspellPath"
    fi
}

toolTextBreak() {
    printBold
    echo "$1:"
    printNormal
}
buildProject() {
    printBold
    echo "Building project..."
    printNormal

    gofmt -s -w .

    printYellow
    go install
    if [ $? != 0 ]; then
        printRed
        echo  "Build failed"
        printDefaultColour
        exit $?
    fi

    toolTextBreak "govet"
    printYellow
    go vet
    printDefaultColour

    toolTextBreak "gocyclo"
    printYellow
    for d in $(go list ./... | grep -v vendor); do
        gocyclo -over 25 "$GOPATH/src/$d"
    done
    printDefaultColour

    toolTextBreak "golint"
    printYellow
    golint .
    printDefaultColour

    toolTextBreak "ineffassign"
    printYellow
    ineffassign .
    printDefaultColour

    toolTextBreak "misspell"
    printYellow
    for d in $(go list ./... | grep -v vendor); do
        misspell "$GOPATH/src/$d"
    done
    printDefaultColour

    printGreen
    echo "Build Complete"
    printDefaultColour
}

##
# main
##
importCommonScripts
checkForGoPath
getGoCyclo
getGoLint
getIneffassign
getMisspell
buildProject