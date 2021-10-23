package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"inStore/config"
	"inStore/handler"
	"inStore/logger"
)

// JsonFileNameFormat: Defines the file name format to be exported.
// Duration: Defines how often the export runs, in minutes.
const (
	JsonFileNameFormat = "%d_inStoreData.json"
)

//CheckExistingData runs when the application is up. If there is any exported data,
//it imports this data into handler.inMemDB.
func CheckExistingData() {
	files, err := ioutil.ReadDir(config.EXPORT_FILE_PATH)
	if err != nil {
		logger.Info.Println(err)
		return
	}

	for _, file := range files {
		path := filepath.Join(config.EXPORT_FILE_PATH, file.Name())
		f, err := os.ReadFile(path)
		if err != nil {
			logger.Fatal.Println(err)
		}

		err = json.Unmarshal(f, &handler.InMemDB)
		if err != nil {
			logger.Fatal.Println(err)
		}
	}
}

//removeFileInDirectory deletes the file with old data.
func removeFileInDirectory() {
	files, err := ioutil.ReadDir(config.EXPORT_FILE_PATH)
	if err != nil {
		logger.Info.Println(err)
		return
	}

	for _, file := range files {
		path := filepath.Join(config.EXPORT_FILE_PATH, file.Name())
		err := os.Remove(path)
		if err != nil {
			logger.Error.Println(err)
		}
	}
}

// saveData will run every specified minutes and export all data from inMemDB to specified json file.
func saveData(ticker *time.Ticker, quit chan struct{}) {
	for {
		select {
		case <-ticker.C:
			if _, err := os.Stat(config.EXPORT_FILE_PATH); os.IsNotExist(err) {
				err := os.Mkdir(config.EXPORT_FILE_PATH, os.ModePerm)
				if err != nil {
					logger.Info.Println(err)
				}
			} else {
				removeFileInDirectory()
			}

			dataBytes, err := json.Marshal(handler.InMemDB)
			if err != nil {
				logger.Fatal.Println(err)
			}

			fileName := filepath.Join(config.EXPORT_FILE_PATH, fmt.Sprintf(JsonFileNameFormat, time.Now().Unix()))
			err = ioutil.WriteFile(fileName, dataBytes, 0777)
			if err != nil {
				logger.Fatal.Println(err)
			} else {
				logger.Info.Printf("All data added to file: %s", fileName)
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}
}

//StartTask starts the tasks that will run when the application is up.
func StartTask() {
	duration, err := strconv.Atoi(config.RECORD_FREQ)
	if err != nil {
		duration = 10
	}
	ticker := time.NewTicker(time.Duration(duration) * time.Minute)
	quit := make(chan struct{})
	go saveData(ticker, quit)
}
