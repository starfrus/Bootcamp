package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numsSliceOne, err1 := readNumsSlice()
	numsSliceTwo, err2 := readNumsSlice()
	resMap := map[int]bool{}
	if err1 != true && err2 != true {
		return
	} else {
		resMap = uniqSliceNums(numsSliceOne, numsSliceTwo)
	}
	printResult(numsSliceOne, resMap)
}

func readNumsSlice() ([]int, bool) {
	numsSlice := []int{}

	buffer := bufio.NewReader(os.Stdin)
	line, _ := buffer.ReadString('\n')
	line = strings.TrimSpace(line)

	numLine := strings.SplitSeq(line, " ")

	for val := range numLine {
		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Invalid input")
			return numsSlice, false
		} else {
			numsSlice = append(numsSlice, num)
		}
	}
	return numsSlice, true
}

func uniqSliceNums(s1, s2 []int) map[int]bool {
	resMap := map[int]bool{}
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				resMap[v1] = true
			}
		}
	}
	return resMap
}

func printResult(s1 []int, m1 map[int]bool) {
	if len(m1) == 0 {
		fmt.Println("Empty intersection")
	} else {
		for _, v1 := range s1 {
			for k := range m1 {
				if v1 == k {
					fmt.Print(v1, " ")
				}
			}
		}
		fmt.Printf("\b")
	}
}
