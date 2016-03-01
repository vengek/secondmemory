package controllers

import (
	"encoding/json"
	"net/http"
	"secondmemory/common"
)

func CreateRecord(w http.ResponseWriter, r *http.Request)  {
	var dataResource RecordResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Incorrect Record",
			500,
		)
		return
	}
	record := &dataResource.Data
	//todo save record
	if j, err := json.Marshal(RecordResource{Data: *record}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Some error",
			500,
		)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

func UpdateRecord(w http.ResponseWriter, r *http.Request)  {
	//todo
}