package cmd

import (
	"fmt"
	"github.com/mo1ein/makhzan"
)

func Execute() {
	fmt.Println("Welcome to makhzan :)\n")
	client, ctx := makhzan.GhAuth()
	// not forked repos
	allRepos := makhzan.ReposList(client, ctx)
	allLangs := makhzan.LangList(client, ctx, allRepos, "mo1ein")
	makhzan.PieChart(allLangs, "Ratio of languages in github repo")
}
