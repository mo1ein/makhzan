package main

import (
	"fmt"
	"makhzan/makhzan"
)

func main() {
    fmt.Println("OK")
	client, ctx := makhzan.GhAuth()
	// not forked repos
    allRepos := makhzan.ReposList(client, ctx)
    allLangs := makhzan.LangList(client, ctx, allRepos, "mo1ein")
    makhzan.PieChart(allLangs, "Ratio of languages in github repo")
}
