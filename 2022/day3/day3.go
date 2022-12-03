package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
)

func uniqueChars(a string) string {
    chars := make(map[string]int)

    for i := 0; i < len(a); i++ {
        chars[string(a[i])] = 1
    }   

    keys := make([]string, 0, len(chars))
    for key := range chars {
        keys = append(keys, key)
    }

    output := new(bytes.Buffer)
    for _, value := range keys {
        fmt.Fprintf(output, "%s", value)
    }
    return output.String()
}

func compareItems(a []string, b int) string {
    items := make(map[string]int)

    for _, value := range a {
        for i := 0; i < len(value); i++ {
            items[string(value[i])]++
        }
    }    

    keys := make([]string, 0, len(items))
    for key := range items {
        keys = append(keys, key)
    }    

    for i := 0; i < len(keys); i++ {
        if items[keys[i]] == b {
            return keys[i]
        }
    }

    return "false"
}

func getPriority(a string) int {
    unicode := int(a[0])
    priority := 0
    if unicode < 91 { priority = unicode - 65 + 27 }
    if unicode > 91 { priority = unicode - 96 }

    return priority
}


func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("err Open")
    }

    part1 := 0
    part2 := 0

    elfCounter := 0
    elfGroup   := []string{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        inventory := scanner.Text()
        line := bytes.TrimSpace(scanner.Bytes())
        halfway := len(line) / 2

        // find item in both compartments
        compartments := []string{}
        compartments = append(compartments, uniqueChars(inventory[halfway:]))
        compartments = append(compartments, uniqueChars(inventory[:halfway]))
        output := compareItems(compartments, 2)

        // determine item priority 
        priority := getPriority(output)
        
        part1 += priority

        // part2
        if elfCounter == 0 { elfGroup = nil }
        elfCounter++
        elfGroup = append(elfGroup, uniqueChars(inventory))

        if elfCounter == 3 {
            output := compareItems(elfGroup, 3)
            priority := getPriority(output)
            part2 += priority

            elfCounter = 0 
        }
    }

    fmt.Println("Part1:", part1)
    fmt.Println("Part2:", part2)
}