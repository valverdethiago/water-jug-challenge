package restapi

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valverdethiago/water-jug-challenge/code/backend/service"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	shutdownDuration = 5 * time.Second
	port             = 80
)

func TestNewWaterJugServer(t *testing.T) {
	// given
	waterJugService := service.NewWaterJugServiceImpl()
	server := NewWaterJugServer(port, waterJugService)
	server.BindEndpoints()
	defer server.Stop(shutdownDuration)
	payload, err := json.Marshal(Payload{2, 10, 4})
	require.NoError(t, err)
	request, err := http.NewRequest(http.MethodPost, "/", bytes.NewReader(payload))
	require.NoError(t, err)
	recorder := httptest.NewRecorder()

	//when
	server.serverMux.ServeHTTP(recorder, request)

	// then
	require.EqualValues(t, http.StatusOK, recorder.Code)
	states, err := readStatesFromResponse(recorder.Body)
	require.NoError(t, err)
	assert.NotEmpty(t, states)
}

func TestValidationErrors(t *testing.T) {
	tests := []struct {
		name    string
		payload Payload
	}{
		{
			name:    "Should fail with negative numbers",
			payload: Payload{-1, -1, -1},
		},
		{
			name:    "Should fail with Z greater than X and Y",
			payload: Payload{45, 35, 50},
		},
		{
			name:    "Validation should fail when Z is not a multiple of the GCD of x and y",
			payload: Payload{8, 6, 5},
		},
	}

	waterJugService := service.NewWaterJugServiceImpl()
	server := NewWaterJugServer(port, waterJugService)
	server.BindEndpoints()
	defer server.Stop(shutdownDuration)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body, err := json.Marshal(test.payload)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
			require.NoError(t, err)
			recorder := httptest.NewRecorder()
			server.serverMux.ServeHTTP(recorder, request)
			response := recorder.Result()
			require.Equal(t, http.StatusBadRequest, response.StatusCode)
		})
	}
}

func readStatesFromResponse(httpResponse io.Reader) ([]service.State, error) {
	var states []service.State
	decoder := json.NewDecoder(httpResponse)
	err := decoder.Decode(&states)
	if err != nil {
		return nil, err
	}
	return states, nil
}
