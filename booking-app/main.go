package main

// 				 go mod init booking-app just run this cmd the file name has to be same as specified during creating the module

import (
	"booking-app/helper"
	"fmt"
	"time"
	//"sync"
	//"strconv"
)

//var wg = sync.WaitGroup{}  // to syncronize the main thread with the no. of extra threads required

const conferenceTickets uint = 50
var conferenceName string = "Coneference" //can also add ':=' and remove 'var' & 'string' & this cannot be used with conatants

var remainingTickets uint= 50  	//unit is for positive no. only int can go -ive 
	//var bookings [50]string //Array (fixed size)
//var bookings = []string  // slice  (a dynamic array) expands automatucally
//var bookings = make([]map[string]string, 0) //a slice of map (we need to specify the initial size of the map)
var bookings = make([]userData,0) // made a slice of structure with initial size of 0


type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main (){
	GreetUser();

	// alternate syntax :
	// var bookings = []string
	// bookings := []string



//	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, confereceName is %T \n",conferenceTickets,remainingTickets,conferenceName)



for {					//infite loop
		
		firstName, lastName, email, userTickets := getUserInput()

		isValidTicketNumbers,isValidName,isValidEmail := helper.ValidateUserInput(firstName ,lastName ,email ,userTickets,remainingTickets)


		if isValidName && isValidTicketNumbers && isValidEmail{
			
			 bookTickets (userTickets,firstName ,lastName ,email)
			// wg.Add(1)  // to add the no. of threads required ro finish the job '1' is the no.
			 /* made the application concurrent that means that execution won't stop 
			 and will continue to execute the next line of code with just adding a 
			 'go' keyword to the function*/ 
			 go sendTicket(userTickets ,firstName ,lastName ,email ) 


			fmt.Printf("The first names of bookings are : %v\n ",getFirstNames())
	
			if remainingTickets == 0{
				fmt.Println("Our conference is booked out. Come back next year")
				break 
			}
		} else {
				if !isValidName{
					fmt.Println("Your entered first name or last name is too short")
				}
				if !isValidEmail {
					fmt.Println("Your entered email address is not valid doesnot contain '@' sign")
				}
				if !isValidTicketNumbers{
					fmt.Println("your entered number of tickets is not vaid")
				}
				
			 }

		//wg.Wait()    //to wait for all the jobs to end before the application to exit
		
	}


	
}

func GreetUser(){
	fmt.Printf("Welcome to our %v booking application \n",conferenceName)
	fmt.Printf("We have a total of %v tickets & %v are still available \n", conferenceTickets,remainingTickets)
	fmt.Println("Get yout tickets here to attend : ")

}

func getFirstNames ()[]string{
	firstNames := []string{}
		for _, booking :=range bookings{			// '_' underscores are used for unused variables  // for each loop
			//var names = strings.Fields(booking)					//string.Fields() seperate till space and gives us 2 elemets in the array
			
			//firstNames = append(firstNames,booking["firstName"])
			firstNames = append(firstNames,booking.firstName)
		}
	return firstNames
}


func getUserInput()(string,string,string,uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
		// ask for the user name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your Email address :")
	fmt.Scan(&email)
		
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName,lastName,email,userTickets
}

func bookTickets (userTickets uint,firstName string,lastName string,email string ){
	remainingTickets = remainingTickets - userTickets
	
	// create a map for the user
	//var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets),10)	//to convert uint into string 

	UserData := userData {
		firstName:firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	
	//bookings[0]= firstName + " "+ lastName			For array
	//bookings = append(bookings,firstName+ " "+lastName)  //for slice append(sliceVar,var that goes into the slice)
	bookings = append(bookings,UserData) //instead of string we r adding a map to our slice
	fmt.Printf("list of bookings is %v \n", bookings)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n ",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v \n ",remainingTickets,conferenceName)
	
}


func sendTicket(userTickets uint,firstName string,lastName string,email string){
	time.Sleep(30 *time.Second)  	// this function stops the execution for the specified time
	ticket :=fmt.Sprintf("%v tickets for %v %v",userTickets,firstName,lastName) //saving formatted string into a variable
	fmt.Println("------------------")
	fmt.Printf("Sending ticket to: \n %v to email address %v \n",ticket,email)
	fmt.Println("------------------")
	//wg.Done() // to remove the thread after it's done 
}



/* for start< 909{
	logic 
} 
this is the while loop in GoLang same thing just thw keyword while is not present 
instead keyword for is there
break,continue,goto work as well 

syantax of goto Eg:
 for start<=100{
	start+=start
	if start == 32 {
		goto locallabel				we will jump out of the loop, loop will not continus to execute after that
	}
	fmt.Println("Somethig")
 }
locallabel: fmt.Println("something 4")


create an label of any var name but add a coln followed by the logic */


/* 
switch variableName{
case "value":
	execute code for booking for this city
case "value2":
	execute code for booking for this city
case "value3" , "value5":			has same code base so we can use this syntax
	execute code for booking for this city
case "value4":
	execute code for booking for this city
default :
	fmt.Println("NO valid value selected")
}
*/