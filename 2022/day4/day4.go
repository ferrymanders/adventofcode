package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)


func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("err Open")
    }

    containCounter := 0
    overlapCounter := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        
        sections := strings.Split(scanner.Text(), ",")
        section1 := strings.Split(sections[0], "-")
        section2 := strings.Split(sections[1], "-")
        s11, _ := strconv.Atoi(section1[0]) // beginning section 1
        s12, _ := strconv.Atoi(section1[1]) // ending section 1
        s21, _ := strconv.Atoi(section2[0]) // beginning section 2
        s22, _ := strconv.Atoi(section2[1]) // ending section 2

        // part1
        contain := false
        if s11 <= s21 && s12 >= s22 { contain = true }
        if s21 <= s11 && s22 >= s12 { contain = true }

        if contain == true { containCounter++ }

        // part2
        overlap := false
        if s11 <= s22 && s11 >= s21 { overlap = true }
        if s21 <= s12 && s21 >= s11 { overlap = true }

        if overlap == true { overlapCounter++ }
    }

    fmt.Println("Part1:", containCounter)
    fmt.Println("Part2:", overlapCounter)
}