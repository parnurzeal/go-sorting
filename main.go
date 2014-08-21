package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func BubbleSort(data Interface) {
	n := data.Len()
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if !data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}

func SelectionSort(data Interface) {
	n := data.Len()
	for i := 0; i < n; i++ {
		max := i
		for j := i + 1; j < n; j++ {
			if !data.Less(max, j) {
				max = j
			}
		}
		data.Swap(i, max)
	}
}

func MergeSort(data Interface) {
	n := data.Len()
	MergeSortInternal(data, 0, n)
}

func MergeSortInternal(data Interface, f, l int) {
	n := l - f + 1
	if f <= l {
		return
	}
	m := (l - f) / 2
	MergeSortInternal(data, f, m)
	MergeSortInternal(data, m+1, l)
	newSlice := make([]Interface, n)
	// add more
}

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(a, b int)      { s[a], s[b] = s[b], s[a] }
func (s IntSlice) Less(a, b int) bool { return s[a] < s[b] }

var digitRegexp = regexp.MustCompile("-?[0-9]+")

func readIntStdin() []int {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	ints := digitRegexp.FindAll(bytes, -1)
	var nums []int
	for _, numBytes := range ints {
		num, err := strconv.Atoi(string(numBytes))
		if err != nil {
			fmt.Println(err)
		}
		nums = append(nums, int(num))
	}
	return nums
}

func main() {
	nums := readIntStdin()
	nums1 := intSlice(nums)

	fmt.Println("Original\t", nums)

	fmt.Println("Bubble Sort\t", nums1.BubbleSort())
	numsBubble := make([]int, len(nums))
	copy(numsBubble, nums)
	BubbleSort(IntSlice(numsBubble))
	fmt.Println("Bubble Sort2\t", numsBubble)

	fmt.Println("Selection Sort\t", nums1.SelectionSort())
	numsSelection := make([]int, len(nums))
	copy(numsSelection, nums)
	SelectionSort(IntSlice(numsSelection))
	fmt.Println("Selection Sort2\t", numsSelection)

	fmt.Println("Insertion Sort\t", nums1.InsertionSort())
}
