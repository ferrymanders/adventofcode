package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("err Open")
    }

    part1 := map[string]string{
        "1": "RSLFQ",
        "2": "NZQGPT",
        "3": "SMQB",
        "4": "TGZJHCBQ",
        "5": "PHMBNFS",
        "6": "PCQNSLVG",
        "7": "WCF",
        "8": "QHGZWVPM",
        "9": "GZDLCNR",
    }

    part2 := map[string]string{
        "1": "RSLFQ",
        "2": "NZQGPT",
        "3": "SMQB",
        "4": "TGZJHCBQ",
        "5": "PHMBNFS",
        "6": "PCQNSLVG",
        "7": "WCF",
        "8": "QHGZWVPM",
        "9": "GZDLCNR",
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "move") {

            re := regexp.MustCompile("move (?P<moves>[0-9]+) from (?P<stack1>[0-9]+) to (?P<stack2>[0-9]+)")
            if re.MatchString(line) {
                search := re.FindStringSubmatch(line)
                result := make(map[string]string)

                for i, name := range re.SubexpNames() {
                    if i != 0 && name != "" {
                        result[name] = search[i]
                    }
                }

                moves, _ := strconv.Atoi(result["moves"])
                
                // Part1
                stack1 := part1[result["stack1"]]
                stack2 := part1[result["stack2"]]
                for i:=0; i<moves; i++ {
                    x := stack1[len(stack1)-1:]
                    stack1 = stack1[:len(stack1)-1]
                    stack2 = stack2 + x
                }
                part1[result["stack1"]] = stack1
                part1[result["stack2"]] = stack2



                // Part2
                P2stack1 := part2[result["stack1"]]
                P2stack2 := part2[result["stack2"]]

                part2[result["stack2"]] = P2stack2 + P2stack1[len(P2stack1)-moves:]
                part2[result["stack1"]] = P2stack1[:len(P2stack1)-moves]
            }
        }
        
    }

    fmt.Println("Part1:", part1)
    fmt.Println("Part2:", part2)
}