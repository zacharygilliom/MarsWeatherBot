package nasaData

import (
	"encoding/json"
	"fmt"
	"github.com/zacharygilliom/MarsWeatherBot/configs"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sort"
	"time"
)

// Request Data from Nasa API and unmarshal it into our nasaData struct.
func GetData() SolDay {
	api_key := configs.GetNasaCredentials()
	resp, err := http.Get("https://api.nasa.gov/insight_weather/?api_key=" +
		api_key + "&feedtype=json&ver=1.0")
	if err != nil {
		fmt.Println("Get Request Error :")
		fmt.Println(err)
	} else {
		fmt.Println("Successful Get Request with NasaData API")
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
	var result SolDay
	json.Unmarshal([]byte(body), &result)
	fmt.Println(result)
	return result

}

// Get the list of days from our API request.
func GetListDays(days SolDay) *[]int {
	var listday []int
	for day, _ := range days {
		listday = append(listday, day)
	}
	sort.Ints(listday)
	return &listday

}

// Get current mars day based on current earth day.
func GetSolDay() int {
	location, err := time.LoadLocation("America/New_York")
	start := time.Date(2018, time.Month(11), 26, 0, 0, 0, 0, location)
	if err != nil {
		fmt.Println(err)
	}
	now := time.Now()
	earthDays := now.Sub(start).Hours() / 24
	solDays := earthDays / 1.02749125170
	solDayFloor := (math.Floor(solDays)) - 1
	solDayFloorInt := int(solDayFloor)
	return solDayFloorInt
}

// Using the current sol day from the GetSolDay function, we will pull the information from our NASA API Request for this day
func GetDayTemp(solDay int, data SolDay) float64 {
	currentDayTemp := data[solDay].AT.Av
	return currentDayTemp
}

/* Function below is to allow the user to download a copy of the json data and then use that data repeatedly instead of making numerous API calls.
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
*/
