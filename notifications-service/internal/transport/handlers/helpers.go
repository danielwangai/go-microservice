package handlers

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, map[string]string{"error": err.Error()})
}

func respondWithManyErrors(w http.ResponseWriter, code int, errs []string) {
	respondWithJSON(w, code, map[string][]string{"errors": errs})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func convertErrors(errs []error) []string {
	var errList []string
	for i := range errs {
		errList = append(errList, errs[i].Error())
	}

	return errList
}
