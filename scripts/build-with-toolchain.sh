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

runGoCyclo() {
    printYellow
    for d in $(go list ./... | grep -v vendor); do
        gocyclo -over 25 "$GOPATH/src/$d" | grep -v "/vendor/"
    done
    printDefaultColour

}

getGoLint() {
    if [ ! -d "$GOPATH/src/github.com/golang/lint" ]; then
        echo "Getting golint..."
        local golintPath="github.com/golang/lint/golint"
        go get "$golintPath"
        go install "$golintPath"
    fi
}

runGoLint() {
    printYellow
    for d in $(go list ./... | grep -v vendor); do
        golint "$d" | grep -v "/vendor/"
    done
    printDefaultColour
}

getIneffassign() {
    if [ ! -d "$GOPATH/src/github.com/gordonklaus/ineffassign" ]; then
        echo "Getting ineffassign..."
        local ineffassignPath="github.com/gordonklaus/ineffassign"
        go get "$ineffassignPath"
        go install "$ineffassignPath"
    fi
}

runIneffassign() {
    printYellow
    for d in $(go list ./... | grep -v vendor); do
        ineffassign "$GOPATH/src/$d" | grep -v "/vendor/"
    done
    printDefaultColour
}

getMisspell() {
    if [ ! -d "$GOPATH/src/github.com/client9/misspell/" ]; then
        echo "Getting misspell..."
        local misspellPath="github.com/client9/misspell/cmd/misspell"
        go get "$misspellPath"
        go install "$misspellPath"
    fi
}

runMisspell() {
    printYellow
    for d in $(go list ./... | grep -v vendor); do
        misspell "$d" | grep -v "/vendor/"
    done
    printDefaultColour
}

runGoFmt() {
    gofmt -s -w .
}

runGoInstall() {
    printYellow
    go install
    if [ $? != 0 ]; then
        printRed
        echo  "Build failed"
        printDefaultColour
        exit $?
    fi
    printDefaultColour
}

runGoVet() {
    printYellow
    go vet
    printDefaultColour
}

printToolHeader() {
    printBold
    echo "$1"
    printNormal
}

buildProject() {
    printBold
    echo "Building project..."
    printNormal

    runGoFmt

    runGoInstall

    printToolHeader "govet"
    runGoVet

    printToolHeader "gocyclo"
    runGoCyclo

    printToolHeader "golint"
    runGoLint

    printToolHeader "ineffassign"
    runIneffassign

    printToolHeader "misspell"
    runMisspell

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