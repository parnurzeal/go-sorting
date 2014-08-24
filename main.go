package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Get(i int) interface{}
	Set(i int, elem interface{})
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

func InsertionSort(data Interface) {
	n := data.Len()
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}

func MergeSort(data Interface) {
	n := data.Len()
	MergeSortInternal(data, 0, n)
}

func MergeSortInternal(data Interface, f, l int) {
	n := l - f
	if n < 2 {
		return
	}
	m := (l-f)/2 + f
	MergeSortInternal(data, f, m)
	MergeSortInternal(data, m, l)
	newSlice := make([]interface{}, n)
	s1, s2, s3 := f, m, 0
	for s1 < m && s2 < l {
		if data.Less(s1, s2) {
			newSlice[s3] = data.Get(s1)
			s1++
			s3++
		} else {
			newSlice[s3] = data.Get(s2)
			s2++
			s3++
		}
	}
	for s1 < m {
		newSlice[s3] = data.Get(s1)
		s3++
		s1++
	}
	for s2 < l {
		newSlice[s3] = data.Get(s2)
		s3++
		s2++
	}
	for i := 0; i < n; i++ {
		data.Set(i+f, newSlice[i])
	}
	return
}

func QuickSort(data Interface) {
	rand.Seed(40)
	QuickSortInternal(data, 0, data.Len())
}

func QuickSortInternal(data Interface, f, l int) {
	n := l - f
	if n <= 1 {
		return
	}
	pivot := rand.Intn(n) + f
	data.Swap(f, pivot)
	idx := f
	for i := f + 1; i < l; i++ {
		if data.Less(i, f) {
			idx++
			data.Swap(i, idx)
		}
	}
	data.Swap(f, idx)
	QuickSortInternal(data, f, idx)
	QuickSortInternal(data, idx+1, l)
}

type IntSlice []int

func (s IntSlice) Len() int                    { return len(s) }
func (s IntSlice) Swap(a, b int)               { s[a], s[b] = s[b], s[a] }
func (s IntSlice) Less(a, b int) bool          { return s[a] < s[b] }
func (s IntSlice) Get(i int) interface{}       { return s[i] }
func (s IntSlice) Set(i int, elem interface{}) { s[i] = elem.(int) }

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
	numsInsertion := make([]int, len(nums))
	copy(numsInsertion, nums)
	InsertionSort(IntSlice(numsInsertion))
	fmt.Println("Insertion Sort2\t", numsInsertion)

	numsMerge := make([]int, len(nums))
	copy(numsMerge, nums)
	MergeSort(IntSlice(numsMerge))
	fmt.Println("Merge Sort2\t", numsMerge)

	numsQuick := make([]int, len(nums))
	copy(numsQuick, nums)
	QuickSort(IntSlice(numsQuick))
	fmt.Println("Quick Sort2\t", numsQuick)

}
