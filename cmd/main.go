package main

import (
	"fmt"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/nasaData"
)

func main() {
	k := nasaData.GetData()
	fmt.Println(k)
}
