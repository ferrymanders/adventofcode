#!/bin/bash

INPUT=input
PART1=0
PART2=0

function getBiggest()
{
    BANK=$1
    OFFSET=$2
    START=$3

    BIGGEST=0
    for (( i = $START; i < (${#BANK}-$OFFSET); i++ ))
    do
        BATTERY=${BANK:$i:1}
        if [ $BATTERY -gt $BIGGEST ];
        then
            BIGGEST=$BATTERY
            BATTERY_NR=$i
        fi
    done
    echo "${BIGGEST};${BATTERY_NR}"
}

while read BANK;
do
    # Part1
    JOLTAGE_PART1=()
    NEXT_START=0
    for (( x = 1; x >= 0; x-- ))
    {
        OUTPUT=$(getBiggest $BANK $x $NEXT_START)
        IFS=';' read BIGGEST_BATTERY BATTERY_NR <<< `echo $OUTPUT`
        NEXT_START=$(( BATTERY_NR + 1 ))
        JOLTAGE_PART1+=$BIGGEST_BATTERY
    }
    PART1=$(( PART1 + JOLTAGE_PART1 ))    

    # Part2
    JOLTAGE_PART2=()
    NEXT_START=0
    for (( x = 11; x >= 0; x-- ))
    {
        OUTPUT=$(getBiggest $BANK $x $NEXT_START)
        IFS=';' read BIGGEST_BATTERY BATTERY_NR <<< `echo $OUTPUT`
        NEXT_START=$(( BATTERY_NR + 1 ))
        JOLTAGE_PART2+=$BIGGEST_BATTERY
    }
    PART2=$(( PART2 + JOLTAGE_PART2 ))
done < $INPUT

echo "Part 1: $PART1"
echo "Part 2: $PART2"