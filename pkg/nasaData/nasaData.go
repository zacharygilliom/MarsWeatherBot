package nasaData

import (
	"encoding/json"
	"fmt"
	"github.com/zacharygilliom/MarsWeatherBot/configs"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetData() SolDay {
	api_key := configs.GetNasaCredentials()
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
	fmt.Println("JSON string representation")
	var result SolDay
	json.Unmarshal([]byte(body), &result)
	return result
}

func GetSolDay() string {
	start := time.Date(2018, time.Month(11), 26, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	earthDays := now.Sub(start).Hours() / 24
	solDays := earthDays / 1.02749125170
	solDayFloor := math.Floor(solDays)
	stringSolDayFloor := strconv.Itoa(int(solDayFloor))
	return stringSolDayFloor
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
