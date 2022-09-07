package main

import (
	"github/mo1ein/makhzan/makhzan"
)

func main() {
	client := makhzan.ghAuth()
	// not forked repos
	allRepos := makhzan.reposList(client)
	allLangs := makhzan.langList(client, allRepos)
	makhzan.PieChart(allLangs, "Ratio of languages in github repo")
}
