package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodev/jumia-go/src/internal/infrastructure"
)

const (
	ErrWritingResponse = "could not write to response"
)

type M map[string]interface{}

func WriteError(w http.ResponseWriter, error string, errorCode int) {
	WriteJson(w, errorCode, M{"error": error})
}

func WriteJson(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	writeJson(w, data)
}

func writeJson(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		infrastructure.Logger(context.TODO()).Error(fmt.Sprintf("%s: %s", ErrWritingResponse, err.Error()))
	}
}

func GetPathParameter(r *http.Request, name string) string {
	vars := mux.Vars(r)
	param, ok := vars[name]
	if !ok {
		return ""
	}
	return param
}
