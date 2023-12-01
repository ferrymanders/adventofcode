#!/bin/bash

# Part 1
sum=""
while read line;
do
    firstNumber=$(echo $line | grep -Eo '^[a-z]*[0-9]{1}')
    lastNumber=$(echo $line | grep -Eo '[0-9]{1}[a-z]*$')
    code="${firstNumber: -1}${lastNumber:0:1}"
    sum="$code+$sum"
done<<<`cat input`

part1=$(echo $sum | sed 's/.$//' | bc)

# Part 2
sum=""
while read line;
do
    firstNumber=$(echo $line | pcregrep -o1 '(one|two|three|four|five|six|seven|eight|nine|[0-9]).*' | head -1)
    lastNumber=$(echo $line | pcregrep -o1 '.*(one|two|three|four|five|six|seven|eight|nine|[0-9])' | tail -1)
    code=$(echo "${firstNumber}${lastNumber}" | sed -e 's/one/1/g' \
                                                    -e 's/two/2/g' \
                                                    -e 's/three/3/g' \
                                                    -e 's/four/4/g' \
                                                    -e 's/five/5/g' \
                                                    -e 's/six/6/g' \
                                                    -e 's/seven/7/g' \
                                                    -e 's/eight/8/g' \
                                                    -e 's/nine/9/g')
    sum="$code+$sum"
done<<<`cat input`

part2=$(echo $sum | sed 's/.$//' | bc)

echo "Part1 = $part1"
echo "Part2 = $part2"