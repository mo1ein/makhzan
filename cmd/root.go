package cmd

import (
	"fmt"
	"github.com/mo1ein/makhzan/makhzan"
)

func Execute() {
	fmt.Println("Welcome to makhzan :)")
	client, ctx := makhzan.GhAuth()

	// not forked repos
	allRepos, err := makhzan.RepoList(client, ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(allRepos)
	allLangs := makhzan.LangList(client, ctx, allRepos, "mo1ein")
	makhzan.PieChart(allLangs, "Ratio of languages in github repo")
}
