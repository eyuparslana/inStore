package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"inStore/logger"
	"inStore/utils"
)

/*
Home returns all data contained in InMemDB in key value format.
It only works with the GET method. For other methods, http.MethodNotAllowed error is returned.

Example cURL request:
	curl --location --request GET '<API_BASE_URL>:<API_PORT>/'
Example Response:
	{
		"golang": "programming",
		"backend": "developer"
	}
*/
func Home(w http.ResponseWriter, r *http.Request) {
	response := new(ApiResponse)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	defer RecoverPanic(r, response)
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		utils.CreateResponse(w, &InMemDB)
		return
	default:
		err := http.StatusText(http.StatusMethodNotAllowed)
		response.Error = err
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.CreateResponse(w, &response)
		return
	}
}

/*
Set stores the key value in the JSON that comes with the POST request in the inMemDB key value format.
It only works with the POST method. For other methods, http.MethodNotAllowed error is returned.

Example cURL request:
	curl --location --request GET '<API_BASE_URL>:<API_PORT>/set?key=golang&value=programming'

Example Response:
	{
		"result": "The value 'programming' is stored with the key 'golang'."
	}
*/
func Set(w http.ResponseWriter, r *http.Request) {
	response := new(ApiResponse)
	defer RecoverPanic(r, response)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	switch r.Method {
	case http.MethodPost:
		var body LoginRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			logger.Error.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			utils.CreateResponse(w, nil)
			return
		}

		if body.Key == "" {
			response.Error = KeyError
			w.WriteHeader(http.StatusBadRequest)
			utils.CreateResponse(w, &response)
			return
		}

		if body.Value == "" {
			response.Error = ValueError
			w.WriteHeader(http.StatusBadRequest)
			utils.CreateResponse(w, &response)
			return
		}

		StoreData(body.Key, body.Value)
		response.Result = fmt.Sprintf(SetResponsePattern, body.Value, body.Key)
		w.WriteHeader(http.StatusCreated)
		utils.CreateResponse(w, &response)
		return
	default:
		err := http.StatusText(http.StatusMethodNotAllowed)
		response.Error = err
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.CreateResponse(w, &response)
		return
	}
}

/*
Get returns the value corresponding to 'key' in InMemDB.
It only works with the GET method. For other methods, http.MethodNotAllowed error is returned.

Example cURL request:
	curl --location --request GET '<API_BASE_URL>:<API_PORT>/get?key=golang'

Example Response:
	{
		"result": "programming"
	}
*/
func Get(w http.ResponseWriter, r *http.Request) {
	response := new(ApiResponse)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	defer RecoverPanic(r, response)
	switch r.Method {
	case http.MethodGet:
		if key, ok := r.URL.Query()["key"]; ok {
			if value, ok := InMemDB[key[0]]; ok {
				response.Result = value
				w.WriteHeader(http.StatusOK)
				utils.CreateResponse(w, &response)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			err := fmt.Sprintf(KeyNotFoundError, key[0])
			response.Error = err
			utils.CreateResponse(w, &response)
			return
		}
		response.Error = KeyError
		w.WriteHeader(http.StatusBadRequest)
		utils.CreateResponse(w, &response)
		return
	default:
		err := http.StatusText(http.StatusMethodNotAllowed)
		response.Error = err
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.CreateResponse(w, &response)
		return
	}
}

/*
Flush deletes all data in InMemDB.
It only works with the GET method. For other methods, http.MethodNotAllowed error is returned.

Example cURL request:
	curl --location --request GET '<API_BASE_URL>:<API_PORT>/flush'

Example Response:
	{
		"result": "All data has been deleted"
	}
*/
func Flush(w http.ResponseWriter, r *http.Request) {
	response := new(ApiResponse)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	defer RecoverPanic(r, response)
	switch r.Method {
	case http.MethodGet:
		InMemDB = make(map[string]string)
		response.Result = FlushResponse
		w.WriteHeader(http.StatusNoContent)
		return
	default:
		err := http.StatusText(http.StatusMethodNotAllowed)
		response.Error = err
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.CreateResponse(w, &response)
		return
	}
}
