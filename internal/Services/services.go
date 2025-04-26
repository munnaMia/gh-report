package services

import "fmt"

type GhService struct{}

func (g GhService) Fetch(userName string) (string, error) {
	url := fmt.Sprint("https://api.github.com/users/%s/events", userName)
	
	return "", nil
}

func (g GhService) Show() {}
