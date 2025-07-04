package ntu

import (
	"algorithms/ut"
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func Test_ch6(t *testing.T) {
	tests := []struct {
		name string
		f    func(*testing.T)
	}{
		// {"cyclicBinarySearch", cyclicBinarySearchTest},
		// {"quickSort", quickSortTest},
		// {"specialBinarySearch", specialBinarySearchTest},
		// {"stutteringSubsequence", stutteringSubsequenceTest},
		// {"interpolationSearch", interpolationSearchTest},
		// {"straightRadixSort", straightRadixSortTest},
		// {"mergeSort", mergeSortTest},
		// {"FindMinAndMax", FindMinAndMaxTest},
		// {"KthSmallestElement", KthSmallestElementTest},
		// {"computeNext", computeNextTest},
		// {"stringMatch", stringMatchKMPTest},
		// {"minimumEditDistance", minimumEditDistanceTest},
		{"HW6_24_03", HW6_24_03_Test},
		// {"RoundRobinSchedule", RoundRobinScheduleTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(t)
		})
	}
}

func RoundRobinScheduleTest(t *testing.T) {
	tests := []struct {
		left        int
		right       int
		playerCount int
	}{
		{1, 4, 4},
		// {1, 8, 8},
	}
	for _, tt := range tests {
		x := CH6{}
		// x.RoundRobinSchedule(tt.left, tt.right, tt.playerCount, "Begin")
		// x.RoundRobinScheduleGrok(tt.left, tt.right, tt.playerCount)
		// got := x.schedule([]int{1, 2, 3, 4})
		// got := x.scheduleCGPT([]int{1, 2, 3, 4, 5, 6, 7, 8})
		x.scheduleRoundRobin(8)
		// x.RoundRobinScheduleGeminiJuha(tt.playerCount)
		t.Log(tt.left)
		// got := x.RoundRobinCGPT([]int{1, 2, 3, 4})
		// t.Logf("%v", got)
	}
}
func HW6_24_03_Test(t *testing.T) {
	tests := []struct {
		name string
		seq  []int
		// want []int
	}{
		{"worst case 3", []int{1, 3, 5}},
		{"worst case 4", []int{1, 2, 3, 7}},
		{"worst case 5", []int{1, 2, 6, 8, 9}},
		{"worst case 6", []int{1, 2, 7, 10, 9, 10}},
		{"", []int{3, 2, 1}},
		{"", []int{1, 2, 4}},
		{"", []int{3, 4, 2, 1, 5}},
		{"3 el. ans. is valid", []int{1, 2, 3, 7, 8, 6}},
		{"returns a set of length n-1", []int{2, 22, 4, 5, 12, 32, 14, 25, 3, 13}}, // iters = n
	}
	for _, tt := range tests {
		x := CH6{}
		got := x.HW6_24_03_pseudo(tt.seq)
		length, sum := len(tt.seq), ut.Sum(got)
		if sum%length != 0 {
			t.Errorf("FAIL got: %v, got sum: %d, sum mod %d = %d", got, sum, length, sum%length)
		} else {
			ratio := float64(x.iters) / float64(length)
			t.Logf("SUCCESS got: %d, wantSum: %d, iters: %d, ratio: %.2f", got, sum, x.iters, ratio)
		}
	}
}
func MajorityTest(t *testing.T) {
	tests := []struct {
		name string
		seq  []int
		want int
	}{
		{"", []int{1, 2, 3}, -1},
		{"", []int{1, 1}, 1},
		{"", []int{1}, 1},
		{"", []int{1, 1, 1, 2, 2, 3}, -1},
		{"", []int{1, 1, 1, 1, 2, 2, 3}, 1},
		{"", []int{3, 3, 3, 3, 3, 4, 4, 2, 2, 1, 1, 3, 3}, 3},
		{"", []int{2, 1, 2, 3, 2, 1, 2, 3, 2, 1, 2, 3, 2}, 2},
	}
	for _, tt := range tests {
		x := CH6{}
		got := x.Majority(tt.seq)
		if got != tt.want {
			t.Errorf("FAIL got: %d, want: %d", got, tt.want)
		}
	}
}

func minimumEditDistanceTest(t *testing.T) {
	ch6_notes_b := [][]int{{0, 1, 2, 3, 4, 5}, {1, 1, 1, 2, 3, 4}, {2, 1, 2, 1, 2, 3}, {3, 2, 2, 2, 1, 2}, {4, 3, 3, 3, 2, 2}}
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  [][]int
	}{
		// {"juha", "AAA", "AAB"},
		{"ch6_notes_b", "abbc", "babba", ch6_notes_b},
	}
	for _, tt := range tests {
		x := CH6{}
		got := x.minimumEditDistance(tt.text1, tt.text2)
		if !ut.Equal2DSlices(got, tt.want) {
			t.Errorf("FAIL got: %v, want: %v", got, tt.want)
		}

		t.Logf("GOT: %v", got)
	}
}

