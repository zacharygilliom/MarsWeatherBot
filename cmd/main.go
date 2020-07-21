package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/twitter"
)

func main() {

	/*Create our cron job and execute our daily tweet in the cron job.
	Also will initialize the stream of tweets
	*/

	c := cron.New()
	c.AddFunc("30 20 * * *", twitter.PostTweet)
	c.Start()
	fmt.Println("CronJob Executed")

	twitter.GetMessages()

}
