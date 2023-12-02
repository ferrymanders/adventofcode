#!/bin/bash

inputFile="input"

part1=0
part2=0

while read line;
do
    gameId=$(echo $line | awk -F ':' '{print $1}' | grep -Eo '[0-9]+')

    reds=$(echo $line | grep -Eo '[0-9]+ red' | grep -Eo '[0-9]+' | sort -h | tail -1)
    greens=$(echo $line | grep -Eo '[0-9]+ green' | grep -Eo '[0-9]+' | sort -h | tail -1)
    blues=$(echo $line | grep -Eo '[0-9]+ blue' | grep -Eo '[0-9]+' | sort -h | tail -1)

    # Part1
    possible=true
    [ $reds -gt 12 ]   && possible=false
    [ $greens -gt 13 ] && possible=false
    [ $blues -gt 14 ]  && possible=false

    [ "$possible" == "true" ] && part1=$(( part1 + gameId ))

    # Part2
    power=$(( reds * greens * blues ))
    part2=$(( part2 + power ))

done<<<`cat $inputFile`

echo "Part1: $part1"
echo "Part2: $part2"