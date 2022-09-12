package cmd

import (
	"fmt"
	"github.com/mo1ein/makhzan/makhzan"
)

func Execute() {
	fmt.Println("Welcome to makhzan :)")
	client, ctx := makhzan.GhAuth()

	// not forked repos
	allRepos, err := makhzan.RepoList(client, ctx, "")
	if err != nil {
		fmt.Println(err)
	}
	var username string
	fmt.Print("Enter repo username: ")
	fmt.Scan("%s", &username)
	allLangs := makhzan.LangList(client, ctx, allRepos, username)
	makhzan.PieChart(allLangs, "Ratio of languages in github repo")
}
