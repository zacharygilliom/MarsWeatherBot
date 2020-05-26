package main

import (
	//"encoding/json"
	"fmt"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/nasaData"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/twitterPost"
	//"log"
)

func main() {
	// k := nasaData.ParseJson()
	/*for key, value := range k {
		fmt.Println(key)
		if key == "sol_keys" {
			days_sol = value
		}
		fmt.Println(value)
	}
	*/
	k := nasaData.GetData()
	fmt.Println("GetData configured correctly")
	if k == nil {
		fmt.Println(k)
	}
	// fmt.Println(k["522"].AT.Av)
	dailyTemp := k["530"].AT.Av
	tweet := fmt.Sprintf("%f", dailyTemp)

	client := twitterPost.CreateClient()
	twitterPost.PostTweet(client, tweet)
}

/* func PostDailyWeather(client *twitter.Client, tweet string) {
	twitterPost.PostTweet(client, tweet)
}
*/
