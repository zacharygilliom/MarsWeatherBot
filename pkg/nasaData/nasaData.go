package nasaData

import (
	"encoding/json"
	"fmt"
	"github.com/zacharygilliom/MarsWeatherBot/configs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetData() SolDay {
	api_key := configs.GetCredentials()
	resp, err := http.Get("https://api.nasa.gov/insight_weather/?api_key=" + api_key + "&feedtype=json&ver=1.0")
	if err != nil {
		fmt.Println("Get Request Error : %d", err)
	} else {
		fmt.Println("Successful Get Request")
		//fmt.Println(resp)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//string_body := string(body)
	fmt.Println("JSON string representation")
	//jsonData := []byte(body)
	var result SolDay
	json.Unmarshal([]byte(body), &result)

	return result
}

func ParseJsonFile() SolDay {
	jsonFile, err := os.Open("nasa.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("nasa.json file opened successfully")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result SolDay
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
