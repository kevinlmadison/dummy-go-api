package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// A numResponse is the number to be returned
// swagger:response numResponse
type numResponseWrapper struct {
	// in:body
	Body struct {
		Status string `json:"status"`
		Num    string `json:"num"`
	}
}

// swagger:route GET /api/v1/health health healthEndpoint
// Returns health status of the API
// responses:
// 200: healthResponse
func (a *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	// Health response
	// swagger:parameters healthEndpoint
	response := map[string]string{"status": "OK"}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
		return
	}
}

// swagger:route GET /api/v1/num/{id} num numEndpoint
// Returns the given number
// responses:
// 200: numResponse
func (a *application) numHandler(w http.ResponseWriter, r *http.Request) {
	// swagger:parameters numEndpoint
	id := strings.TrimPrefix(r.URL.Path, "/api/v1/num/")
	fmt.Println("[INF] id: ", id)

	response := map[string]string{"status": "OK", "num": id}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
		return
	}
}
