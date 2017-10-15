package pickpeaks_test

import (
	. "github.com/manhtai/gogo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(array []int, posPeaks PosPeaks) {
	var ans = SolPickPeaks(array)
	Expect(ans).To(Equal(posPeaks))
}

var _ = Describe("Example Test Example", func() {
	It("should support finding peaks", func() {
		dotest(
			[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		dotest(
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		)
	})

	It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
		dotest(
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
			PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		)
	})

	It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
		dotest(
			[]int{2, 1, 3, 1, 2, 2, 2, 2, 1},
			PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		dotest(
			[]int{2, 1, 3, 1, 2, 2, 2, 2},
			PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		dotest(
			[]int{2, 1, 3, 2, 2, 2, 2, 5, 6},
			PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		)
	})

	It("should support finding peaks, despite the plateau", func() {
		dotest(
			[]int{2, 1, 3, 2, 2, 2, 2, 1},
			PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		)
	})

	It("should support finding peaks", func() {
		dotest(
			[]int{1, 2, 5, 4, 3, 2, 3, 6, 4, 1, 2, 3, 3, 4, 5, 3, 2, 1, 2, 3, 5, 5, 4, 3},
			PosPeaks{Pos: []int{2, 7, 14, 20}, Peaks: []int{5, 6, 5, 5}},
		)
	})

	It("should return an object with empty arrays if the input is an empty array", func() {
		dotest(
			[]int{},
			PosPeaks{Pos: []int{}, Peaks: []int{}},
		)
	})

	It("should return an object with empty arrays if the input does not contain any peak", func() {
		dotest(
			[]int{1, 1, 1, 1},
			PosPeaks{Pos: []int{}, Peaks: []int{}},
		)
	})
})
