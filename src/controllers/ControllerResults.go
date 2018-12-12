package controllers

import ()

const (
	// keep adding codes here.
	Status  string = "status"
	Failed  string = "failed"
	Success string = "success"
)

// keep adding for different codes here.

var (
	HttpResults = map[int]interface{}{
		400: map[string]string{Status: Failed, "message": "request format is not as required."},
		404: map[string]string{Status: Success, "message": "requested resource was not found with system."},
		200: map[string]string{Status: Success},
		500: map[string]string{Status: Failed, "message": "uknown error occurred"},
	}
)

func GetHttpResult(httpCode int, debugMessage string) interface{} {
	result := HttpResults[httpCode].(map[string]string) // this is the way casting happens in golang
	if debugMessage != "" {
		result["debugMessage"] = debugMessage
	}
	return result
}
