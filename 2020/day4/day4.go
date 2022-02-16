package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
    "strconv"
)

func findData(data string) map[string]string {
    output := make(map[string]string)
    chunks := strings.Split(data, " ")

    for _, value := range chunks {
        if len(value) > 0 {
            parts := strings.Split(value, ":")
            output[parts[0]] = parts[1]
        }
    }

    return output
}

func checkValid(data map[string]string) bool {
    counter := 0
    if _, ok := data["byr"]; ok { counter++ }
    if _, ok := data["iyr"]; ok { counter++ }
    if _, ok := data["eyr"]; ok { counter++ }
    if _, ok := data["hgt"]; ok { counter++ }
    if _, ok := data["hcl"]; ok { counter++ }
    if _, ok := data["ecl"]; ok { counter++ }
    if _, ok := data["pid"]; ok { counter++ }

    if counter >= 7 {
        return true
    } else {
        return false
    }
}


func checkValidP2(data map[string]string) bool {
    counter := 0

    if val, ok := data["byr"]; ok {
        byr, _ := strconv.Atoi(val)
        if byr >= 1920 && byr <= 2002 {
            counter++ 
        }
    }

    if val, ok := data["iyr"]; ok {         
        iyr, _ := strconv.Atoi(val)
        if iyr >= 2010 && iyr <= 2020 {
            counter++ 
        }
    }

    if val, ok := data["eyr"]; ok { 
        eyr, _ := strconv.Atoi(val)
        if eyr >= 2020 && eyr <= 2030 {
            counter++ 
        }
    }

    if val, ok := data["hgt"]; ok {
        re := regexp.MustCompile("(?P<height>[0-9]+)(?P<format>[a-z]+)")
        if re.MatchString(val) {
            search := re.FindStringSubmatch(val)
            result := make(map[string]string)

            for i, name := range re.SubexpNames() {
                if i != 0 && name != "" {
                    result[name] = search[i]
                }
            }
        
            height, _ := strconv.Atoi(result["height"])
            switch result["format"] {
                case "cm":
                    if height >= 150 && height <= 193 {
                        counter++
                    }
                case "in":
                    if height >= 59 && height <= 76 {
                        counter++
                    }
            }    
        }
    }

    if val, ok := data["hcl"]; ok { 
        re := regexp.MustCompile("#[0-9a-f]{6}")
        if re.MatchString(val) {
            counter++
        }
    }

    if val, ok := data["ecl"]; ok { 
        re := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
        if re.MatchString(val) {
            counter++
        }
    }
    
    if val, ok := data["pid"]; ok { 
        re := regexp.MustCompile("^[0-9]{9}$")
        if re.MatchString(val) {
            counter++
        }
    }

    if counter >= 7 {
        return true
    } else {
        return false
    }
}

func main() {
    var passport []string
    data := ""

    file, err := os.Open("input")
    if err != nil {
        fmt.Println("err Open")
    }

    // Collect data per passport
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        // New passport starts on empty newline
        matched, _ := regexp.MatchString(`^\s*$`, line)
        if matched {
            passport = append(passport, data)
            data = ""
        } else {
            data = data + " " + line
        }
    }
    // we also need the last passport
    passport = append(passport, data)

    // Part 1
    counter := 0
    for _, line := range passport {

        data := findData(line)
        check := checkValid(data)
        if check {
            counter++
        }
    }
    fmt.Println("Part1:", counter)

    // Part 2
    counterP2 := 0
    for _, line := range passport {

        data := findData(line)      
        check := checkValidP2(data)
        if check {
            counterP2++
        }
    }

    fmt.Println("Part2:", counterP2)
}