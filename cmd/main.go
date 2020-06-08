package main

import (
	"fmt"
	//"github.com/robfig/cron/v3"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/client"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/nasaData"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/twitter"
	"strconv"
)

func main() {

	// Query our data and configure into our struct
	nasadata := nasaData.GetData()
	fmt.Println("GetData configured correctly")
	if nasadata == nil {
		fmt.Println("Error while parsing JSON data in GetData function")
	}

	solDay := nasaData.GetSolDay()
	fmt.Println(solDay)

	// Pull the Average Temperature data out and format

	// Create our Twitter Client through the go-twitter API.
	// And Post the tweet
	listday := nasaData.GetListDays(nasadata)
	client := client.New()
	tweet := ConfigureTweet(listday, solDay, nasadata)
	fmt.Println(tweet)
	//twitter.NewTweet(client, tweet)
	tweets := twitter.GetTimeline(client)
	lastTweet := tweets[0]
	fmt.Println(lastTweet.Text)

}

func GetDayTemp(solDay int, data nasaData.SolDay) float64 {
	currentDayTemp := data[solDay].AT.Av
	return currentDayTemp
}

func OneDayTweet(day int, data nasaData.SolDay) string {
	currentTemp := GetDayTemp(day, data)
	strDay := strconv.Itoa(day)
	currentTempStr := fmt.Sprintf("%.2f", currentTemp)
	tweet := "Most up to date Sol Day on Mars is: " +
		strDay + "\n" + "The Average Temperature is " + currentTempStr + " degrees celsius"
	return tweet
}

func MultiDayTweet(days *[]int, data nasaData.SolDay) string {
	var tweet string
	tweet = "No New Data. Most Recent Readings:\n"
	for _, day := range *days {
		fmt.Println(day)
		currentTemp := GetDayTemp(day, data)
		fmt.Println(currentTemp)
		currentTempStr := fmt.Sprintf("%.2f", currentTemp)
		strday := strconv.Itoa(day)
		tweet += "Sol Day - " + strday + ": " + currentTempStr + " °C\n"
	}
	return tweet
}

func ConfigureTweet(days *[]int, day int, data nasaData.SolDay) string {
	check := intInSlice(day, days)
	if check == true {
		tweet := OneDayTweet(day, data)
		return tweet
	} else {
		tweet := MultiDayTweet(days, data)
		return tweet
	}
}

func intInSlice(a int, list *[]int) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}
	return false
}

//TODO: create function to reply to tweets
