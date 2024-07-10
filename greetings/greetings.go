package greetings

import "fmt"

//Hello returns a greeting for the named person

func Hello(pName string) string { 
	//returns a greeting that embeds the name in the message
	message:= fmt.Sprintf("Hi %v. Welcome!", pName)
	return message
} 


