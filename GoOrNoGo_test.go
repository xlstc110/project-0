//Project-0 func test.
package main

import "testing"

//func find will loop through the slice and check if the number has previously picked and return a boolean result.
func Test_find(t *testing.T) {
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	find(x, 3)
}

func Test_record(t *testing.T) {
	record(1, 33333, 66666)
}

func Test_ScoreCheck(t *testing.T) {
	ScoreCheck()
}

func Test_load(t *testing.T) {
	load("Robbie")
}

func Benchmark_find(b *testing.B) {
	y := []int{}
	for i := 0; i < 8000000; i++ {
		y[i] = i
	}
	find(y, 5000000)
}

func Benchmark_timesUp(b *testing.B) {
	timesUp(30000, 40000)
}
