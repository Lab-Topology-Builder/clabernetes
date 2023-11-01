package util

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	claberneteserrors "github.com/srl-labs/clabernetes/errors"
)

// IsURL returns true if the given path string starts with http or https.
func IsURL(path string) bool {
	return strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://")
}

// WriteHTTPContentsFromPath writes content at the http path `path` into the writer w.
func WriteHTTPContentsFromPath(path string, w io.Writer) error {
	resp, err := http.Get(path) //nolint:noctx,gosec
	if err != nil {
		return err
	}

	defer resp.Body.Close() //nolint

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"%w: non 200 status attempting to load content at '%s', status code: %d",
			claberneteserrors.ErrUtil,
			path,
			resp.StatusCode,
		)
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// GitHubNormalToRawLink try to convert a "normal" github link to a raw link.
func GitHubNormalToRawLink(path string) string {
	if !strings.Contains(path, "github.com") {
		return path
	}

	if strings.Contains(path, "githubusercontent") {
		return path
	}

	p := regexp.MustCompile(
		`(?mi)https?:\/\/(?:www\.)?github\.com\/(?P<GroupRepo>.*?\/.*?)\/(?:(blob)|(tree))(?P<Path>.*)`, //nolint:lll
	)

	paramsMap := RegexStringSubMatchToMap(p, path)

	return fmt.Sprintf(
		"https://raw.githubusercontent.com/%s/%s", paramsMap["GroupRepo"], paramsMap["Path"],
	)
}

// GitHubGroupAndRepoFromURL attempts to return the GitHub group/user and repository from a GitHub
// URL.
func GitHubGroupAndRepoFromURL(path string) (group, repo string) {
	if !strings.Contains(path, "github.com") {
		return group, repo
	}

	p := regexp.MustCompile(
		`(?mi).*.\.com/(?P<Group>.*?)/(?P<Repo>.*?)/`,
	)

	paramsMap := RegexStringSubMatchToMap(p, path)

	return paramsMap["Group"], paramsMap["Repo"]
}
