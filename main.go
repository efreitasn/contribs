package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/efreitasn/contribs/internal/config"
	"github.com/efreitasn/contribs/internal/github"
	"github.com/efreitasn/contribs/internal/logs"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] != "set" {
			logs.Error.Fatalln("while parsing arguments")
		}

		flags := flag.NewFlagSet("set", flag.ExitOnError)
		apiKey := flags.String("key", "", "Set the GitHub API key.")
		flags.Parse(os.Args[2:])

		if *apiKey == "" {
			logs.Error.Fatalln("while parsing arguments")
		}

		c := config.Config{
			GitHubAPIKey: *apiKey,
		}

		err := config.Write(&c)
		if err != nil {
			logs.Error.Fatalln(fmt.Errorf("writing config to file: %w", err))
		}

		return
	}

	config, err := config.Get()
	if err != nil {
		logs.Error.Fatalln(fmt.Errorf("reading config file: %w", err))
	}

	if config == nil {
		logs.Error.Fatalln("config file doesn't exist")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()
	from := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		time.Local,
	)
	to := from.Add((time.Hour * 24) - time.Second)

	numContribs, err := github.GetNumContribs(ctx, config.GitHubAPIKey, from, to)
	if err != nil {
		logs.Error.Fatalln(fmt.Errorf("fetching contribs: %w", err))
	}

	fmt.Println(numContribs)
}
