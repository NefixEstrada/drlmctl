package software

import (
	"fmt"
	"testing"
)

func TestGitHub(t *testing.T) {
	fmt.Println(gitServiceGitHub.GetRelease("brainupdaters", "drlm", "latest"))
	t.Fail()
}
