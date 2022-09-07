package makhzan

import (
	"context"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
	"os"
)

// *charts.Pie
func PieChart(d map[string]int, t string) {
	destinations := make([]opts.PieData, len(d))
	i := 0
	for k, v := range d {
		destinations[i] = opts.PieData{
			Name:  k,
			Value: v,
			ItemStyle: &opts.ItemStyle{
				Color: colors[k],
			},
		}
		i++
	}
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: t,
		TitleStyle: &opts.TextStyle{
			FontFamily: "sens-serif",
		},
		Left: "center",
	}),
	)

	pie.SetGlobalOptions(charts.WithTooltipOpts(opts.Tooltip{
		Show: true,
		// item or axis?
		Trigger: "item",
		// mousemove or click or both
		TriggerOn: "mousemove",
	}))

	pie.SetGlobalOptions(charts.WithLegendOpts(opts.Legend{
		Show:   true,
		Orient: "horizontal",
		Top:    "bottom",
	}))

	// TODO: add Toolbox for save image
	//TODO: change these like smooth add series
	pie.AddSeries("Languages", destinations)
	pie.SetSeriesOptions(charts.WithLabelOpts(
		opts.Label{
			Show:      true,
			Formatter: "{b}: {d}%",
		}),
	)

	f, _ := os.Create("pie.html")
	pie.Render(f)
	// return pie
}

func ghAuth() *github.Client {
	ctx := context.Background()
	//TODO: input gh token
	var string ghToken
	fmt.Scanf("%s", &ghToken)
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	var string username
	fmt.Scanf("%s", &username)

	//TODO: this for use other funcs
	// list all repositories for the authenticated user
	return client
}

func reposList(c *github.Client) []string {
	// not forked
	// add contribute langs...
	// todo: add private repos
	repos, _, _ := c.Repositories.List(ctx, "", nil)
	repoList := make([]string, len(repos))
	for i, r := range repos {
		if !*r.Fork {
			repoList[i] = *r.Name
			fmt.Println(*r.Name)
		}
	}
	return repoList
}

func langList(c *github.Client, r []string) map[string]int {
	langsMap := make(map[string]int, len(r))
	for _, name := range r {
		langs, _, _ := c.Repositories.ListLanguages(ctx, username, name)
		for k, v := range langs {
			langsMap[k] += v
		}
	}
	return langsMap
}
