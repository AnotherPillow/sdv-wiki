package main

import (
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "51013"
	}
	return port
}
