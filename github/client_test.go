package github_test

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/ia17011/hollein/github"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetTodaysPublicContributions(t *testing.T) {
	// NOTE: recorded request when contribution count was 16
	// if you have the fixtures/github.yaml, this test uses the data without requesting api.github.com
	const recordedContributionCount = 16
	tests := []struct {
		testCase string
		testDate string
		count    int
	}{
		{
			testCase: "fetch todasy GitHub contributions",
			testDate: "Sun, 30 Aug 2020 09:50:25 GMT",
			count: recordedContributionCount,
		},
	}

	// create go-vcr recorder
	r, _ := recorder.New("../fixtures/github")
	defer r.Stop()
	customHTTPClient := &http.Client{
		Transport: r,
	}	

	githubClient := github.New(customHTTPClient)

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			count, _ := githubClient.GetTodaysPublicContributions("ia17011")
			assert.Equal(t, tt.count, count)
		})
	}
}
