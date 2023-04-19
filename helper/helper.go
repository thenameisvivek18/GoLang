package helper

import "strings"

func UserInputValid(firstname string , lastname string , email string , userTicket uint ,remainingTicket uint)(bool,bool,bool){
	isValidName := len(firstname) >=2 && len(lastname) >=2
	isValiEmail :=strings.Contains(email,"@")
	isValidTickets := userTicket > 0 && userTicket <= remainingTicket
	
	return isValidName , isValiEmail ,isValidTickets
}
