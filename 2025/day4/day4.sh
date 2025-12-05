#!/bin/bash

INPUT=input
PART1=0
PART2=0

# Lets create a map
declare -A MAP
ROW=0
while read LINE;
do
    LINELEN=${#LINE}
    for (( COL = 0; COL < $LINELEN; COL++ ));
    do
        MAP[$ROW,$COL]=${LINE:$COL:1}
    done
    ROW=$(( ROW + 1 ))
done < $INPUT
ROWS=$ROW
COLS=$COL

# Part1
for (( X = 0; X < $ROWS; X++ ))
do
    for (( Y = 0; Y < $COLS; Y++ ))
    do
        if [ "${MAP[$X,$Y]}" == "@" ];
        then
            COUNTER=-1
            for XX in {-1..1};
            do
                for YY in {-1..1};
                do
                    if [ "${MAP[$(( X + XX )),$(( Y + YY ))]}" == "@" ];
                    then
                        COUNTER=$((COUNTER + 1))
                    fi
                done
            done
            [ $COUNTER -lt 4 ] && PART1=$(( PART1 + 1 ))
        fi
    done
done

# Part2
REMOVED_PAPER=666
while [ $REMOVED_PAPER -gt 0 ];
do
    REMOVED_PAPER=0
    for (( X = 0; X < $ROWS; X++ ))
    do
        for (( Y = 0; Y < $COLS; Y++ ))
        do
            if [ "${MAP[$X,$Y]}" == "@" ];
            then
                COUNTER=-1
                for XX in {-1..1};
                do
                    for YY in {-1..1};
                    do
                        if [ "${MAP[$(( X + XX )),$(( Y + YY ))]}" == "@" ];
                        then
                            COUNTER=$((COUNTER + 1))
                        fi
                    done
                done

                if [ $COUNTER -lt 4 ];
                then
                    REMOVED_PAPER=$(( REMOVED_PAPER + 1 ))
                    PART2=$(( PART2 + 1 ))
                    MAP[$X,$Y]="X"
                fi
            fi
        done
    done
done

echo "Part 1: $PART1"
echo "Part 2: $PART2"