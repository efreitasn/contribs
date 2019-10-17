# contribs
contribs is a CLI tool to show how many contributions you've done as of the current day.

## Install
You can download one of the binaries available on the [releases page](https://github.com/efreitasn/contribs/releases) or install it using Go v1.13 or higher.

### Installing using go
```bash
go get -u github.com/efreitasn/contribs
$(go env GOPATH)/bin/contribs
```

## API Key
To use contribs, you need to create a new GitHub personal access token with the `read:user` scope to be used as the GitHub API key. Once this is done, you need to add it to the tool using the following:

```bash
contribs set --key YOUR_KEY
```

## How to use
Just run

```bash
contribs
```

### Last year contributions
If you want to see the number of contributions you made in the last year (same data presented in your GitHub profile), run with the `--last-year` flag.

```bash
contribs --last-year
```