package findunique

// FindUniq is solution for http://www.codewars.com/kata/585d7d5adb20cf33cb000235
func FindUniq(arr []float32) float32 {
	m := make(map[float32]int)
	var n float32
	for _, num := range arr {
		m[num]++
	}
	for key, value := range m {
		if value == 1 {
			n = key
			break
		}
	}
	return n
}
