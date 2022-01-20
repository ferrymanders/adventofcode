package main

import (
    "bufio"
    "fmt"
    "os"
		"strconv"
		"regexp"
		"strings"
)

func main() {
		counter := 0
		counter2 := 0

		file, err := os.Open("input")
		if err != nil {
				fmt.Println("err Open")
		}

		scanner := bufio.NewScanner(file)

		// Line example "4-5 t: ftttttrvts"
		re := regexp.MustCompile("(?P<from>[0-9]+)-(?P<to>[0-9]+) (?P<needle>[a-z]): (?P<haystack>[a-z]+)$")

		for scanner.Scan() {
				line := scanner.Text()
				search := re.FindStringSubmatch(line)
				result := make(map[string]string)

				for i, name := range re.SubexpNames() {
						if i != 0 && name != "" {
								result[name] = search[i]
						}
				}

				from, err := strconv.Atoi(result["from"])
				to, err := strconv.Atoi(result["to"])
				if err != nil {
					fmt.Println(err)
				}

				// Part 1 
				count := strings.Count(result["haystack"], result["needle"])
				if count >= from && count <= to {
						counter++
				}

				// Part 2
				s := strings.Split(result["haystack"], "")
				test := false
				if(from-1 <= len(result["haystack"])) {
						if s[from-1] == result["needle"] {
							test = true
						}
				}
				if(to-1 <= len(result["haystack"])) {
						if s[to-1] == result["needle"] {
							test = true
						}
						if s[from-1] == result["needle"] && s[to-1] == result["needle"] {
							test = false
						}
				}

				if test == true {
						counter2++
				}
		}

		fmt.Println("part 1: ", counter)
		fmt.Println("part 2: ", counter2)
}