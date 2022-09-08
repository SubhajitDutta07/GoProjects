package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"

)

func main(){

	os.Setenv("SLACK-BOT-TOKEN","xoxb-3876397852503-4001654540983-crbYeHRrYY1e4WD0AqhFED0R")
	os.Setenv("CHANNEL-ID","C0415UFQQDN")
	api := slack.New(os.Getenv("SLACK-BOT-TOKEN"))  //getting the environment 
	channelArr := []string{os.Getenv("CHANNEL-ID")} 	//getting the the evironment variable and passed it to the string
	fileArr := []string{"git-cheat-sheet-education.pdf"} // file to be uploaded


	for i:=0;i<len(fileArr);i++{		
		params := slack.FileUploadParameters{ 		// getting a parameter 
			Channels : channelArr,
			File : fileArr[i],

		}
		file,err :=api.UploadFile(params)  // uploading the file
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Printf("Name : %s, URL : %s \n",file.Name,file.URL) // and printing the name and URL of the file in the terminal
	}
}