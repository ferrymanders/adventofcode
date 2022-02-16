package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)


func replaceAt(s string, i int, c rune) string {
    r := []rune(s)
    r[i] = c
    return string(r)
}

func treeCounter(field []string, stepx int, stepy int) int {
    counter := 0
    x := 0
    y := 0

    for y < len(field) - stepy {
        y = y + stepy        
        line := string(field[y])

        x = x + stepx
        if x >= len(line) {
            x = x - len(line)
        }

        if strings.Contains(string(line[x]), "#") {
            counter++
        }             
    }

    return int(counter)
}

func main() {
    var field []string

    file, err := os.Open("input")
    if err != nil {
            fmt.Println("err Open")
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        field = append(field, line)
    }

    // Part 1
    slope := treeCounter(field, 3, 1)

    fmt.Println("# Part1")
    fmt.Println("Slope:", slope)

    // Part 2
    slope1 := treeCounter(field, 1, 1)
    slope2 := treeCounter(field, 3, 1)
    slope3 := treeCounter(field, 5, 1)
    slope4 := treeCounter(field, 7, 1)
    slope5 := treeCounter(field, 1, 2)

    product := slope1 * slope2 * slope3 * slope4 * slope5

    fmt.Println("# Part2")
    fmt.Println("Slope1:", slope1)
    fmt.Println("Slope2:", slope2)
    fmt.Println("Slope3:", slope3)
    fmt.Println("Slope4:", slope4)
    fmt.Println("Slope5:", slope5)
    fmt.Println("Product:", product)
}