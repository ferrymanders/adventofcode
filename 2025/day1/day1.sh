#!/bin/bash

INPUT=input

POSITION=50
COUNT_PART1=0
COUNT_PART2=0

while read LINE;
do
    DIRECTION=$(echo $LINE| grep -Eo '^(L|R)')
    CLICKS=$(echo $LINE| grep -Eo '[0-9]+$')

    case $DIRECTION in
        L) OPERATOR='-';;
        R) OPERATOR='+';;
    esac

    while [ $CLICKS -gt 0 ];
    do
        POSITION=$(( POSITION $OPERATOR 1 ))
        CLICKS=$(( CLICKS - 1 ))
        case $POSITION in
            -1) POSITION=99;;
            100) POSITION=0;;
        esac
        [ $POSITION -eq 0 ] && COUNT_PART2=$(( COUNT_PART2 + 1 ))
    done

    [ $POSITION -eq 0 ] && COUNT_PART1=$(( COUNT_PART1 + 1 ))
done < $INPUT

echo "# Part 1: $COUNT_PART1"
echo "# Part 2: $COUNT_PART2"