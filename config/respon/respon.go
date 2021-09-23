package respon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR returns a jsonified error response along with a status code.
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Code    int    `json:"code"`
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{
			Code:    400,
			Status:  false,
			Message: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

func FAILED(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	JSON(w, statusCode, struct {
		Code    int         `json:"code"`
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Code:    400,
		Status:  false,
		Message: message,
		Data:    data,
	})
	return
}
