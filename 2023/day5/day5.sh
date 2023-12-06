#!/bin/bash

inputFile="input"
debug="true"

highestMapNumber=$(cat $inputFile | grep -Eo '[0-9]+' | sort -h | tail -1)

seedList=$(cat $inputFile | grep -E '^seeds:' | awk -F ':' '{print $2}')

# Create data save location
[ -d "./tmp" ] || mkdir ./tmp

while read line
do
    [ -z "$line" ] && continue
    [[ "$line" =~ "seeds:" ]] && continue

    if [[ "$line" =~ "-to-" ]]
    then
        mapName=$(echo $line | awk '{print $1}')
        src=$(echo $mapName | awk -F '-to-' '{print $1}')
        dest=$(echo $mapName | awk -F '-to-' '{print $2}')
        
        echo -n '' > tmp/${mapName}.map
    else
        echo "$line" >> tmp/${mapName}.map
    fi

done<<<`cat $inputFile`


part1Loc=$highestMapNumber
part1Seed=0

for seed in $(echo $seedList);
do 
    prev="seed"
    seedInfo="$seed"
    prevData="$seed"
    for type in soil fertilizer water light temperature humidity location;
    do
        data=""
        while read line
        do
            startDest=$(echo $line | awk '{print $1}')
            startSrc=$(echo $line | awk '{print $2}')
            range=$(echo $line | awk '{print $3}')

            [ "$prevData" -lt "$startSrc" ] && continue
            [ "$prevData" -gt "$((startSrc + range))" ] && continue
            distance=$(( prevData - startSrc ))
            data=$(( startDest + distance ))
        done<<<`cat tmp/${prev}-to-${type}.map`

        [ -z "$data" ] && data=$prevData

        seedInfo="${seedInfo};$data"
        
        prev=$type
        prevData=$data
    done

    seedNr=$(echo $seedInfo | awk -F ';' '{print $1}')
    seedLoc=$(echo $seedInfo | awk -F ';' '{print $8}')
    [ $seedLoc -lt $part1Loc ] && part1Loc=$seedLoc && part1Seed=$seedNr && echo "#closer seed found : $seedNr - $seedLoc"
    [ "$debug" == "true" ] && echo $seedInfo
done

echo "Part1 - Seed: $part1Seed - Loc: $part1Loc"