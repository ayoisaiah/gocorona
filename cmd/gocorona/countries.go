package main

import (
	"fmt"
	"sort"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

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
			tableHeader[i] = fmt.Sprintf("[%s](fg:yellow) ▼", tableHeader[i])
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
	table.TextStyle = ui.NewStyle(ui.ColorClear)
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

// FilterByName returns the table data for the specified country name
func (c *Countries) FilterByName(name string) *TableData {
	for _, v := range c.Data {
		if v.Country == name {
			return &v
		}
	}

	return nil
}