func stringMatchKMPTest(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		// {"Manber exp.","xyxxyxyxyyxyxyxyyxyxyxx", "xyxyyxyxyxx", 13},
		// {"ByteQuest VOD exp.","BABABABABCABABCABAB", "ABABCABAB", 6},
	}
	for _, tt := range tests {
		x := CH6{}
		got := x.stringMatchKMP(tt.text1, tt.text2)
		if got != tt.want {
			t.Errorf("FAIL got: %d, want: %d", got, tt.want)
		} else {
			t.Logf("SUCCESS got: %v", got)
		}

	}
}
func computeNextTest(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []int
	}{
		// {"xyxyy", []int{-1, 0, 0, 1, 2}},
		// {"xyxyyx", []int{-1, 0, 0, 1, 2, 0}},
		// {"xyxyyxyxyxx", []int{-1, 0, 0, 1, 2, 0, 1, 2, 3, 4, 3}},
		// {"ABABCABAB", []int{-1, 0, 0, 1, 2, 0, 1, 2, 3}},
		{"24/hw6/5", "abaababaa", []int{-1, 0, 0, 1, 1, 2, 3, 2, 3}},
	}
	for _, tt := range tests {
		x := CH6{}
		got := x.computeKMPNext(tt.text)
		// got := x.computeKMPNext(tt.text)
		if !slices.Equal(got, tt.want) {
			t.Errorf("FAIL got: %v, want: %v", got, tt.want)
		} else {
			t.Logf("SUCCESS got: %v, want: %v", got, tt.want)
		}

	}
}

func KthSmallestElementTest(t *testing.T) {
	tests := []struct {
		seq  []int
		k    int
		want int
	}{
		{[]int{10, 20, 30}, 1, 10},
		{[]int{10, 20, 30}, 2, 20},
		{[]int{10, 20, 30}, 3, 30},
		{[]int{10, 11, 12, 13}, 3, 12},
		{[]int{10, 11, 12, 13, 14}, 3, 12},
		{[]int{10, 11, 12, 13, 14}, 5, 14},

		{[]int{12, 11, 10}, 1, 10},
		{[]int{12, 11, 10}, 2, 11},
		{[]int{12, 11, 10}, 3, 12},
		{[]int{12, 11, 10, 13}, 1, 10},
		{[]int{14, 13, 12, 11, 10}, 2, 11},
		// {[]int{4, 2, 5, 3, 1}, 4, 0},
		// {[]int{11, 22, 3, 55, 6, 7, 8, 9}, 3, 2},
		// {[]int{11, 22, 3, 55, 6, 7, 8, 9}, 1, 2},
		// {[]int{11, 22, 3, 55, 6, 7, 8, 9}, 6, 1},
		// {[]int{11, 22, 3, 55, 6, 7, 8, 9}, 4, 9},
	}
	for _, tt := range tests {
		x := CH6{t: t, verbose: true}
		seqOrign := slices.Clone(tt.seq)
		got := x.KthSmallestElement(seqOrign, tt.k)
		if got != tt.want {
			if got == -1 {
				t.Errorf("Didn't find the element")
			} else {
				// t.Errorf("Got %d/%d, want %d/%d", gotIndex, tt.seq[gotIndex], tt.want, tt.seq[tt.want])
				t.Errorf("Got %d, want %d", got, tt.want)
			}
		} else {
			n := len(tt.seq)

			ratio := float64(x.iters) / float64(n)
			t.Logf("n: %d, iters: %d, ratio: %.2f", n, x.iters, ratio)
		}

	}
}
func FindMinAndMaxTest(t *testing.T) {
	tests := []struct {
		seq []int
		min int
		max int
	}{
		// {[]int{1, 2, 3}, 1, 3},
		// {[]int{4, 2, 3, 1}, 1, 4},
		// {[]int{4, 2, 5, 3, 1}, 1, 5},
		// {[]int{1, 2, 3, 5, 6, 7, 8, 9}, 1, 9},
		{ut.TestDdata4813, 20, 99986},
	}
	for _, tt := range tests {
		x := CH6{t: t}
		min, max := x.FindMinAndMax(tt.seq)
		if min != tt.min || max != tt.max {
			t.Errorf("Got %d/%d, want %d/%d", min, max, tt.min, tt.max)
		} else {
			n := len(tt.seq)

			ratio := float64(x.iters) / float64(n)
			t.Logf("n: %d, iters: %d, ratio: %.2f", n, x.iters, ratio)
		}

	}
}

func mergeSortTest(t *testing.T) {
	testData := ut.TestData{}
	testData.Init()
	alg2022mid_6 := []int{8, 3, 2, 6, 5, 9, 10, 7, 1, 12, 4, 11}
	tests := []struct {
		seq  []int
		want []int
	}{
		{testData.Test3, ut.GetSorted(testData.Test3)},
		{testData.Test4, ut.GetSorted(testData.Test4)},
		{testData.Test6, ut.GetSorted(testData.Test6)},
		{alg2022mid_6, ut.GetSorted(alg2022mid_6)},
		{testData.Test7, ut.GetSorted(testData.Test7)},
		{testData.Test8, ut.GetSorted(testData.Test8)},
		{testData.Test10, ut.GetSorted(testData.Test10)},
		// {testData.TestDdata4813, ut.GetSorted(testData.TestDdata4813)},
		{[]int{8, 2, 4, 6, 9, 7, 10, 1, 5, 3}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{8, 3, 2, 6, 5, 9, 10, 7, 1, 12, 4, 11}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
	}
	for i, tt := range tests {
		x := CH6{t: t}
		got := x.mergeSortImproved(tt.seq)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FAIL %d: iters: %x, inputs %v, expected %v, got %v", i, x.iters, tt.seq, tt.want, got)
		} else {
			fmt.Printf("SUCCESS merge iters: %d\n", x.iters)
		}
	}
}

