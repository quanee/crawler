package engine

import (
	"crawler/fetcher"
	"log"
)

func worker(req Request) (Result, error) {
	body, err := fetcher.Fetch(req.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", req.URL, err)
		return Result{}, err
	}
	return req.Callback(body, req.URL), err
}
