package store

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type UrlData struct {
	LongUrl   string `json:"longUrl"`
	CreatedAt string `json:"createdAt"`
}

type DataStore struct {
	DataMap             map[string]UrlData `json:"data"`
	LastUpdateTimeStamp string             `json:"lastUpdateTimeStamp"`
}

var _dataStoreInstance *DataStore

func init() {
	initStoreInstance()
}

func initStoreInstance() {
	lock.Lock()
	defer lock.Unlock()

	if _dataStoreInstance == nil {
		fmt.Println("Creating single instance now.")
		_dataStoreInstance = &DataStore{
			DataMap: make(map[string]UrlData),
		}
		readFile(_dataStoreInstance)
	} else {
		fmt.Println("Single instance already created.")
	}
}

func GetDataStoreInstance() *DataStore {
	if _dataStoreInstance == nil {
		initStoreInstance()
	}
	return _dataStoreInstance
}
