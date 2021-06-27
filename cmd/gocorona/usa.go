package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// USA represents the USA states table.
type USA struct {
	Table
}

// FetchData retrieves the latest data for each USA state
// that has stats available, and sorts it by total cases.
func (u *USA) FetchData() error {
	url := "https://disease.sh/v3/covid-19/states"
	return u.Table.FetchData(url)
}

// Construct constructs the USA states table widget.
func (u *USA) Construct() {
	p := message.NewPrinter(language.English)
	table := widgets.NewTable()
	tableHeader := []string{"#", "State", "Cases", "Cases (today)", "Deaths", "Deaths (today)", "Recoveries", "Active", "Mortality (IFR)", "Mortality (CFR)"}

	for i, v := range tableHeader {
		if v == u.Sort {
			tableHeader[i] = fmt.Sprintf("[%s](fg:yellow) ▼", tableHeader[i])
			break
		}
	}

	table.Rows = [][]string{tableHeader}

	for i := range u.Data {
		v := u.Data[i]
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
			p.Sprintf("%.2f%s", v.MortalityIFR*100, "%"), //nolint:gomnd // 100 is a constant
			p.Sprintf("%.2f%s", v.MortalityCFR*100, "%"), //nolint:gomnd // 100 is a constant
		})
	}

	table.ColumnWidths = []int{5, 30, 20, 20, 18, 18, 18, 18, 20, 20}
	table.TextAlignment = ui.AlignCenter
	table.TextStyle = ui.NewStyle(ui.ColorClear)
	table.FillRow = true
	table.RowSeparator = false
	table.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table.BorderLeft = false
	table.BorderRight = false

	if u.Widget == nil {
		u.Widget = table
		return
	}

	u.Widget.Rows = table.Rows
}
