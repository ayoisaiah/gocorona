package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func getCountriesTotals() (*Countries, error) {
	url := "https://corona.lmao.ninja/countries"
	countries := &Countries{}

	return countries, fetch(url, countries)
}

func affectedCountriesTable() (*widgets.Table, error) {
	p := message.NewPrinter(language.English)
	countryTotals, err := getCountriesTotals()
	if err != nil {
		return nil, err
	}

	table := widgets.NewTable()
	table.Rows = [][]string{
		[]string{"#", "Country", "Cases", "Cases (today)", "Deaths", "Deaths (today)", "Recoveries", "Active", "Critical"},
	}

	for i, v := range *countryTotals {
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
		})

		if i%2 == 0 {
			table.RowStyles[i] = ui.NewStyle(ui.ColorBlack, ui.ColorYellow)
		}
	}

	table.ColumnWidths = []int{5, 20, 15, 15, 15, 15, 15, 15, 15}
	table.TextAlignment = ui.AlignCenter
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.FillRow = true
	table.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table.SetRect(0, 11, 140, 50)

	return table, nil
}
