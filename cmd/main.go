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
		fmt.Println(nasadata)
	}

	solDay := nasaData.GetSolDay()
	if solDay == nil {
		fmt.Println("Get Sol Day function returned invalid data")
	}

	// Pull the Average Temperature data out and format
	dailyTemp := GetDayTemp(solDay, nasadata)
	dailyTempFormatted := fmt.Sprintf("%.2f", dailyTemp)

	// Create our Twitter Client through the go-twitter API.
	// And Post the tweet
	client := twitterPost.CreateClient()
	twitterPost.PostTweet(client, solDay, dailyTempFormatted)

}

func GetDayTemp(solDay *string, data nasaData.SolDay) float64 {
	currentDayTemp := data[*solDay].AT.Av
	return currentDayTemp
}

//DONE: Create function in nasadata.go to calculate current day of the week,
// and to convert it to Sol Days.
//DONE: Create function that will tweet out the weather info for the current SOL
// day.
