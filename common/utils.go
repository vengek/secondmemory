package common
import (
	"encoding/json"
	"log"
	"os"
	"net/http"
)

type configuration struct {
	Server string
}

var AppConfig configuration

func initConfig() {
	loadAppConfig()
}

func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()

	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}

	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}

type (
	appError struct {
		Error string `json:"error"`
		Message string `json:"message"`
		HttpStatus string `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
)

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int)  {
	errObject := appError{
		Error: handlerError.Error(),
		Message: message,
		HttpStatus: code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObject}); err == nil {
		w.Write(j)
	}
}