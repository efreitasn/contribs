package github

import (
	"context"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type contribsCollection struct {
	TotalCommitContributions            int
	TotalIssueContributions             int
	TotalPullRequestContributions       int
	TotalPullRequestReviewContributions int
	TotalRepositoryContributions        int
}

type contribsQueryByTime struct {
	Viewer struct {
		ContributionsCollection contribsCollection `graphql:"contributionsCollection(from:$from,to:$to)"`
	}
}

type contribsQueryLastYear struct {
	Viewer struct {
		ContributionsCollection contribsCollection
	}
}

// NewClient creates a new githubv4 client.
func NewClient(apiKey string) *githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiKey},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	return githubv4.NewClient(httpClient)
}

// GetNumContribsByTime returns the number of contributions made by the user in the current day.
func GetNumContribsByTime(ctx context.Context, client *githubv4.Client, from time.Time, to time.Time) (int, error) {
	query := new(contribsQueryByTime)

	err := client.Query(ctx, &query, map[string]interface{}{
		"from": githubv4.DateTime(struct {
			time.Time
		}{from}),
		"to": githubv4.DateTime(struct {
			time.Time
		}{to}),
	})
	if err != nil {
		return 0, err
	}

	numContribs := sumContribs(query.Viewer.ContributionsCollection)

	return numContribs, nil
}

// GetNumContribsLastYear returns the number of contributions made by the user in the last year.
func GetNumContribsLastYear(ctx context.Context, client *githubv4.Client) (int, error) {
	query := new(contribsQueryLastYear)

	err := client.Query(ctx, &query, nil)
	if err != nil {
		return 0, err
	}

	numContribs := sumContribs(query.Viewer.ContributionsCollection)

	return numContribs, nil
}

func sumContribs(cCollection contribsCollection) int {
	return cCollection.TotalCommitContributions +
		cCollection.TotalIssueContributions +
		cCollection.TotalPullRequestContributions +
		cCollection.TotalPullRequestReviewContributions +
		cCollection.TotalRepositoryContributions
}
