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
func (self *Global) FetchData() error {
	url := "https://corona.lmao.ninja/v2/all"
	return fetch(url, self)
}

// Construct creates the global statistics widget
func (self *Global) Construct() {
	p := message.NewPrinter(language.English)

	widget := widgets.NewParagraph()
	widget.Title = "🌐 Global statistics"
	widget.Text = p.Sprintf("[Infections](fg:blue): %d (%d today)\n", self.Cases, self.TodayCases)
	widget.Text += p.Sprintf("[Deaths](fg:red): %d (%d today)\n", self.Deaths, self.TodayDeaths)
	widget.Text += p.Sprintf("[Recoveries](fg:green): %d (%d remaining)\n", self.Recovered, self.Active)
	widget.Text += p.Sprintf("[Critical](fg:yellow): %d (%.2f%% of cases)\n", self.Critical, float64(self.Critical)/float64(self.Cases)*100)
	widget.Text += p.Sprintf("[Mortality rate](fg:cyan): %.2f%%\n", float64(self.Deaths)/float64(self.Cases)*100)
	widget.Text += p.Sprintf("[Affected Countries](fg:magenta): %d\n", self.AffectedCountries)
	widget.SetRect(0, 0, 50, 10)
	widget.BorderStyle.Fg = ui.ColorYellow

	self.Widget = widget
}

// Countries represents the countries table
type Countries struct {
	Table
}

// FetchData retrieves the latest data for each country
// that has stats available, and sorts it by total cases
func (self *Countries) FetchData() error {
	self.parent = self
	url := "https://corona.lmao.ninja/v2/countries"
	return self.Table.FetchData(url)
}

// Construct constructs the countries table widget
func (self *Countries) Construct() {
	p := message.NewPrinter(language.English)
	table := widgets.NewTable()
	tableHeader := []string{"#", "Country", "Total Cases", "Cases (today)", "Total Deaths", "Deaths (today)", "Recoveries", "Active", "Critical", "Mortality"}
	for i, v := range tableHeader {
		if v == self.Sort {
			tableHeader[i] = fmt.Sprintf("[%s](fg:red) ▼", tableHeader[i])
			break
		}
	}

	table.Rows = [][]string{tableHeader}

	for i, v := range self.Data {
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
			p.Sprintf("%.2f%s", float64(v.Deaths)/float64(v.Cases)*100, "%"),
		})
	}

	table.ColumnWidths = []int{5, 22, 20, 20, 18, 18, 15, 15, 15, 15}
	table.TextAlignment = ui.AlignCenter
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.FillRow = true
	table.RowSeparator = false
	table.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table.BorderLeft = false
	table.BorderRight = false

	if self.Widget == nil {
		self.Widget = table
	} else {
		self.Widget.Rows = table.Rows
	}
}

// SortByCritical sorts the countries by number of critical cases
func (self *Countries) SortByCritical() {
	sort.SliceStable(self.Data, func(i, j int) bool {
		return self.Data[i].Critical > self.Data[j].Critical
	})
	self.Sort = "Critical"
	self.Construct()
}
