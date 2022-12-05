#!/bin/bash

BASEDIR=${PWD}
WORKDIR=${PWD}/work

[ -d ${WORKDIR} ] || mkdir ${WORKDIR}

createStart() {
    echo "RSLFQ"    > ${WORKDIR}/stack1
    echo "NZQGPT"   > ${WORKDIR}/stack2
    echo "SMQB"     > ${WORKDIR}/stack3
    echo "TGZJHCBQ" > ${WORKDIR}/stack4
    echo "PHMBNFS"  > ${WORKDIR}/stack5
    echo "PCQNSLVG" > ${WORKDIR}/stack6
    echo "WCF"      > ${WORKDIR}/stack7
    echo "QHGZWVPM" > ${WORKDIR}/stack8
    echo "GZDLCNR"  > ${WORKDIR}/stack9
}

getState() {
    for i in {1..9};
    do
        stack=$(cat ${WORKDIR}/stack$i)
        crate=${stack: -1}
        echo -n $crate
    done
}

# Part1
createStart

cat ${BASEDIR}/input | grep -E '^move' | while read _ moves _ from _ to;
do
    stackFrom=${WORKDIR}/stack${from}
    stackTo=${WORKDIR}/stack${to}

    i=0
    while [ $i -lt $moves ];
    do
        i=$((i+1))
        stack=$(cat $stackFrom)
        crate=${stack: -1}
        sed -i "s/$crate$//" $stackFrom
        sed -i "s/$/$crate/" $stackTo
    done
done
part1=$(getState)

# Part2
createStart

cat ${BASEDIR}/input | grep -E '^move' | while read _ moves _ from _ to;
do
    stackFrom=${WORKDIR}/stack${from}
    stackTo=${WORKDIR}/stack${to}

    stack=$(cat $stackFrom)
    crates=${stack: -$moves}
    sed -i "s/$crates$//" $stackFrom
    sed -i "s/$/$crates/" $stackTo
done
part2=$(getState)

# Results
echo "Part1: $part1"
echo "Part1: $part2"
