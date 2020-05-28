package main

import (
	"fmt"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/nasaData"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/twitterPost"
)

func main() {
	// Query our data and configure into our struct
	nasadata := nasaData.GetData()
	fmt.Println("GetData configured correctly")
	if nasadata == nil {
		fmt.Println("Error while parsing JSON data in GetData function")
	}

	solDay := nasaData.GetSolDay()
	if solDay == nil {
		fmt.Println("Get Sol Day function returned invalid data")
	}

	// Pull the Average Temperature data out and format
	dailyTemp := GetDayTemp(solDay, nasadata)

	// Create our Twitter Client through the go-twitter API.
	// And Post the tweet
	client := twitterPost.CreateClient()
	twitterPost.PostTweet(client, solDay, dailyTemp)

}

func GetDayTemp(solDay *string, data nasaData.SolDay) string {
	currentDayTemp := data[*solDay].AT.Av
	currentDayTempFormatted := fmt.Sprintf("%.2f", currentDayTemp)
	return currentDayTempFormatted
}

//DONE: Create function in nasadata.go to calculate current day of the week,
// and to convert it to Sol Days.
//DONE: Create function that will tweet out the weather info for the current SOL
// day.
