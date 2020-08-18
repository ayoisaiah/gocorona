package main

import (
	"fmt"
	"sort"

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
	url := "https://corona.lmao.ninja/v2/all"
	return fetch(url, g)
}

// Construct creates the global statistics widget
func (g *Global) Construct() {
	p := message.NewPrinter(language.English)

	widget := widgets.NewParagraph()
	widget.Title = "🌐 Global statistics"
	widget.Text = p.Sprintf("[Infections](fg:blue): %d (%d today)\n", g.Cases, g.TodayCases)
	widget.Text += p.Sprintf("[Deaths](fg:red): %d (%d today)\n", g.Deaths, g.TodayDeaths)
	widget.Text += p.Sprintf("[Recoveries](fg:green): %d (%d remaining)\n", g.Recovered, g.Active)
	widget.Text += p.Sprintf("[Critical](fg:yellow): %d (%.2f%% of cases)\n", g.Critical, float64(g.Critical)/float64(g.Cases)*100)
	widget.Text += p.Sprintf("[Mortality rate (IFR)](fg:cyan): %.2f%%\n", float64(g.Deaths)/float64(g.Cases)*100)
	widget.Text += p.Sprintf("[Mortality rate (CFR)](fg:cyan): %.2f%%\n", float64(g.Deaths)/(float64(g.Recovered)+float64(g.Deaths))*100)
	widget.Text += p.Sprintf("[Affected Countries](fg:magenta): %d\n", g.AffectedCountries)
	widget.SetRect(0, 0, 50, 10)
	widget.BorderStyle.Fg = ui.ColorYellow

	g.Widget = widget
}

// Countries represents the countries table
type Countries struct {
	Table
}

// FetchData retrieves the latest data for each country
// that has stats available, and sorts it by total cases
func (c *Countries) FetchData() error {
	url := "https://corona.lmao.ninja/v2/countries"
	return c.Table.FetchData(url)
}

// Construct constructs the countries table widget
func (c *Countries) Construct() {
	p := message.NewPrinter(language.English)
	table := widgets.NewTable()
	tableHeader := []string{"#", "Country", "Total Cases", "Cases (today)", "Total Deaths", "Deaths (today)", "Recoveries", "Active", "Critical", "Mortality (IFR)", "Mortality (CFR)"}
	for i, v := range tableHeader {
		if v == c.Sort {
			tableHeader[i] = fmt.Sprintf("[%s](fg:red) ▼", tableHeader[i])
			break
		}
	}

	table.Rows = [][]string{tableHeader}

	for i, v := range c.Data {
		table.Rows = append(table.Rows, []string{
			p.Sprintf("%d", i+1),
			v.Country,
			p.Sprintf("%d", v.Cases),
			p.Sprintf("%d", v.TodayCases),
			p.Sprintf("%d", v.Deaths),
			p.Sprintf("%d", v.TodayDeaths),
			p.Sprintf("%d", v.Recovered),
			p.Sprintf("%d", v.Active),
			p.Sprintf("%d", v.Critical),
			p.Sprintf("%.2f%s", v.MortalityIFR*100, "%"),
			p.Sprintf("%.2f%s", v.MortalityCFR*100, "%"),
		})
	}

	table.ColumnWidths = []int{5, 22, 20, 20, 18, 18, 15, 15, 15, 20, 20}
	table.TextAlignment = ui.AlignCenter
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.FillRow = true
	table.RowSeparator = false
	table.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table.BorderLeft = false
	table.BorderRight = false

	if c.Widget == nil {
		c.Widget = table
	} else {
		c.Widget.Rows = table.Rows
	}
}

// SortByCritical sorts the countries by number of critical cases
func (c *Countries) SortByCritical() {
	sort.SliceStable(c.Data, func(i, j int) bool {
		return c.Data[i].Critical > c.Data[j].Critical
	})
	c.Sort = "Critical"
	c.Construct()
}
