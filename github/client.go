package github

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/go-github/v32/github"
	"github.com/pkg/errors"
)

type GitHubClient struct {
	service *github.Client
}

func New() GitHubClient {
	return GitHubClient{service: github.NewClient(nil)}
}

// NOTE: 1日以内のPushEventかどうか
func isValidEvent(event string, durationMinutes float64) bool {
	const minutesADay = 1440.0
	return event == "PushEvent" &&  minutesADay >= durationMinutes
}

func (gc *GitHubClient) GetTodaysContributions(userName string) (int, error) {
	events, _, err := gc.service.Activity.ListEventsPerformedByUser(context.Background(), userName, true, nil)
	if _, ok := err.(*github.RateLimitError); ok {
		log.Println("hit late limit")
	}

	count := 0

	for _, event := range events {
		payload := event.GetRawPayload()
		eventDay := event.GetCreatedAt()
		durationMinutes := time.Since(eventDay).Minutes()

		if isValidEvent(event.GetType(), durationMinutes) != true {
			continue
		}

		var pushEventPayload PushEventPayload
		err := json.Unmarshal([]byte(payload), &pushEventPayload)
		if err != nil {
			return 0, errors.Wrapf(err, "invalid event payload")
		}

		count += pushEventPayload.Size
	}

	return count, nil
}

type Commit struct {
	Sha    string `json:"sha"`
	Author struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	Url      string `json:"url"`
}

type PushEventPayload struct {
	PushId       int      `json:"push_id"`
	Size         int      `json:"size"`
	DistinctSize int      `json:"distinct_size"`
	Ref          string   `json:"ref"`
	Head         string   `json:"head"`
	Before       string   `json:"before"`
	Commits      []Commit `json:"commits"`
}
