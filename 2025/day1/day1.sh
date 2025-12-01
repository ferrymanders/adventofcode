#!/bin/bash

INPUT=input

POSITION=50
COUNT=0

while read LINE;
do
    DIRECTION=$(echo $LINE| grep -Eo '^(L|R)')
    CLICKS=$(echo $LINE| grep -Eo '[0-9]+$')

    while [ $CLICKS -gt 100 ];
    do
        CLICKS=$(( CLICKS - 100 ))
    done

    case $DIRECTION in
        L) POSITION=$(( POSITION - CLICKS ));;
        R) POSITION=$(( POSITION + CLICKS ));;
    esac

    [ $POSITION -lt 0 ] && POSITION=$(( 100 + POSITION ))
    [ $POSITION -gt 99 ] && POSITION=$(( 0 + (POSITION-100) ))

    [ $POSITION -eq 0 ] && COUNT=$(( COUNT + 1 ))

    echo "# $DIRECTION - $CLICKS - $POSITION"

done < $INPUT 
#| column -t | grep -E ' 0$' | wc -l

echo -e "\n\n# Part 1 = $COUNT"