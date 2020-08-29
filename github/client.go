package github

import (
	"context"
	"encoding/json"
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

func isValidEventType(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func isValidEvent(event string, durationMinutes float64) bool {
	const minutesADay = 1440.0
	contributionCountEvents := []string{"PushEvent", "IssuesEvent", "PullRequestEvent"}

	return isValidEventType(contributionCountEvents, event) && minutesADay >= durationMinutes
}

func handlePushEvent(payload json.RawMessage) (int, error) {
	var pushEventPayload *github.PushEvent
	if err := json.Unmarshal([]byte(payload), &pushEventPayload); err != nil {
		return 0, err
	}
	return *pushEventPayload.Size, nil
}

func handleIssuesEvent(payload json.RawMessage) (int, error) {
	var issuesEventPayload *github.IssuesEvent
	if err := json.Unmarshal([]byte(payload), &issuesEventPayload); err != nil {
		return 0, err
	}
	action := issuesEventPayload.GetAction()
	if action == "opened" {
		return 1, nil
	}
	return 0, nil
}

func handlePullRequestEvent(payload json.RawMessage) (int, error) {
	var pullRequestEventPayload *github.PullRequestEvent
	if err := json.Unmarshal([]byte(payload), &pullRequestEventPayload); err != nil {
		return 0, err
	}
	action := pullRequestEventPayload.GetAction()
	if action == "opened" {
		return 1, nil
	}
	return 0, nil
}

func (gc *GitHubClient) GetTodaysContributions(userName string) (int, error) {
	// TODO: add option to handle private contribution in the future
	events, _, err := gc.service.Activity.ListEventsPerformedByUser(context.Background(), userName, true, nil)
	if _, ok := err.(*github.RateLimitError); ok {
		return 0, errors.Wrapf(err, "hit late limit")
	}

	count := 0

	for _, event := range events {
		var contributionNum int
		var err error

		payload := event.GetRawPayload()
		eventType := event.GetType()
		eventDay := event.GetCreatedAt()
		durationMinutes := time.Since(eventDay).Minutes()

		if isValidEvent(eventType, durationMinutes) != true {
			continue
		}

		switch eventType {
		case "PushEvent":
			contributionNum, err = handlePushEvent(payload)
		case "IssuesEvent":
			contributionNum, err = handleIssuesEvent(payload)
		case "PullRequestEvent":
			contributionNum, err = handlePullRequestEvent(payload)
		default:
			if err != nil {
				return 0, errors.Wrapf(err, "invalid event payload")
			}
		}
		count += contributionNum
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

type PullRequestEventPayload struct {
	Action string `json:"action"`
	Number int    `json:"number"`
}
