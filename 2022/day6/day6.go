package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func findUniqueCharset(input string, chars int) int {

    s := strings.Split(input, "")

    buffer := []string{}

    for i:=0;i<len(s);i++ {

        if(len(buffer) >= chars){
            buffer = buffer[1:]
        }
        buffer = append(buffer, s[i])
        
        check := map[string]int{}
        for k, v := range buffer {
            check[v] = k
        }
        uniqueChars := len(check)
        
        if(uniqueChars >= chars){ 
            return i+1 
        }
    }

    return 0
}

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("err Open")
    }

    part1 := 0
    part2 := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        part1 = findUniqueCharset(line, 4)
        part2 = findUniqueCharset(line, 14)
    }

    fmt.Println("Part1:", part1)
    fmt.Println("Part2:", part2)
}