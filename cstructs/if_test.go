package cstructs_test

import (
	"fmt"
	"testing"

	"example.com/m/cstructs"
)

func ExampleIfTrue() {
	count := 0
	if true {
		count = count + 1
	}
	fmt.Println(count)
	// output: 1
}
func TestIfTrue(t *testing.T) {
	got := cstructs.IfTrue()
	want := 1

	assertCorrectMessageInt(got, want, t)
}
func ExampleIfFalse() {
	count := 0
	if false {
		count = count + 1
	}
	fmt.Println(count)
	// output: 0
}
func TestIfFalse(t *testing.T) {
	got := cstructs.IfFalse()
	want := 0

	assertCorrectMessageInt(got, want, t)
}
func ExampleIfInitializationTrue() {
	count := 0
	if one := 1; true {
		count = count + one
	}
	fmt.Println(count)
	// output: 1
}
func TestIfInitializationTrue(t *testing.T) {
	got := cstructs.IfInitializationTrue()
	want := 1

	assertCorrectMessageInt(got, want, t)
}
func ExampleIfInitializationFalse() {
	count := 0
	if one := 1; false {
		count = count + one
	}
	fmt.Println(count)
	// output: 0
}
func TestIfInitializationFalse(t *testing.T) {
	got := cstructs.IfInitializationFalse()
	want := 0

	assertCorrectMessageInt(got, want, t)
}
