#!/bin/bash

INPUT="test_input"

PART1=0
while read LINE;
do
    PREV_LEVEL=start
    PREV_TREND=unknown
    ERRORS=0
    for LEVEL in $(echo $LINE);
    do
        if [ $PREV_LEVEL != "start" ];
        then
            DIFFERENCE=$(( $LEVEL - $PREV_LEVEL ))
            [[ $DIFFERENCE -gt 3 || $DIFFERENCE -lt -3 || $DIFFERENCE -eq 0 ]] && ERRORS=$((ERRORS+1))
            [[ $DIFFERENCE -gt 0 ]] && TREND=positive
            [[ $DIFFERENCE -lt 0 ]] && TREND=negative
            [[ $PREV_TREND != $TREND && $PREV_TREND != "unknown" ]] && ERRORS=$((ERRORS+1))
            PREV_TREND=$TREND
        fi
        PREV_LEVEL=$LEVEL
    done
    [[ $ERRORS -eq 0 ]] && PART1=$((PART1 + 1))
done <<< `cat $INPUT`

echo "# Part1 : $PART1"