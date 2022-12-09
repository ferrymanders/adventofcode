package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "regexp"
    "strconv"
)

func checkAdjacent(hLoc, tLoc map[string]float64) bool {
    adjacent := false
    deltaX := math.Abs(hLoc["x"] - tLoc["x"])
    deltaY := math.Abs(hLoc["y"] - tLoc["y"])
    if ( deltaX <= 1 ) && ( deltaY <= 1 ){ adjacent = true }
    return adjacent
}

func moveT(hLoc, tLoc map[string]float64) map[string]float64 {
    deltaX := hLoc["x"] - tLoc["x"]
    deltaY := hLoc["y"] - tLoc["y"]

    if ( deltaX ==  2 && deltaY ==  0 ){ tLoc["x"]++; }
    if ( deltaX ==  0 && deltaY ==  2 ){ tLoc["y"]++; }
    if ( deltaX == -2 && deltaY ==  0 ){ tLoc["x"]--; }
    if ( deltaX ==  0 && deltaY == -2 ){ tLoc["y"]--; }

    if ( deltaX ==  2 && deltaY ==  1 ){ tLoc["x"]++; tLoc["y"]++ }
    if ( deltaX ==  2 && deltaY == -1 ){ tLoc["x"]++; tLoc["y"]-- }
    if ( deltaX ==  1 && deltaY ==  2 ){ tLoc["x"]++; tLoc["y"]++ }
    if ( deltaX == -1 && deltaY ==  2 ){ tLoc["x"]--; tLoc["y"]++ }

    if ( deltaX == -2 && deltaY == -1 ){ tLoc["x"]--; tLoc["y"]-- }
    if ( deltaX == -2 && deltaY ==  1 ){ tLoc["x"]--; tLoc["y"]++ }
    if ( deltaX == -1 && deltaY == -2 ){ tLoc["x"]--; tLoc["y"]-- }
    if ( deltaX ==  1 && deltaY == -2 ){ tLoc["x"]++; tLoc["y"]-- }
    
    return tLoc
}


func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("err Open")
    }

    part1 := 0
    part2 := 0

    tSpots := map[string]int{ "0_0": 1 }
    hLoc := map[string]float64{ "x": 0, "y": 0, }
    tLoc := map[string]float64{ "x": 0, "y": 0, }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        re := regexp.MustCompile("(?P<direction>[UDLR]+) (?P<steps>[0-9]+)")
        if re.MatchString(line) {
            search := re.FindStringSubmatch(line)
            result := make(map[string]string)

            for i, name := range re.SubexpNames() {
                if i != 0 && name != "" {
                    result[name] = search[i]
                }
            }

            direction := result["direction"]
            steps, _ := strconv.Atoi(result["steps"])

            // Move H
            for i := 0; i<steps; i++  {
                switch direction {
                    case "U": 
                        hLoc["y"]++
                    case "D":
                        hLoc["y"]--
                    case "L":
                        hLoc["x"]--
                    case "R":
                        hLoc["x"]++
                }
                // check if T is adjacent to H
                if ( checkAdjacent(hLoc, tLoc) == false ) {
                    tLoc = moveT(hLoc, tLoc)
                    tSpot := fmt.Sprintf("%g_%g", tLoc["x"], tLoc["y"])
                    tSpots[tSpot] = 1
                }

            }
        }
    }

    part1 = len(tSpots)

    fmt.Println("Part1:", part1)
    fmt.Println("Part2:", part2)
}