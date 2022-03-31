package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)


func main() {
    questionMap := make(map[string]int)
    part1 := int(0)
    part2 := int(0)
    members := int(0)

    file, err := os.Open("input")
    if err != nil {
        fmt.Println("err Open")
    }

    // Collect data per customs declaration group
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        // New group starts on empty newline
        matched, _ := regexp.MatchString(`^\s*$`, line)
        if matched {
            for key, value := range questionMap {
                part1++
                if ( value == members ) {
                    part2++
                }
                delete(questionMap, key)
            }
            members = int(0)
        } else {
            for _, char := range line {
                questionMap[string(char)]++
            }
            members++
        }

    }

    // Part 1
    fmt.Println("Part1: ", part1)
    // Part 2
    fmt.Println("Part2: ", part2)

}