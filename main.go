package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

type intSlice []int

func (s intSlice) swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s intSlice) BubbleSort() intSlice {
	newS := make(intSlice, len(s))
	copy(newS, s)
	for i := 0; i < len(newS); i++ {
		for j := 0; j < len(newS)-1; j++ {
			if newS[j] > newS[j+1] {
				newS.swap(j, j+1)
			}
		}
	}
	return newS
}

func (s intSlice) SelectionSort() intSlice {
	newS := make(intSlice, len(s))
	copy(newS, s)
	for i := 0; i < len(newS); i++ {
		k := i
		for j := i + 1; j < len(newS); j++ {
			if newS[k] > newS[j] {
				k = j
			}
		}
		newS.swap(i, k)
	}
	return newS
}

func (s intSlice) InsertionSort() intSlice {
	newS := make(intSlice, len(s))
	copy(newS, s)
	for i := 0; i < len(newS); i++ {
		for j := i; j > 0 && newS[j] < newS[j-1]; j-- {
			newS.swap(j, j-1)
		}
	}
	return newS
}

func (s intSlice) MergeSort() intSlice {
	return []int{}
}

func (s intSlice) QuickSort() intSlice {
	return []int{}
}

var digitRegexp = regexp.MustCompile("-?[0-9]+")

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	ints := digitRegexp.FindAll(bytes, -1)
	var nums intSlice
	for _, numBytes := range ints {
		num, err := strconv.Atoi(string(numBytes))
		if err != nil {
			fmt.Println(err)
		}
		nums = append(nums, int(num))
	}
	fmt.Println("Original", nums)
	fmt.Println("Bubble Sort", nums.BubbleSort())
	fmt.Println("Selection Sort", nums.SelectionSort())
	fmt.Println("Insertion Sort", nums.InsertionSort())
}
