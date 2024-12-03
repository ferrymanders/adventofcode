#!/bin/bash

INPUT="input"

PART1=0
PART2=0
ENABLED=true

REGEX="(do(|n't)\(\)|mul\([0-9]+,[0-9]+\))"


for COMMAND in $(grep -Eo "$REGEX" $INPUT);
do
    [[ $COMMAND == "do()" ]] && ENABLED=true
    [[ $COMMAND == "don't()" ]] && ENABLED=false

    CALC=$(echo $COMMAND | sed -e 's/mul(//' -e 's/)//')
    IFS=',' read FIRST SECOND <<<`echo $CALC`

    if [[ $COMMAND == mul* && $ENABLED == true ]];
    then 
        PART2=$(( PART2 + $((FIRST * SECOND)) ))
    fi

    [[ $COMMAND == mul* ]] && PART1=$(( PART1 + $((FIRST * SECOND)) ))
done

echo "# Part1: $PART1"
echo "# Part2: $PART2"