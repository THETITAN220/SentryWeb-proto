package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-METHODS", "POST,OPTIONS")
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}
	var page PageRequest

	json.NewDecoder(r.Body).Decode(&page)

	results, err := AnalyzePage(page.Data)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type:", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"analysis": results,
	})

	reqLog, _ := json.MarshalIndent(page.Data, "", "  ")
	log.Printf("Request: %v", string(reqLog))
	log.Printf("Response: %v", results)
}

func main() {

	log.Println("The backend server is starting... ")
	http.HandleFunc("/analyze", analyzeHandler)
	http.ListenAndServe(":8000", nil)

}
