package main

import "fmt"

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	secs := g * 3600 / (v2 - v1)
	hour := secs / 3600
	minute := (secs - hour*3600) / 60
	second := secs - hour*3600 - minute*60
	return [3]int{hour, minute, second}
}

func main() {
	fmt.Println(Race(80, 91, 37))
}
