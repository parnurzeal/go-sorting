package main

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
