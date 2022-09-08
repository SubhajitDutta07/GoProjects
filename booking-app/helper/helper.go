package helper

import "strings"


func ValidateUserInput(firstName string,lastName string,email string,userTickets uint,remainingTickets uint) (bool,bool,bool){
	isValidName := len(firstName) >= 2 && len(lastName) >=2
	isValidEmail:= strings.Contains(email,"@")				//strings.Contains(variable, character to check) This func checks whether the character is present in the string or not
	isValidTicketNumbers := userTickets > 0 && userTickets <= remainingTickets
	return isValidTicketNumbers,isValidName,isValidEmail			// in go we can retuen any number of values also seethe return type it will be same as the return types
		
}

// to export a function we just capitalizethe name of the function just the 1st letter
// we can also export variables ny the same method