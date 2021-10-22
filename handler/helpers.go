package handler

import (
	"net/http"

	"inStore/logger"
)

//RecoverPanic allows the application to log the error and continue working in panic situations.
func RecoverPanic(r *http.Request, response *ApiResponse) {
	if rec := recover(); rec != nil {
		logger.Error.Printf("[RECOVERED PANIC] %+v", rec)
	}
	logger.Info.Printf("Received %s %s [Rsp: {%+v}]", r.Method, r.URL.Path, response)
}

//StoreData allows data to be written to InMemDB.
func StoreData(key, value string) {
	InMemDB[key] = value
}
