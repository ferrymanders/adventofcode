#!/bin/bash

INPUT="input"

PART1=0

REGEX="mul\([0-9]+,[0-9]+\)"


for MUL in $(grep -Eo "$REGEX" $INPUT);
do
    # Cleanup
    MUL=$(echo $MUL | sed -e 's/mul(//' -e 's/)//')
    IFS=',' read FIRST SECOND <<<`echo $MUL`

    PART1=$(( PART1 + $((FIRST * SECOND)) ))
done

echo "# Part1: $PART1"