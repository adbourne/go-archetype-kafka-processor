#!/usr/bin/env bash

# Text Weight
BOLD_WEIGHT=$(tput bold)
NORMAL_WEIGHT=$(tput sgr0)

# Colours
RED='\033[0;31m'
YELLOW='\033[0;33m'
GREEN='\033[0;32m'
NC='\033[0m'

printBold() {
    echo -n "$BOLD_WEIGHT"
}

printNormal() {
    echo -n "$NORMAL_WEIGHT"
}

printRed() {
    echo -ne "$RED"
}

printYellow() {
    echo -ne "$YELLOW"
}

printGreen() {
    echo -ne "$GREEN"
}

printDefaultColour() {
    echo -ne "$NC"
}