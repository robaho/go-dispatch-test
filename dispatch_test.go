package main_test

import (
	"testing"
)

type I interface {
	GetValue() int
}

type A struct {
	x int
}

// pointer dispatch is 3x slower
func (a A) GetValue() int {
	return a.x
}

type B struct {
	y int
}

// pointer dispatch is 3x slower
func (b *B) GetValue() int {
	return b.y
}

var gsum int

func BenchmarkIDispatch(b *testing.B) {
	var ifs = [2]I{&A{1}, &A{2}}
	var sum int

	for i := 0; i < b.N; i++ {
		sum += ifs[i%2].GetValue()
	}

	if sum > 0 {
		gsum = sum
	}
}

func BenchmarkIDispatchPtr(b *testing.B) {
	var ifs = [2]I{&B{1}, &B{2}}
	var sum int

	for i := 0; i < b.N; i++ {
		sum += ifs[i%2].GetValue()
	}

	if sum > 0 {
		gsum = sum
	}
}


func BenchmarkDispatch(b *testing.B) {
	var ifs = [2]A{A{1}, A{2}}
	var sum int

	for i := 0; i < b.N; i++ {
		sum += ifs[i%2].GetValue()
	}

	if sum > 0 {
		gsum = sum
	}
}

var bytes [16 * 1024]byte

func BenchmarkArrayParms(b *testing.B) {
	var slice = bytes[:]

	var sum int

	for i := 0; i < b.N; i++ {
		sum += calculateChecksum(slice)
	}

	if sum > 0 {
		gsum = sum
	}
}

func calculateChecksum(slice []byte) int {
	var sum int
	for i, b := range slice {
		sum = sum ^ int(b)
		slice[i] = byte(sum) // modify slice so to avoid loop removal
	}
	return sum
}
