package mathfunc

import (
	"fmt"
	"math"
)

func Ceil()  {
	var val float64

	fmt.Println("Enter Value:")
	fmt.Scan(&val)
	//up value
	fmt.Printf("Entered Value %v Ceil Value %v\n",val,math.Ceil(val))
}
func Floor()  {
	var val float64

	fmt.Println("Enter Value:")
	fmt.Scan(&val)
	//down value
	fmt.Printf("Entered Value %v Floor Value %v\n",val,math.Floor(val))
}

func ModFunc()  {
	var a, b float64

	fmt.Println("Enter Num 1:")
	fmt.Scan(&a)
	fmt.Println("Enter Num 2:")
	fmt.Scan(&b)

	fmt.Printf("Num1 %v And Num2 %v Mode Value is %v \n",a,b,math.Mod(a,b))
}



func Min(){
	var a, b float64

	fmt.Println("Enter Num 1:")
	fmt.Scan(&a)
	fmt.Println("Enter Num 2:")
	fmt.Scan(&b)

	fmt.Printf("Num 1 %v And Num 2 %v Max Value is %v\n",a,b,math.Min(a,b))
}
func Max(){
	var a, b float64

	fmt.Println("Enter Num 1:")
	fmt.Scan(&a)
	fmt.Println("Enter Num 2:")
	fmt.Scan(&b)

	fmt.Printf("Num 1 %v And Num 2 %v Max Value is %v\n",a,b,math.Max(a,b))
}

func Abs() {
	var val float64

	fmt.Println("Enter Value:")
	fmt.Scan(&val)

	fmt.Printf("your Entered Value %f And Absolute Value %f\n", val, math.Abs(val))
}

func Pow() {
	var num float64
	var p float64

	fmt.Println("Enter Num :")
	fmt.Scan(&num)
	fmt.Println("Enter Power:")
	fmt.Scan(&p)

	fmt.Printf("Entered Num %v or Enterd Power Num %v And Power %v\n", num, p, math.Pow(num, p))

}

func AbsInt() {
	var val int
	fmt.Println("Enter Value:")
	fmt.Scan(&val)

	fmt.Printf("your Entered Value %d of Integer Absolute Value %v\n", val, math.Abs(float64(val)))
}

