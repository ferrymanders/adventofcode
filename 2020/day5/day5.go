package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
    "strconv"
)

func in_array(val int, array []int) (exists bool, index int) {
    exists = false
    index = -1

    for i, v := range array {
            if val == v {
                    index = i
                    exists = true
                    return
            }
    }

    return
}

func main() {
    replacer := strings.NewReplacer("F", "0", 
                                    "B", "1",
                                    "L", "0",
                                    "R", "1")

    var seats []int

    file, err := os.Open("input")
    if err != nil {
        fmt.Println("err Open")
    }

    // Collect data per boardingpass
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        rowData := line[0:7]
        row, _ := strconv.ParseInt(replacer.Replace(rowData), 2, 8)
        
        columnData := line[7:10]
        column, _ := strconv.ParseInt(replacer.Replace(columnData), 2, 8)
        
        seatId := int(row * 8 + column)

        seats = append(seats, seatId)
    }

    // Part 1
    sort.Ints(seats)
    lastSeat := seats[len(seats)-1]
    fmt.Println("Last seat: ", lastSeat)

    // Part 2
    firstSeat := seats[0]
    mySeat := int(0)

    for i := firstSeat; i <= lastSeat; i++ {
        test, _ := in_array(i, seats)
        if test == false {
            prevCheck, _ := in_array(i-1, seats)
            nextCheck, _ := in_array(i+1, seats)
            if ( prevCheck && nextCheck ) {
                mySeat = i
            }
        }
    }
    fmt.Println("My seat: ", mySeat)

}