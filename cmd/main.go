package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	//"github.com/zacharygilliom/MarsWeatherBot/pkg/client"
	//"github.com/zacharygilliom/MarsWeatherBot/pkg/nasaData"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/twitter"
	//"strconv"
	"time"
)

func main() {

	c := cron.New()
	// Query our data and configure into our struct
	// Pull the Average Temperature data out and format
	// Create our Twitter Client through the go-twitter API.
	// And Post the tweet

	//clientOauth1 := client.NewOauth1()
	//clientOauth2 := client.NewOauth2()
	c.AddFunc("13 0 * * *", twitter.PostTweet)
	c.Start()
	select {}
	fmt.Println("CronJob Executed")
	//time.Sleep(120 * time.Second)

	//tweets := twitter.GetTimeline(clientOauth1)
	//lastTweet := tweets[0]
	//fmt.Println(lastTweet.Text)
	//stream := twitter.Stream(clientOauth1)
	/*for message := range stream.Messages {
		fmt.Println(message)
	}
	*/
	//demux.HandleChan(stream.Messages)
	//stream.Stop()
	twitter.GetMessages()
	time.Sleep(120 * time.Second)

}

//TODO: create function to reply to tweets
