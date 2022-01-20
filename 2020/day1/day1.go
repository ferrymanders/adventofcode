package main

import (
    "bufio"
    "fmt"
    "os"
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
		var numbers []int
		needle := 2020

		file, err := os.Open("numbers")
		if err != nil {
				fmt.Println("err Open")
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
				number, err := strconv.Atoi(scanner.Text())
				if err != nil {	
						fmt.Println("err Atoi")
				}

				numbers = append(numbers, number)
		}

		for i, s := range numbers {
				searchNumber := needle - s
				test, index := in_array(searchNumber, numbers)
				if test == true {
						result := numbers[index]
						answer := numbers[i] * result
						fmt.Println(s, result, answer)
						break
				}
		}

		stop := false
		for i, s := range numbers {
				searchNumber := needle - s
				for ii, ss := range numbers {
						test := false
						searchSecondNumber := searchNumber - ss
						test, index := in_array(searchSecondNumber, numbers)
						if test == true {
								firstNumber := numbers[i]
								secondNumber := numbers[ii]
								thirdNumber := numbers[index]
								answer := firstNumber * secondNumber * thirdNumber
								fmt.Println(firstNumber, secondNumber, thirdNumber, answer)
								stop = true
								break
						}
						if stop == true {
								break
						}
				}
		}
}