package pickpeaks

// http://www.codewars.com/kata/5279f6fe5ab7f447890006a7

// PosPeaks datastruture to return
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// SolPickPeaks is function to solve Pick Peaks
// TODO: Must we export this to use by test?
func SolPickPeaks(array []int) PosPeaks {
	posPeaks := PosPeaks{[]int{}, []int{}}
	candidate := 0
	for i := 1; i < len(array); i++ {
		if array[i-1] < array[i] {
			candidate = i
		} else if array[i-1] > array[i] && candidate > 0 {
			posPeaks.Pos = append(posPeaks.Pos, candidate)
			posPeaks.Peaks = append(posPeaks.Peaks, array[candidate])
			candidate = 0
		}
	}
	return posPeaks
}
