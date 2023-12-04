package restapi

import (
	"encoding/json"
	"errors"
	"github.com/valverdethiago/water-jug-challenge/code/backend/service"
	"log"
	"net/http"
)

type Payload struct {
	XCap   int `json:"x_cap"`
	YCap   int `json:"y_cap"`
	Target int `json:"target"`
}

func solveChallengeHandler(service service.WaterJugService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var err error
		if err = validateMethod(request); err != nil {
			handleError(writer, http.StatusMethodNotAllowed, err)
		}
		var payload Payload
		if err = json.NewDecoder(request.Body).Decode(&payload); err != nil {
			handleError(writer, http.StatusBadRequest, err)
		}
		if states, hasSolution, err := service.SolveWaterJugProblem(payload.XCap, payload.YCap, payload.Target); err != nil {
			handleError(writer, http.StatusBadRequest, err)
		} else {
			if !hasSolution {
				writer.WriteHeader(http.StatusNoContent)
				writer.Write([]byte("No solution"))
				return
			}
			if err = json.NewEncoder(writer).Encode(states); err != nil {
				handleError(writer, http.StatusInternalServerError, err)
				return
			}
		}

	}
}

func corsHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		enableCors(&writer)
	}
}

func handleError(writer http.ResponseWriter, status int, err error) {
	response := HttpErrorResponse{Error: err.Error()}
	jsonBody, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte("Error serializing response"))
		log.Printf("Error serializing response: %s \n", err)
	}
	writer.WriteHeader(status)
	_, err = writer.Write(jsonBody)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Printf("error writing HTTP response: %s \n", err)
	}

}

func validateMethod(request *http.Request) error {
	if request.Method != http.MethodPost && request.Method != http.MethodOptions {
		return errors.New("method not supported")
	}
	return nil
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Request-Headers", "*")
}
