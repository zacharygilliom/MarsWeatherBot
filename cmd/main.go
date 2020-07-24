package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	// "github.com/zacharygilliom/MarsWeatherBot/pkg/client"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/twitter"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	/*Create our cron job and execute our daily tweet in the cron job.
	Also will initialize the stream of tweets
	*/

	c := cron.New()
	c.AddFunc("30 20 * * *", twitter.PostTweet)
	c.Start()
	fmt.Println("CronJob Executed")

	demux, client := twitter.GetMessages()
	fmt.Println("Stream Started...")
	// client := client.NewOauth1()
	stream := twitter.Stream(client)
	go demux.HandleChan(stream.Messages)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Stream Stopped...")
	stream.Stop()
	c.Stop()
}
