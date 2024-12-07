#!/usr/bin/env bash

INPUT="input"
PART1=0

function getBinaryList(){
    # This is stupid, but my brain cant figure out how to do this better atm.
    case $1 in
        1)  D2B=({0..1});;
        2)  D2B=({0..1}{0..1});;
        3)  D2B=({0..1}{0..1}{0..1});;
        4)  D2B=({0..1}{0..1}{0..1}{0..1});;
        5)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1});;
        6)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
        7)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
        8)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
        9)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
       10)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
       11)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
       12)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
       13)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
       14)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
       15)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
       16)  D2B=({0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1}{0..1});;
    esac
}

while IFS=":" read PRODUCT MATH;
do
    echo "# $PRODUCT - $MATH";
    NUMBERS=$(echo "$MATH" | wc -w)
    OPERATORS=$((NUMBERS-1))
    getBinaryList $OPERATORS

    SUCCESS=false
    for BIN in ${D2B[@]};
    do
        BUCKET=$(echo "$BIN" | sed -e 's/1/\*/g' -e 's/0/\+/g')
        i=0
        for NUMBER in $MATH;
        do
            [[ $i == 0 ]] && SUM=$NUMBER || SUM=$((SUM ${BUCKET:$((i-1)):1} NUMBER))
            i=$((i+1))
        done
        [[ $SUM == $PRODUCT ]] && SUCCESS=true
    done

    [[ $SUCCESS == true ]] && PART1=$((PART1+PRODUCT))

done <$INPUT

echo "Part1: $PART1"