package fizzbuzz

import (
	"fmt"
	"strconv"
)

type FizzBuzz struct {
	index int
	length int
}

func NewFizzBuzz() *FizzBuzz {
	f := &FizzBuzz{index: 1, length: 50}
	return f
}

func Run(fizzBuzz FizzBuzz) []string {
	var arr []string

	for i := fizzBuzz.index; i < fizzBuzz.length; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if i % 3 == 0 {
			arr = append(arr, "Fizz")
		} else if i % 5 == 0 {
			arr = append(arr, "Buzz")
		} else {
			arr = append(arr, strconv.Itoa(i))
		}
	}
	return arr
}

func main() {
	fizzBuzz := NewFizzBuzz()
	arr := Run(*fizzBuzz)

	for i, str := range arr {
		fmt.Println(i, string(str))
	}
}

