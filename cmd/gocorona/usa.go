package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// USA represents the USA states table
type USA struct {
	Table
}

// FetchData retrieves the latest data for each USA state
// that has stats available, and sorts it by total cases
func (u *USA) FetchData() error {
	url := "https://disease.sh/v3/covid-19/states"
	return u.Table.FetchData(url)
}

// Construct constructs the USA states table widget
func (u *USA) Construct() {
	p := message.NewPrinter(language.English)
	table := widgets.NewTable()
	tableHeader := []string{"#", "State", "Total Cases", "Cases (today)", "Total Deaths", "Deaths (today)", "Recoveries", "Active", "Mortality (IFR)", "Mortality (CFR)"}
	for i, v := range tableHeader {
		if v == u.Sort {
			tableHeader[i] = fmt.Sprintf("[%s](fg:yellow) ▼", tableHeader[i])
			break
		}
	}

	table.Rows = [][]string{tableHeader}

	for i, v := range u.Data {
		u.Data[i].Recovered = v.Cases - v.Deaths - v.Active
		table.Rows = append(table.Rows, []string{
			p.Sprintf("%d", i+1),
			v.State,
			p.Sprintf("%d", v.Cases),
			p.Sprintf("%d", v.TodayCases),
			p.Sprintf("%d", v.Deaths),
			p.Sprintf("%d", v.TodayDeaths),
			p.Sprintf("%d", v.Recovered),
			p.Sprintf("%d", v.Active),
			p.Sprintf("%.2f%s", v.MortalityIFR*100, "%"),
			p.Sprintf("%.2f%s", v.MortalityCFR*100, "%"),
		})
	}

	table.ColumnWidths = []int{5, 25, 20, 20, 18, 18, 15, 15, 20, 20}
	table.TextAlignment = ui.AlignCenter
	table.TextStyle = ui.NewStyle(ui.ColorClear)
	table.FillRow = true
	table.RowSeparator = false
	table.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table.BorderLeft = false
	table.BorderRight = false

	if u.Widget == nil {
		u.Widget = table
	} else {
		u.Widget.Rows = table.Rows
	}
}
