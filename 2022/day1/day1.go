package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "sort"
    "strconv"
)


func main() {
        file, err := os.Open("input")
        if err != nil {
                fmt.Println("err Open")
        }

        elfs := make(map[int]int)
        elf := 1

        elfs[elf] = 0

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                line := bytes.TrimSpace(scanner.Bytes())

                if len(line) == 0 {
                    elf++
                    elfs[elf] = 0
                } else {
                    cal, err := strconv.Atoi(scanner.Text())
                    if err != nil {
                        fmt.Println("err Open")
                    }					
                    elfs[elf] += cal
                }

        }

        keys := make([]int, 0, len(elfs))
        for key := range elfs {
            keys = append(keys, key)
        }

        //fmt.Println(elfs)
        //fmt.Println(keys)		
        
        sort.SliceStable(keys, func(i, j int) bool{
            return elfs[keys[i]] < elfs[keys[j]]
        })		

        //fmt.Println(keys)

        // Part1
        fmt.Println("\n## Part 1")

        topElf := keys[len(keys)-1]
        topElfCal := elfs[topElf]

        fmt.Println("Top Elf: #", topElf)
        fmt.Println("Top Elf carrying:", topElfCal)

        // Part2
        fmt.Println("\n## Part 2")

        top3Elfs := 0
        for i := 1; i <= 3; i++ {
            pickElf := keys[len(keys)-i]
            elfCal := elfs[pickElf]
    
            fmt.Println("Elf spot : #", pickElf, "\t- carrying :", elfCal)
            top3Elfs += elfCal
        }
        fmt.Println("top3 Elfs carrying:", top3Elfs)
}