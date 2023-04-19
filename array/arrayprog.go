package array

import "fmt"

func DynamicArray() {
	var n int

	fmt.Println("Enter Length Of the Array")
	fmt.Scan(&n)

	arr := []int{}
	var s int
	for i := 0; i < n; i++ {
		fmt.Println("Enter Array Elements:")
		fmt.Scan(&s)
		arr = append(arr, s)
	}

	fmt.Println("Array Of Elements:")
	for i := 0; i < n; i++ {
		fmt.Printf("Number Of Elements:%d\n", arr[i])
	}

	var sort string
	fmt.Println("If You Want To Change Your Array Format ? ")
	fmt.Print("Ascending Order To Type : a\nDescending Order To Type : d\n")
	fmt.Scan(&sort)
	if  sort=="a" {
		var temp int
	for i:=0; i<n ; i++{
		for j:=i; j<n ; j++{
				if arr[j] < arr[i] {
					temp = arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}	
	}
	fmt.Println("Sorted Array Of Elements:")
	for i := 0; i < n; i++ {
		fmt.Printf("Number Of Elements:%d\n", arr[i])
	}	
	}else if sort == "d"{
		var temp int
		for i:=0; i<n ; i++{
			for j:=0; j<n-i-1; j++{
					if arr[j] < arr[i] {
						temp = arr[i]
						arr[i] = arr[j]
						arr[j] = temp
					}
				}	
		}
		fmt.Println("Sorted Array Of Elements:")
		for i := 0; i < n; i++ {
			fmt.Printf("Number Of Elements:%d\n", arr[i])
		}
	}else{
		fmt.Println("Thank You")
		
	}

	

}
