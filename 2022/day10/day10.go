package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getPixelNr(cycle int) int {
    if ( cycle > 200 ){ cycle = cycle - 200 }
    if ( cycle > 160 ){ cycle = cycle - 160 }
    if ( cycle > 120 ){ cycle = cycle - 120 }
    if ( cycle >  80 ){ cycle = cycle -  80 }
    if ( cycle >  40 ){ cycle = cycle -  40 }
    return cycle
}

func drawPixel(line string, cycle, sprite int) string {
    pixel := "."
    pixelLoc := getPixelNr(cycle) - 1
    deltaS := Abs(pixelLoc - sprite)
    if ( deltaS <= 1 ){ pixel = "#" }
    return fmt.Sprintf("%s%s", line,pixel)
}

func main() {
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("err Open")
    }

    part1 := 0
    part2 := 0

    cycle := 0
    x := 1

    signal := map[int]int{}
    screen := map[int]string{}

    lineNr := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        cycle++

        if ( getPixelNr(cycle) == 1 ){ lineNr++ }

        s := strings.Split(scanner.Text(), " ")
        command := s[0]
        switch command {
            case "noop":
                signal[cycle] = x * cycle
                screen[lineNr] = drawPixel(screen[lineNr], cycle, x)
            case "addx":
                arg, _ := strconv.Atoi(s[1])
                signal[cycle] = x * cycle
                screen[lineNr] = drawPixel(screen[lineNr], cycle, x)
                cycle++
                signal[cycle] = x * cycle
                if ( getPixelNr(cycle) == 1 ){ lineNr++ }
                screen[lineNr] = drawPixel(screen[lineNr], cycle, x)
                x = x + arg
        }
    }

    part1Signals := map[int]int{
         20: signal[20],
         60: signal[60],
        100: signal[100],
        140: signal[140],
        180: signal[180],
        220: signal[220],
    }

    for _, v := range part1Signals {
        part1 += v
    }


    fmt.Println("Part1:", part1)
    fmt.Println("Part2:", part2)
    fmt.Println(screen[1])
    fmt.Println(screen[2])
    fmt.Println(screen[3])
    fmt.Println(screen[4])
    fmt.Println(screen[5])
    fmt.Println(screen[6])
}