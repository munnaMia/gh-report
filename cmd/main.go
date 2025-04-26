package main

import (
	"os"

	services "github.com/munnaMia/gh-report/internal/Services"
	"github.com/munnaMia/gh-report/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		utils.PrintFatal("Usage: gh-report <github-username>")
	}
	
	username:= os.Args[2]

	service := services.GhService{}

	switch os.Args[1] {
	case "gh-report":
		utils.Must(service.Fetch(username))
	}
}
