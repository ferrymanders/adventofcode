#!/bin/bash

INPUT=input

PART1=0
PART2=0

for RANGE in $(cat $INPUT | grep -Eo '[0-9]+\-[0-9]+');
do
    IFS='-' read START FINISH <<< `echo $RANGE`

    for (( NUMBER = $START; NUMBER <= $FINISH; NUMBER++ ));
    do
        STRLEN=${#NUMBER}
        HALV=$(( STRLEN / 2 ))

        # Part1
        FIRST=${NUMBER::$HALV}
        LAST=${NUMBER:$HALV}
        if [ "$FIRST" == "$LAST" ];
        then 
            PART1=$(( PART1 + $NUMBER ))
        fi

        # Part2
        for (( STEP = 1; STEP <= $HALV; STEP++ ));
        do
            PREV_CHECK=""
            MATCH=true
            for (( START = 0; START < $STRLEN; START = START + STEP ));
            do
                CHECK=${NUMBER:$START:$STEP}

                if [[ "$PREV_CHECK" != "" ]];
                then
                    if [[ $PREV_CHECK != $CHECK ]];
                    then
                        MATCH=false
                        break # Stops the "for START"-loop, this number didn't match
                    fi  
                fi
                PREV_CHECK=$CHECK
            done

            if [ $MATCH == true ];
            then 
                PART2=$(( PART2 + $NUMBER ))
                break # Stops the "for STEP"-loop, we've found a pattern match
            fi
        done
    done
done

echo "# Part1: $PART1"
echo "# Part2: $PART2"