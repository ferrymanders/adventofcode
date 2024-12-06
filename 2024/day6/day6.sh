#!/usr/bin/env bash

INPUT="input"

declare -A MAP

PART1=0

X=0
while read LINE;
do
    for (( Y=0; Y<${#LINE}; Y++ ));
    do
        MAP[$X,$Y]=${LINE:$Y:1}
        [[ ${LINE:$Y:1} == "^" ]] && CURRENT_POS="$X $Y"
    done
    X=$((X+1))
done<$INPUT

MAX_X=$X
MAX_Y=$Y

echo "# Guard start pos: $CURRENT_POS"
CURRENT_DIRECTION="UP"

CONTINUE=true
while [[ $CONTINUE == true ]];
do
    read X Y <<<`echo $CURRENT_POS`

    MAP[$X,$Y]="X"

    [[ $X == 0 || $X == $MAX_X || $Y == 0 || $Y == $MAX_Y ]] && CONTINUE=false

    case $CURRENT_DIRECTION in
        "UP")
            X=$((X-1)); 
            [[ ${MAP[$X,$Y]} == "#" ]] && CURRENT_DIRECTION="RIGHT" && continue;
            MAP[$X,$Y]="^"
            ;;
        "RIGHT")
            Y=$((Y+1)); 
            [[ ${MAP[$X,$Y]} == "#" ]] && CURRENT_DIRECTION="DOWN" && continue;            
            MAP[$X,$Y]=">";;
        "DOWN")
            X=$((X+1)); 
            [[ ${MAP[$X,$Y]} == "#" ]] && CURRENT_DIRECTION="LEFT" && continue;            
            MAP[$X,$Y]="V";;
        "LEFT")
            Y=$((Y-1)); 
             [[ ${MAP[$X,$Y]} == "#" ]] && CURRENT_DIRECTION="UP" && continue;            
           MAP[$X,$Y]="<";;
    esac
    
    CURRENT_POS="$X $Y"
done

for (( X=0; X<$MAX_X; X++))
do
    for (( Y=0; Y<$MAX_Y; Y++))
    do
        echo -n "${MAP[$X,$Y]}"
        [[ ${MAP[$X,$Y]} == "X" ]] && PART1=$((PART1+1))
    done
    echo ""
done


echo "# Part1: $PART1"