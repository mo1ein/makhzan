package makhzan

import (
	"context"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/google/go-github/v47/github"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
	"log"
	"os"
	"syscall"
)

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

	// pie chart configs
	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: t,
		TitleStyle: &opts.TextStyle{
			FontFamily: "sens-serif",
		},
		Left: "center",
	}),
	)
	pie.SetGlobalOptions(charts.WithTooltipOpts(opts.Tooltip{
		Show:      true,
		Trigger:   "item",
		TriggerOn: "mousemove",
	}))
	pie.SetGlobalOptions(charts.WithLegendOpts(opts.Legend{
		Show:   true,
		Orient: "horizontal",
		Top:    "bottom",
	}))
	pie.AddSeries("Languages", destinations)
	pie.SetSeriesOptions(charts.WithLabelOpts(
		opts.Label{
			Show:      true,
			Formatter: "{b}: {d}%",
		}),
	)
	f, _ := os.Create("pie.html")
	pie.Render(f)
}

func GhAuth() (*github.Client, context.Context) {
	fmt.Print("Enter GitHub Token: ")
	byteToken, _ := terminal.ReadPassword(int(syscall.Stdin))
	println()
	token := string(byteToken)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client, ctx
}

// TODO: add contrubuted langs
// TODO: add private repos

// Return string slice of Repos(Not forked repos)
func RepoList(c *github.Client, ctx context.Context) ([]string, error) {
	repos, resp, err := c.Repositories.List(ctx, "", nil)
	projects := make([]string, len(repos))
	if err != nil {
		return projects, err
	}

	// Rate.Limit should most likely be 5000 when authorized.
	log.Printf("Rate: %#v\n", resp.Rate)

	// If a Token Expiration has been set, it will be displayed.
	if !resp.TokenExpiration.IsZero() {
		log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
	}

	for i, r := range repos {
		if !*r.Fork {
			projects[i] = *r.Name
		}
	}
	return projects, nil
}

// Return map of languages `string` as key and `int` as value
func LangList(c *github.Client, ctx context.Context, repos []string, u string) map[string]int {
	langMap := make(map[string]int, len(repos))
	for _, name := range repos {
		// this is not beautiful
		langs, _, _ := c.Repositories.ListLanguages(ctx, u, name)
		for k, v := range langs {
			langMap[k] += v
		}
	}
	return langMap
}
