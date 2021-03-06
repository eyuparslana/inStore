package config

import (
	"inStore/utils"
)

//EXPORT_FILE_PATH: Defines the path of the JSON file.
//API_PORT: Defines which port the API will run on.
var (
	EXPORT_FILE_PATH = utils.GetEnv("EXPORT_FILE_PATH", "./output")
	API_PORT         = utils.GetEnv("API_PORT", "8080")
	RECORD_FREQ      = utils.GetEnv("RECORD_FREQ", "10")
)
