// SPDX-License-Identifier: AGPL-3.0-only

package software

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// gitService is a hosting Git service. It's used to query it's API for getting latest releases, etc...
type gitService interface {
	// CloneURL returns the clone URL of a repository
	CloneURL(user, repo string) string
	// GetRelease returns the latest release of a repo (it's tag)
	GetRelease(user, repo, release string) (string, error)
}

var gitServiceGitHub = &gitHub{
	Scheme: "https",
	Host:   "api.github.com",
}

type gitHub url.URL

// CloneURL returns the clone URL of a repository
func (g *gitHub) CloneURL(user, repo string) string {
	return fmt.Sprintf("https://github.com/%s/%s", user, repo)
}

// GetRelease returns the release tag of a release of a repo
func (g *gitHub) GetRelease(user, repo, release string) (string, error) {
	type response struct {
		TagName string `json:"tag_name,omitempty"`
	}

	var req = &url.URL{
		Scheme: g.Scheme,
		Host:   g.Host,
	}
	req.Path = path.Join(g.Path, "repos", user, repo, "releases", release)

	rsp, err := http.Get(req.String())
	if err != nil {
		return "", fmt.Errorf("error getting the release tag: %v", err)
	}
	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error getting the release tag: %v", rsp.StatusCode)
	}

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the response body: %v", err)
	}

	var r response
	if err = json.Unmarshal(b, &r); err != nil {
		return "", fmt.Errorf("error unmarshaling the API response: %v", err)
	}

	return r.TagName, nil
}
