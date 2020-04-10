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
		[]string{"#", "Country", "Cases (today)", "Deaths (today)", "Recoveries", "Active", "Critical", "Mortality"},
	}

	for i, v := range *countryTotals {
		table.Rows = append(table.Rows, []string{
			p.Sprintf("%d", i+1),
			v.Country,
			p.Sprintf("%d (%d)", v.Cases, v.TodayCases),
			p.Sprintf("%d (%d)", v.Deaths, v.TodayDeaths),
			p.Sprintf("%d", v.Recovered),
			p.Sprintf("%d", v.Active),
			p.Sprintf("%d", v.Critical),
			p.Sprintf("%.2f%s", float64(v.Deaths)/float64(v.Cases)*100, "%"),
		})
	}

	table.ColumnWidths = []int{5, 20, 20, 20, 15, 15, 13, 13, 14}
	table.TextAlignment = ui.AlignCenter
	table.TextStyle = ui.NewStyle(ui.ColorWhite)
	table.FillRow = true
	table.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	table.SetRect(0, 11, 150, 50)

	return table, nil
}
