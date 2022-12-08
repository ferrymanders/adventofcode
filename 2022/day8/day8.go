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

    part1 := 0
    part2 := 0

    trees := [][]int{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        treeLineRaw := strings.Split(line, "")

        // This is stupid..
        treeLine := []int{}
        for _, tree := range treeLineRaw {
            height, _ := strconv.Atoi(tree)
            treeLine = append(treeLine, height)
        }

        trees = append(trees, treeLine)
    }

    // Part1
    countVisible := 0
    
    southEdge := len(trees)
    for treeRow, valueRow := range trees {
        eastEdge := len(valueRow)
        for treeCol, treeHeight := range valueRow {
            hiddenSide := map[string]int{}

            // check west
            for i:=treeCol-1; i >= 0; i-- {
                if ( trees[treeRow][i] >= treeHeight ){ hiddenSide["west"] = 1 }
            }
            // check east
            for i:=treeCol+1; i < eastEdge; i++ { 
                if ( trees[treeRow][i] >= treeHeight ){ hiddenSide["east"] = 1 }
            }

            // check north
            for i:=treeRow-1; i >= 0; i-- { 
                if ( trees[i][treeCol] >= treeHeight ){ hiddenSide["north"] = 1 }
            }
            // check south
            for i:=treeRow+1; i < southEdge; i++ { 
                if ( trees[i][treeCol] >= treeHeight ){ hiddenSide["south"] = 1 }
            }

            if ( len(hiddenSide) < 4 ){
                countVisible++
            }

        }
    }
    part1 = countVisible

    // Part2
    bestScore := 0
    for treeRow, valueRow := range trees {
        eastEdge := len(valueRow)
        for treeCol, treeHeight := range valueRow {
            sideCounter := map[string]int{
                "west":  0,
                "east":  0,
                "north": 0,
                "south": 0,
            }

            // check west
            for i:=treeCol-1; i >= 0; i-- {
                if ( trees[treeRow][i] >= treeHeight ){ sideCounter["west"]++; break }else{ sideCounter["west"]++ }
            }
            // check east
            for i:=treeCol+1; i < eastEdge; i++ { 
                if ( trees[treeRow][i] >= treeHeight ){ sideCounter["east"]++; break }else{ sideCounter["east"]++ }
            }

            // check north
            for i:=treeRow-1; i >= 0; i-- { 
                if ( trees[i][treeCol] >= treeHeight ){ sideCounter["north"]++; break }else{ sideCounter["north"]++ }
            }
            // check south
            for i:=treeRow+1; i < southEdge; i++ { 
                if ( trees[i][treeCol] >= treeHeight ){ sideCounter["south"]++; break }else{ sideCounter["south"]++ }
            }

            currentScore := sideCounter["west"] * sideCounter["east"] * sideCounter["north"] * sideCounter["south"]

            if ( currentScore > bestScore ){
                bestScore = currentScore
            }

        }
    }
    part2 = bestScore

    fmt.Println("Part1:", part1)
    fmt.Println("Part2:", part2)
}