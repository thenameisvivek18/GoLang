package strings


import "fmt"

func Ascii() {
	var val string

	fmt.Println("Enter Character:")
	fmt.Scan(&val)

	for i := 0; i < len(val); i++ {
		fmt.Printf("The ASCII value of %c is %d \n", val[i], val[i])
	 }
}
