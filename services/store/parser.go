package store

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func checkError(e error, errMsg string) {
	if e != nil {
		fmt.Println(errMsg, e)
		panic(e)
	}
}

func readFile(dataStore *DataStore) {
	filepath := ("./services/data.json")
	file, err := os.Open(filepath)
	checkError(err, "Error opening File: ")
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dataStore)
	checkError(err, "Error decoding JSON: ")
}

func WriteFile(dataStore *DataStore) {
	dataStore.LastUpdateTimeStamp = fmt.Sprint(time.Now().UnixNano())
	filepath := ("./services/data.json")
	file, err := os.Create(filepath)
	checkError(err, "Error opening File: ")
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(dataStore)
	checkError(err, "Error encoding JSON: ")
}
