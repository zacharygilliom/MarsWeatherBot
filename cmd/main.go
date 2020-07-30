package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/zacharygilliom/MarsWeatherBot/internal/twitter"
	"github.com/zacharygilliom/MarsWeatherBot/internal/visuals"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	/*Create our cron job and execute our daily tweet in the cron job.
	Also will initialize the stream of tweets
	*/
	file, err := os.OpenFile("/home/zacharygilliom/goProjects/MarsWeatherBot/logs/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	x, y := visuals.GetAxesValues()
	visuals.PlotGraph(x, y)

	c := cron.New()
	c.AddFunc("30 20 * * *", twitter.PostTweet)
	c.Start()

	demux, client := twitter.GetMessages()
	fmt.Println("Stream Started...")
	stream := twitter.Stream(client)
	go demux.HandleChan(stream.Messages)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Stream Stopped...")
	stream.Stop()
	c.Stop()
}
