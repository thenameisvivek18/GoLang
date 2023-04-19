package numeric


import "fmt"


func OddEven()  {
	var num int

	fmt.Println("Enter Num:")
	fmt.Scan(&num)

	if num%2==0 {
		fmt.Println("Num Is Even")
	}else{
		fmt.Println("Num Is Odd")
	}

}

func Add(){
	var a uint
	var b uint
	fmt.Println("Enter Num 1:")
	fmt.Scan(&a)
	fmt.Println("Enter Num 2:")
	fmt.Scan(&b)


	fmt.Println("Addition of Two Numbers:",a+b)
}

func Average(){
	var a,b,c int	
	fmt.Println("Enter Num 1:")
	fmt.Scan(&a)
	fmt.Println("Enter Num 2:")
	fmt.Scan(&b)
	fmt.Println("Enter Num 3:")
	fmt.Scan(&c)

	fmt.Println("Average Of The Three Numbers:",a+b+c/3)
	
	
}

func Swap() {

	var a, b int

	fmt.Println("Enter Num 1:")
	fmt.Scan(&a)
	fmt.Println("Enter Num 2:")
	fmt.Scan(&b)
	fmt.Println("Before Swap Number:", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After Swap Number:", a, b)

}
