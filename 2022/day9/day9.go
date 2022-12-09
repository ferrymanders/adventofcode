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

    if ( deltaX >=  2 && deltaY ==  0 ){ tLoc["x"]++ }
    if ( deltaX ==  0 && deltaY >=  2 ){ tLoc["y"]++ }
    if ( deltaX <= -2 && deltaY ==  0 ){ tLoc["x"]-- }
    if ( deltaX ==  0 && deltaY <= -2 ){ tLoc["y"]-- }

    if ( deltaX ==  2 && deltaY ==  1 ){ tLoc["x"]++; tLoc["y"]++ }
    if ( deltaX ==  2 && deltaY == -1 ){ tLoc["x"]++; tLoc["y"]-- }
    if ( deltaX ==  1 && deltaY ==  2 ){ tLoc["x"]++; tLoc["y"]++ }
    if ( deltaX == -1 && deltaY ==  2 ){ tLoc["x"]--; tLoc["y"]++ }

    if ( deltaX == -2 && deltaY == -1 ){ tLoc["x"]--; tLoc["y"]-- }
    if ( deltaX == -2 && deltaY ==  1 ){ tLoc["x"]--; tLoc["y"]++ }
    if ( deltaX == -1 && deltaY == -2 ){ tLoc["x"]--; tLoc["y"]-- }
    if ( deltaX ==  1 && deltaY == -2 ){ tLoc["x"]++; tLoc["y"]-- }

    if ( deltaX ==  2 && deltaY ==  2 ){ tLoc["x"]++; tLoc["y"]++ }
    if ( deltaX ==  2 && deltaY == -2 ){ tLoc["x"]++; tLoc["y"]-- }
    if ( deltaX == -2 && deltaY ==  2 ){ tLoc["x"]--; tLoc["y"]++ }
    if ( deltaX == -2 && deltaY == -2 ){ tLoc["x"]--; tLoc["y"]-- }
    
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
    t9Spots := map[string]int{ "0_0": 1 }
    hLoc := map[string]float64{ "x": 0, "y": 0, }
    tLoc := map[string]float64{ "x": 0, "y": 0, }

    t1Loc := map[string]float64{ "x": 0, "y": 0, }
    t2Loc := map[string]float64{ "x": 0, "y": 0, }
    t3Loc := map[string]float64{ "x": 0, "y": 0, }
    t4Loc := map[string]float64{ "x": 0, "y": 0, }
    t5Loc := map[string]float64{ "x": 0, "y": 0, }
    t6Loc := map[string]float64{ "x": 0, "y": 0, }
    t7Loc := map[string]float64{ "x": 0, "y": 0, }
    t8Loc := map[string]float64{ "x": 0, "y": 0, }
    t9Loc := map[string]float64{ "x": 0, "y": 0, }


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
                if ( checkAdjacent(hLoc, t1Loc) == false )  { t1Loc = moveT(hLoc, t1Loc); }
                if ( checkAdjacent(t1Loc, t2Loc) == false ) { t2Loc = moveT(t1Loc, t2Loc); }
                if ( checkAdjacent(t2Loc, t3Loc) == false ) { t3Loc = moveT(t2Loc, t3Loc); }
                if ( checkAdjacent(t3Loc, t4Loc) == false ) { t4Loc = moveT(t3Loc, t4Loc); }
                if ( checkAdjacent(t4Loc, t5Loc) == false ) { t5Loc = moveT(t4Loc, t5Loc); }
                if ( checkAdjacent(t5Loc, t6Loc) == false ) { t6Loc = moveT(t5Loc, t6Loc); }
                if ( checkAdjacent(t6Loc, t7Loc) == false ) { t7Loc = moveT(t6Loc, t7Loc); }
                if ( checkAdjacent(t7Loc, t8Loc) == false ) { t8Loc = moveT(t7Loc, t8Loc); }
                if ( checkAdjacent(t8Loc, t9Loc) == false ) { 
                    t9Loc = moveT(t8Loc, t9Loc)
                    tSpot := fmt.Sprintf("%g_%g", t9Loc["x"], t9Loc["y"])
                    t9Spots[tSpot] = 1
                }

            }
        }
    }

    part1 = len(tSpots)
    part2 = len(t9Spots)

    fmt.Println("Part1:", part1)
    fmt.Println("Part2:", part2)
}