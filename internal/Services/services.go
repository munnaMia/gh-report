package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/munnaMia/gh-report/internal/entity"
	"github.com/munnaMia/gh-report/internal/utils"
)

type GhService struct{}

func (g GhService) Fetch(userName string) ([]entity.Activity, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", userName)

	response := utils.Must(http.Get(url))
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP status code", response.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	var activity []entity.Activity

	err = json.Unmarshal(body, &activity)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		os.Exit(1)
	}

	return activity, nil
}

func (g GhService) Show(userData []entity.Activity) {
	for _, act := range userData {
		action := ""

		switch act.Type {
		case "PushEvent":
			action = "Pushed commits to"
		case "IssuesEvent":
			action = "Opened an issue in"
		case "IssueCommentEvent":
			action = "Commented on an issue in"
		case "PullRequestEvent":
			action = "Opened a pull request in"
		case "WatchEvent":
			action = "Starred"
		case "ForkEvent":
			action = "Forked"
		default:
			action = act.Type // fallback
		}

		fmt.Printf("- %s %s\n", action, act.Repo.Name)
	}
}
