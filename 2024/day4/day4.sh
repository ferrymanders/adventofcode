#!/bin/bash

INPUT="input"

IFS=$'\r\n' GLOBIGNORE='*' command eval 'LINES=($(cat $INPUT))'
LINE_COUNT=$(cat $INPUT | wc -l)

function findNext() {
    local X=$1
    local Y=$2
    local PREV_CHAR=$3
    local FOLLOW="$4"

    local SEARCHX
    local SEARCHY

    [ -z $FOLLOW ] && FOLLOW="TOPLEFT TOPMID TOPRIGHT LEFT RIGHT BOTTOMLEFT BOTTOMMID BOTTOMRIGHT"

    for DIRECTION in $FOLLOW
    do
        local CHAR=''
        case $DIRECTION in
            'TOPLEFT')      SEARCHX=$((X-1)); SEARCHY=$((Y-1));;
            'TOPMID')       SEARCHX=$((X-1)); SEARCHY=$((Y));;
            'TOPRIGHT')     SEARCHX=$((X-1)); SEARCHY=$((Y+1));;
            'LEFT')         SEARCHX=$((X));   SEARCHY=$((Y-1));;
            'RIGHT')        SEARCHX=$((X));   SEARCHY=$((Y+1));;
            'BOTTOMLEFT')   SEARCHX=$((X+1)); SEARCHY=$((Y-1));;
            'BOTTOMMID')    SEARCHX=$((X+1)); SEARCHY=$((Y));;
            'BOTTOMRIGHT')  SEARCHX=$((X+1)); SEARCHY=$((Y+1));;
        esac
        CHAR=$(getChar $SEARCHX $SEARCHY)
        [[ "$PREV_CHAR" == "X" && "$CHAR" == 'M' ]] && findNext $SEARCHX $SEARCHY 'M' $DIRECTION
        [[ "$PREV_CHAR" == "M" && "$CHAR" == 'A' ]] && findNext $SEARCHX $SEARCHY 'A' $DIRECTION
        [[ "$PREV_CHAR" == "A" && "$CHAR" == 'S' ]] && PART1=$((PART1+1))
        
    done
}

function getChar() {
    X=$1; Y=$2

    [[ $X -lt 0 || $X -gt $LINE_COUNT ]] && return
    LINE=${LINES[$X]}

    [[ $Y -lt 0 || $Y -gt ${#LINE} ]] && return

    CHAR=${LINE:$Y:1}

    echo "$CHAR"
}


PART1=0

for (( x=0; x<$LINE_COUNT; x++ ));
do
    LINE=${LINES[$x]}
    for (( y=0; y<${#LINE}; y++))
    do
        CHAR=${LINE:$y:1}
        [ $CHAR == 'X' ] && findNext $x $y 'X' ''
    done
done

echo "# Part1: $PART1"