package main

//aa bb cc aa cc cc cc aa ab ac bb
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var numWords int
	resultArr := []string{}
	arrNumWords := []int{}
	mapWordsNum := map[string]int{}
	uniqWords := []string{}
	words := ReadStdin()
	if len(words) == 0 {
		fmt.Println("")
		return
	}
	_, err := fmt.Scanf("%d", &numWords)
	if err != nil {
		fmt.Println("Error read K")
		return
	}
	//считаем количество слов и добавляем это в мапу
	// key - string, value - countWords
	for _, v := range words {
		mapWordsNum[v]++
	}

	//разделяем мапу на slice value
	for k, v := range mapWordsNum {
		arrNumWords = append(arrNumWords, v)
		uniqWords = append(uniqWords, k)
	}
	if len(words) < numWords {
		resultArr = SortWords(uniqWords, mapWordsNum)
		fmt.Println(resultArr)
	} else {
		//сортируем слайс по возрастанию
		arrNumWords = SortNumWords(arrNumWords)

		//сравниваем отсортированный слайс с value в мапе
		resultSlice := CompareNum(numWords, mapWordsNum, arrNumWords)

		result := strings.Join(resultSlice, " ")
		fmt.Println(result)
	}
}

func ReadStdin() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Fields(line)
	return words
}

func SortNumWords(arr []int) []int {
	for j := 1; j < len(arr); j++ {
		for i := 0; i < len(arr)-j; i++ {
			if arr[i] < arr[i+1] {
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
			}
		}
	}
	return arr
}

func CompareNum(numWords int, mapWords map[string]int, sliceNumCount []int) []string {
	resultSclice := []string{}
	for i := range numWords {
		for k, v := range mapWords {
			if sliceNumCount[i] == v {
				resultSclice = append(resultSclice, k)
			}
		}
	}
	return resultSclice
}

func SortWords(uniqueWords []string, freq map[string]int) []string {
	sort.Slice(uniqueWords, func(i, j int) bool {
		return freq[uniqueWords[i]] > freq[uniqueWords[j]]
	})
	return uniqueWords
}