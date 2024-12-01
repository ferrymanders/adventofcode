#!/bin/bash

INPUT="input"

LEFT_LIST=$(cat $INPUT | awk '{print $1}')
RIGHT_LIST=$(cat $INPUT | awk '{print $2}')
LINES=$(echo $LEFT_LIST | wc -w)

PART1=0
PART2=0
for i in $(seq 1 $LINES);
do
    LEFT_NR=$(echo "$LEFT_LIST" | sort -h | head -n ${i} | tail -1)
    RIGHT_NR=$(echo "$RIGHT_LIST" | sort -h | head -n ${i} | tail -1)
    DIFFERENCE=$(( LEFT_NR - RIGHT_NR ))
    [ $DIFFERENCE -lt 0 ] && DISTANCE=$(( DIFFERENCE * -1 )) || DISTANCE=$DIFFERENCE

    PART1=$(( PART1 + $DISTANCE ))

    PART2_LEFT_NR=$(echo "$LEFT_LIST" | head -n ${i} | tail -1)
    PART2_COUNTER=$(echo "$RIGHT_LIST" | grep -E "^$PART2_LEFT_NR$" | wc -l)
    PART2=$(( PART2 + ( PART2_LEFT_NR * PART2_COUNTER ) ))
done

echo "PART1 = $PART1"
echo "PART2 = $PART2"