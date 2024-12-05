#!/usr/bin/env bash

INPUT='input'

PART1=0

# Setup rules array
RULE_DATA=$(grep '|' $INPUT)
declare -A RULES
i=0
for RULE in $(echo "$RULE_DATA");
do
    RULES[$i,0]=$(echo "$RULE" | awk -F '|' '{print $1}')
    RULES[$i,1]=$(echo "$RULE" | awk -F '|' '{print $2}')
    i=$((i+1))
done
RULES_COUNT=$i


grep ',' $INPUT | while read LINE;
do
    ERRORS=0
    for (( i=0; i<RULES_COUNT; i++));
    do
        SPLIT=${RULES[$i,0]}
        NEEDLE=${RULES[$i,1]}

        # Check if both SPLIT and NEEDLE are in the "haystack"
        [[ $LINE == *"$SPLIT"* && $LINE == *"$NEEDLE"* ]] || continue

        # Check if NEEDLE comes after SPLIT, add error if not
        REGEX="(.*)$SPLIT(.*)"
        [[ $LINE =~ $REGEX ]] || continue
        POST=${BASH_REMATCH[2]}
        [[ $POST == *"$NEEDLE"* ]] || ERRORS=$((ERRORS+1))
    done
    
    if [[ $ERRORS == 0 ]];
    then
        NUMBERS=$(echo "$LINE" | wc -c)
        MID=$(( (NUMBERS/2) - 1))
        MIDDLE="${LINE:$MID:2}"
        PART1=$(echo $MIDDLE + $PART1 | bc)
        echo $PART1
    fi
done