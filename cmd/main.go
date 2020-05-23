package main

import (
	//"encoding/json"
	"fmt"
	"github.com/zacharygilliom/MarsWeatherBot/pkg/nasaData"
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
	// fmt.Println(k)
	/*prettyJSON, err := json.MarshalIndent(k, "", "   ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Println("Prettified JSON Output")
	fmt.Printf("%s\n", string(prettyJSON))
	*/
}
