package handler

import (
	"encoding/json"
	"net/http"
)

func JSONHandler(js []byte, fallback http.Handler) (http.HandlerFunc, error) {
	json, err := parseJSON(js)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(json)
	return MapHandler(pathMap, fallback), nil
}

func parseJSON(js []byte) ([]urlPair, error) {
	var urlPairs []urlPair
	err := json.Unmarshal(js, &urlPairs)
	if err != nil {
		return nil, err
	}
	return urlPairs, nil
}
