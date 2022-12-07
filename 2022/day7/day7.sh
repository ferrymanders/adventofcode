#!/bin/bash

diskSize=70000000
updateSize=30000000

BASEDIR=${PWD}/work/data
CURRENTDIR=${PWD}

[ -d $BASEDIR ] || mkdir -p $BASEDIR

cat input | while read line;
do
    checkCommand=$(echo $line | grep -cE '^\$')
    checkOutputDir=$(echo $line | grep -cE '^dir ')

    if [ $checkCommand == 1 ];
    then
        read foo action arg <<< `echo $line`

        case "$action" in
            cd) 
                [ -d $CURRENTDIR/$arg ] || mkdir $CURRENTDIR/$arg
                case "$arg" in
                    /)
                        cd $BASEDIR
                        CURRENTDIR=$BASEDIR
                        ;;
                    ..)
                        cd ..
                        CURRENTDIR=$PWD
                        ;;                   
                    *)
                        cd $arg
                        CURRENTDIR=$CURRENTDIR/$arg
                        ;;
                esac
                ;;
            *) ;;    
        esac
    fi

    if [ $checkOutputDir == 1 ];
    then
        read foo dir <<< `echo $line`
        [ -d $CURRENTDIR/$dir ] || mkdir $CURRENTDIR/$dir
    fi

    if [ $checkCommand != 1 ] && [ $checkOutputDir != 1 ];
    then
        read size file <<< `echo $line`
        [ -f $file ] || dd bs=1 count=$size if=/dev/zero of=$file
    fi
done

counter=0
list=""

while read dir;
do
    dirCounter=0

    for file in $(find $dir -type f);
    do
        read fileSize fileName <<< `du -b $file`
        dirCounter=$((dirCounter + fileSize))
    done

    if [ $dirCounter -lt 100001 ];
    then
        counter=$((counter + dirCounter))
    fi

    list="$list $dirCounter;$dir"

done <<< `find $BASEDIR -type d`

totalCounter=0
for file in $(find $BASEDIR -type f);
do
    read fileSize fileName <<< `du -b $file`
    totalCounter=$((totalCounter + fileSize))
done

currentSpace=$((diskSize - totalCounter))
neededSpace=$((updateSize - currentSpace))

closest=$updateSize
for item in $list;
do
     IFS=';' read dirSize dirName <<< `echo $item`
     [ $dirSize -lt $closest ] && [ $dirSize -gt $neededSpace ] && closest=$dirSize
done


echo "Part1: $counter"
echo "Part2: $closest"