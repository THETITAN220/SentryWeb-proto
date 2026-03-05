package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func QueryOllama(prompt string) (string, error) {
	reqBody := OllamaRequest{
		Model:  "qwen2.5:1.5b",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, _ := json.Marshal(reqBody)

	resp, err := http.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result OllamaResponse

	json.NewDecoder(resp.Body).Decode(&result)

	return result.Response, nil
}
