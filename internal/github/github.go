package github

import (
	"context"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type contribsQuery struct {
	Viewer struct {
		ContributionsCollection struct {
			TotalCommitContributions            int
			TotalIssueContributions             int
			TotalPullRequestContributions       int
			TotalPullRequestReviewContributions int
			TotalRepositoryContributions        int
		} `graphql:"contributionsCollection(from:$from,to:$to)"`
	}
}

// GetNumContribs returns the number of contributions of the user that owns the provided API key.
func GetNumContribs(ctx context.Context, apiKey string, from time.Time, to time.Time) (int, error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiKey},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	query := new(contribsQuery)

	client := githubv4.NewClient(httpClient)

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

	numContribs := query.Viewer.ContributionsCollection.TotalCommitContributions +
		query.Viewer.ContributionsCollection.TotalIssueContributions +
		query.Viewer.ContributionsCollection.TotalPullRequestContributions +
		query.Viewer.ContributionsCollection.TotalPullRequestReviewContributions +
		query.Viewer.ContributionsCollection.TotalRepositoryContributions

	return numContribs, nil
}
