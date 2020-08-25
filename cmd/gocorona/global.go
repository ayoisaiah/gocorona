package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// All represents up to date Global totals
type All struct {
	Cases               int     `json:"cases"`
	TodayCases          int     `json:"todayCases"`
	Deaths              int     `json:"deaths"`
	TodayDeaths         int     `json:"todayDeaths"`
	Recovered           int     `json:"recovered"`
	Active              int     `json:"active"`
	Critical            int     `json:"critical"`
	CasesPerOneMillion  int     `json:"casesPerOneMillion"`
	DeathsPerOneMillion float64 `json:"deathsPerOneMillion"`
	Updated             int64   `json:"updated"`
	AffectedCountries   int     `json:"affectedCountries"`
}

// Global represents the Global statistics widget
type Global struct {
	All
	Widget *widgets.Paragraph
}

// FetchData retrieves global statistics for
// cases, deaths, recovered, time last updated,
// and active cases
func (g *Global) FetchData() error {
	url := "https://disease.sh/v3/covid-19/all"
	return fetch(url, g)
}

// Construct creates the global statistics widget
func (g *Global) Construct(title string) {
	p := message.NewPrinter(language.English)

	widget := widgets.NewParagraph()
	widget.Title = title
	widget.Text = p.Sprintf("[Infections](fg:blue): %d (%d today)\n", g.Cases, g.TodayCases)
	widget.Text += p.Sprintf("[Deaths](fg:red): %d (%d today)\n", g.Deaths, g.TodayDeaths)
	widget.Text += p.Sprintf("[Recoveries](fg:green): %d (%d remaining)\n", g.Recovered, g.Active)
	if g.Critical > 0 {
		widget.Text += p.Sprintf("[Critical](fg:yellow): %d (%.2f%% of cases)\n", g.Critical, float64(g.Critical)/float64(g.Cases)*100)
	}
	widget.Text += p.Sprintf("[Mortality rate (IFR)](fg:cyan): %.2f%%\n", float64(g.Deaths)/float64(g.Cases)*100)
	widget.Text += p.Sprintf("[Mortality rate (CFR)](fg:cyan): %.2f%%\n", float64(g.Deaths)/(float64(g.Recovered)+float64(g.Deaths))*100)
	if g.AffectedCountries > 0 {
		widget.Text += p.Sprintf("[Affected Countries](fg:magenta): %d\n", g.AffectedCountries)
	}
	widget.SetRect(0, 0, 50, 10)
	widget.BorderStyle.Fg = ui.ColorYellow
	widget.TitleStyle = ui.NewStyle(ui.ColorClear)
	widget.TextStyle = ui.NewStyle(ui.ColorClear)

	g.Widget = widget
}
