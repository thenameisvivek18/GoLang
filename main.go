package main

import (
	"fmt"
	"practice/mathfunc"
	"practice/numeric"
	"practice/strings"
	"practice/array"
)

func main() {

	fmt.Println("**************************************")
	fmt.Println("Hello Welcome To Go Lang Basic Program")
	fmt.Println("**************************************")
	fmt.Println("Enter 1 For the Addition Of The TWo Numbers")
	fmt.Println("Enter 2 For the Average Of The Three Numbers")
	fmt.Println("Enter 3 For the Swap Of The TWo Numbers")
	fmt.Println("Enter 4 For the Ascii Value Of The Character")
	fmt.Println("Enter 5 To Find Absolute Values Of Float The Number")
	fmt.Println("Enter 6 To Find Absolute Values Of Int The Number")
	fmt.Println("Enter 7 To Find Power Of The Number")
	fmt.Println("Enter 8 To Find Max Of The Two Number")
	fmt.Println("Enter 9 To Find Min Of The Two Number")
	fmt.Println("Enter 10 To Find Ceil Of The Two Number")
	fmt.Println("Enter 11 To Find Floor Of The Two Number")
	fmt.Println("Enter 12 To Find Mod Of The Two Number")
	fmt.Println("Enter 13 To Find Number is Even Or Odd")
	fmt.Println("Enter 14 To Generate Dynamic Array")
	fmt.Println("")

	var option int

	fmt.Scan(&option)

	switch option {
	case 1:
		numeric.Add()
	case 2:
		numeric.Average()
	case 3:
		numeric.Swap()
	case 4:
		strings.Ascii()
	case 5:
		mathfunc.Abs()
	case 6:
		mathfunc.AbsInt()
	case 7:
		mathfunc.Pow()
	case 8:
		mathfunc.Max()
	case 9:
		mathfunc.Min()
	case 10:
		mathfunc.Ceil()
	case 11:
		mathfunc.Floor()
	case 12:
		mathfunc.ModFunc()
	case 13:
		numeric.OddEven()
	case 14:
		array.DynamicArray()
	default:
		fmt.Println("Not Matched Option")

	}

}