func straightRadixSortTest(t *testing.T) {
	tests := []struct {
		seq []int
	}{
		{[]int{170, 45, 75, 90, 802, 24, 2, 66}},
		{[]int{1, 321, 5, 22, 987, 44, 23, 5432}},
	}

	x := CH6{}
	for _, tt := range tests {
		got := x.straightRadixSort(tt.seq)
		t.Logf("%v", got)
	}
}
func interpolationSearchTest(t *testing.T) {
	tests := []struct {
		sortedSeq []int
		target    int
		want      int
	}{
		{[]int{1, 2, 3}, 0, -1},
		{[]int{1, 2, 3}, 1, 0},
		{[]int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, 2, 3}, 4, -1},
		{[]int{1, 20, 33}, 2, -1},
		{[]int{1, 2, 30, 44, 55, 100}, 54, -1},
		{[]int{1, 2, 30, 44, 55, 100}, 55, 4},
		{[]int{1, 2, 30, 44, 55, 100}, 100, 5},
		{[]int{1, 2, 30, 44, 55, 100}, 101, -1},
	}
	for i, tt := range tests {
		x := CH6{t: t, verbose: true}
		got := x.interpolationSearch(tt.sortedSeq, tt.target)
		if got != tt.want {
			t.Errorf("FAIL got: %d, want: %d", got, tt.want)
		} else {
			t.Logf("i: %d got: %v, iters: %d\n", i, got, x.iters)
		}
	}
}
func stutteringSubsequenceTest(t *testing.T) {
	tests := []struct {
		text    string
		pattern string
		want    int
	}{
		// {"abacaba", "ab", 1},
		// {"abacabab", "ab", 2},
		{"banana", "bna", 1},
		{"zzzaacccbbb", "acb", 2},
		{"geeksforgeeks", "gks", 1},
	}
	for i, tt := range tests {
		x := CH6{t: t, verbose: true}
		got := x.stutteringSubsequence(tt.text, tt.pattern)
		if got != tt.want {
			t.Errorf("FAIL got: %d, want: %d", got, tt.want)
		} else {
			t.Logf("i: %d got: %v, iters: %d\n", i, got, x.iters)
		}
	}
}
func specialBinarySearchTest(t *testing.T) {
	tests := []struct {
		seq  []int
		want int
	}{
		// {[]int{1, 2, 3}, -1},
		// {[]int{-1, 0, 2, 10}, 2},
		{[]int{10, 20, 30}, -1},
		{[]int{-1, 0, 1, 3, 5, 6, 7, 8, 9}, 3},
		{[]int{-1, 0, 1, 2, 4, 5, 6, 7, 8, 9}, 4},
	}
	for i, tt := range tests {
		x := CH6{t: t, verbose: true}
		got := x.specialBinarySearch(tt.seq)
		if got != tt.want {
			t.Errorf("FAIL got: %d, want: %d", got, tt.want)
		} else {
			t.Logf("i: %d got: %v, iters: %d\n", i, got, x.iters)
		}
	}
}
func quickSortTest(t *testing.T) {
	tests := []struct {
		seq []int
	}{
		// {[]int{1, 2, 3}},
		// {[]int{2, 3, 1}},
		// {[]int{3, 1, 2}},
		// {[]int{5, 6, 7, 8, 1, 2, 3, 4}},
		// {[]int{2, 3, 4, 5, 6, 7, 8, 1}},
		{[]int{3, 2, 8, 5, 1, 4, 7, 6}},
	}
	for i, tt := range tests {
		want := make([]int, len(tt.seq))
		copy(want, tt.seq)
		slices.Sort(want)
		x := CH6{t: t}
		got := x.quickSort(tt.seq)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FAIL got: %d, want: %d", got, want)
		} else {
			t.Logf("%d got: %v, iters: %d\n", i, got, x.iters)
		}
	}
}
func cyclicBinarySearchTest(t *testing.T) {
	tests := []struct {
		seq  []int
		want int
	}{
		// {[]int{1, 2, 3}, 0},
		// {[]int{2, 3, 1}, 2},
		// {[]int{3, 1, 2}, 1},
		{[]int{5, 6, 7, 8, 1, 2, 3, 4}, 4},
		{[]int{2, 3, 4, 5, 6, 7, 8, 1}, 7},
	}
	for i, tt := range tests {
		x := CH6{t: t}
		got := x.cyclicBinarySearch(tt.seq)
		if got != tt.want {
			t.Errorf("FAIL got: %d, want: %d", got, tt.want)
		} else {
			t.Logf("i: %d SUCCESS got: %d", i, got)
		}
	}
}
