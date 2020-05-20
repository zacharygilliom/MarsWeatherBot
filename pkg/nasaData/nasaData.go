package nasaData

import (
	"fmt"
	"github.com/zacharygilliom/MarsWeatherBot/configs"
	"io/ioutil"
	"log"
	"net/http"
)

func GetData() string {
	api_key := configs.GetCredentials()
	resp, err := http.Get("https://api.nasa.gov/insight_weather/?api_key=" + api_key + "&feedtype=json&ver=1.0")
	if err != nil {
		fmt.Println("Get Request Error : %d", err)
	} else {
		fmt.Println("Successful Get Request")
		fmt.Println(resp)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	string_body := string(body)

	return string_body
}
