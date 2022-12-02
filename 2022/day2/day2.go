package main

import (
    "bufio"
    "fmt"
    "os"
)


func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("err Open")
    }

    outcomePart1 := map[string]int{
        "A X": 4,
        "A Y": 8,
        "A Z": 3,
        "B X": 1,
        "B Y": 5,
        "B Z": 9,
        "C X": 7,
        "C Y": 2,
        "C Z": 6,
    }

    outcomePart2 := map[string]int{
        "A X": 3,
        "A Y": 4,
        "A Z": 8,
        "B X": 1,
        "B Y": 5,
        "B Z": 9,
        "C X": 2,
        "C Y": 6,
        "C Z": 7,
    }

    scorePart1 := 0
    scorePart2 := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        scorePart1 += outcomePart1[scanner.Text()]
        scorePart2 += outcomePart2[scanner.Text()]
    }

    fmt.Println("Part1 score:", scorePart1)
    fmt.Println("Part2 score:", scorePart2)
}