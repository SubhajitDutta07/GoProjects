package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func main(){
	// setting the environment in the OS 
	os.Setenv("SLACK_BOT_TOKEN","xoxb-3876397852503-4035640003584-C1B6z90lGix3NNnD3uvT4R2w")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A0411HR4S9E-4008941107749-17adce0c7e04b6dde933a6ea34e997881f7974e492275638cc957603a2351a9d")
	// calling the environment variable
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))

	// prints the cmd events
	go printCommandEvents(bot.CommandEvents())

	//the actual things which runs the program

	//cmd which is to be used to run the bot program
	// int the angular brackets there is the parameter
	bot.Command("my yob is <year>",&slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year") // gettig the parameter usinf Param()
			yob , err := strconv.Atoi(year)
			if err != nil{
				fmt.Println("error")
			}
			age := 2022 - yob
			r :=fmt.Sprintf("Age is %d",age)	// to get the format which is to be printed 
			response.Reply(r)		// reply on the slack 
		},
	})

	// to stop your bot 
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//listen receives event from slack
	err := bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}

func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		// prnting the information 
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("##################")
	}

}
