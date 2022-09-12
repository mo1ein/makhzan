package makhzan_test

import (
	"context"
	"github.com/google/go-github/v47/github"
	"github.com/mo1ein/makhzan/makhzan"
	"reflect"
	"testing"
)

func TestRepoList(t *testing.T) {
	ctx := context.Background()
	c := github.NewClient(nil)
	repos, err := makhzan.RepoList(c, ctx, "mo1ein")
	if err != nil {
		t.Errorf("RepoList returned error: %v", err)
	}

	expected := []string{
		"AIC-2021",
		"Codeforces-status",
		"Game-of-Life",
		"Golchin",
		"makhzan",
		"mo1ein",
		"My-CP-Submissions",
		"My-dotfiles",
		"RenameIt",
		"spline-drawer-danaxa",
		"TelegramArchive",
	}
	available := removeEmptyString(repos, "")
	if !reflect.DeepEqual(available, expected) {
		t.Errorf("RepoList returned error: %v", err)
	}
}

func removeEmptyString(s []string, r string) []string {
	var res []string
	for _, v := range s {
		if v != r {
			res = append(res, v)
		}
	}
	return res
}
