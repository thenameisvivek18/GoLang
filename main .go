package main

import ("fmt"
	
	"ticketbooking/helper"
	"time"
	"sync"
)

const conferenceTicket = 50
var ConferenceName = "GoConf"	
var remainingTicket uint = 50
var Bookings = make([]UserData,0)

type UserData struct{
	firstname  string
	lastname string
	email string
	NumberOfTickets uint
}

var wg = sync.WaitGroup{}

func main()  {

greetuser()

		
		firstname , lastname ,email , userTicket := getUser()
		isValidName , isValiEmail ,isValidTickets := helper.UserInputValid(firstname,lastname,email,userTicket,remainingTicket)

		if isValidName && isValiEmail && isValidTickets{
			
			BookTicket(userTicket,firstname,lastname,email)
			wg.Add(1)
			go SendTicket(userTicket,firstname,lastname,email)
			/* fmt.Printf("The Whole Bookings %v\n",Bookings)
			 fmt.Printf("The First Index Bookings %v\n",Bookings[0])
			 fmt.Printf("The Type Bookings %T\n",Bookings)
			 fmt.Printf("The length Bookings %v\n",len(Bookings))
			*/
			//call the fumction firstnames

			firstnames:= userFirstNames()
			fmt.Printf("This are all our Bookings %v\n",firstnames)

			var NoTicketsRemainings bool = remainingTicket == 0
			if  NoTicketsRemainings{
				fmt.Printf("our Conference %v Booked Out. Come Out Next Year\n",ConferenceName)
				// break

			}
		}else{
			
			if !isValidName {
				fmt.Println("Youre entered Short length of the FirstName & Lastname , Try Again")
			}
			if !isValiEmail{
				fmt.Println("Youre entered wrong email , Try Again")
			}
			if !isValidTickets{
				fmt.Println("We Have Only 50 Tickets Of Conference , Try Again")
			}
		}
			wg.Wait()
 
	}		



func greetuser(){
	fmt.Println("*******************************************")
	fmt.Printf(" Welcome To Our %v Booking Application \n",ConferenceName)
	
	fmt.Println("*******************************************")
	fmt.Printf("We Have Total %v conference Tickets \n", remainingTicket)
	fmt.Printf("\n")
}

func userFirstNames()[]string{
	firstnames := []string {}
	for _, Booking := range Bookings{
			
		firstnames = append(firstnames, Booking.firstname)
	 }
	
	return firstnames		

	
}



func getUser()(string,string,string,uint){
	var firstname string
		var lastname string
		var email string
		var userTicket uint

		fmt.Println("Enter You First Name :")
		fmt.Scan(&firstname)
		fmt.Println("Enter You Last Name :")
		fmt.Scan(&lastname)
		fmt.Println("Enter You Email :")
		fmt.Scan(&email)
		fmt.Println("Enter Number Of Tikcets You Want :")
		fmt.Scan(&userTicket) 

		return firstname , lastname , email , userTicket
}

func BookTicket( userTicket uint,firstname string , lastname string ,email string){
	remainingTicket = remainingTicket - userTicket

	var userdata = UserData{
		firstname:firstname,
		lastname : lastname,
		email : email,
		NumberOfTickets : userTicket,
	}

	

	Bookings = append(Bookings, userdata)
	
	fmt.Printf("List Of Bookings is %v\n",Bookings)
	fmt.Printf("Thank You %v For Books Your Tickets you have receive confirmation email %v \n",firstname,email)
	fmt.Printf("%v remaining Tickets of Total Tickets %v\n",remainingTicket,conferenceTicket)
}

func SendTicket(userTicket uint,firstname string,lastname string,email string){
	time.Sleep(10*time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v",userTicket,firstname,lastname)
	fmt.Println("----------------------------------------------------------")
	fmt.Printf("Sending Ticket \n %v \n to email address %v\n",ticket,email)
	fmt.Println("----------------------------------------------------------")
	wg.Done()
}