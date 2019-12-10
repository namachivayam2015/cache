package util

import (
	"encoding/json"
	"net/http"
)

type CustomException struct {
	msg string
}

type ErrorResponse struct {
	Message string `json:"error"`
}


func (error *CustomException) Error() string {
	return error.msg
}


func BuildErrorResponse(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if errMsg != "" {
		var errData ErrorResponse
		errData.Message = errMsg
		err, _ := json.Marshal(errData)
		w.Write(err)
	}
}
