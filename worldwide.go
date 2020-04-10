package main

import (
	"fmt"
	"sort"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type WorldwideTable struct {
	Data   Countries
	Widget *widgets.Table
	Sort   string
}

// Construct retrieves the latest data from the API
// if it does not exist already, sorts it by cases
// and constructs the table widget
func (wt *WorldwideTable) Construct() error {
	if len(wt.Data) == 0 {
		totals, err := getCountryTotals()
		if err != nil {
			return err
		}

		wt.Data = *totals
		wt.SortByCases()
		wt.Sort = "Total Cases"
	}

	p := message.NewPrinter(language.English)
	table := widgets.NewTable()
	tableHeader := []string{"#", "Country", "Total Cases", "Cases (today)", "Total Deaths", "Deaths (today)", "Recoveries", "Active", "Critical", "Mortality"}
	for i, v := range tableHeader {
		if v == wt.Sort {
			tableHeader[i] = fmt.Sprintf("[%s](fg:red) ▼", tableHeader[i])
			break
		}
	}

	table.Rows = [][]string{tableHeader}

	for i, v := range wt.Data {
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

	if wt.Widget == nil {
		wt.Widget = table
	} else {
		wt.Widget.Rows = table.Rows
	}

	return nil
}

// SortByCases sorts the countries by number of total cases
func (wt *WorldwideTable) SortByCases() {
	sort.SliceStable(wt.Data, func(i, j int) bool {
		return wt.Data[i].Cases > wt.Data[j].Cases
	})
	wt.Sort = "Total Cases"
	wt.Construct()
}

// SortByCasesToday sorts the countries by number of cases today
func (wt *WorldwideTable) SortByCasesToday() {
	sort.SliceStable(wt.Data, func(i, j int) bool {
		return wt.Data[i].TodayCases > wt.Data[j].TodayCases
	})
	wt.Sort = "Cases (today)"
	wt.Construct()
}

// SortByDeaths sorts the countries by number of total deaths
func (wt *WorldwideTable) SortByDeaths() {
	sort.SliceStable(wt.Data, func(i, j int) bool {
		return wt.Data[i].Deaths > wt.Data[j].Deaths
	})
	wt.Sort = "Total Deaths"
	wt.Construct()
}

// SortByDeathsToday sorts the countries by number of deaths today
func (wt *WorldwideTable) SortByDeathsToday() {
	sort.SliceStable(wt.Data, func(i, j int) bool {
		return wt.Data[i].TodayDeaths > wt.Data[j].TodayDeaths
	})
	wt.Sort = "Deaths (today)"
	wt.Construct()
}

// SortByRecoveries sorts the countries by number of total recoveries
func (wt *WorldwideTable) SortByRecoveries() {
	sort.SliceStable(wt.Data, func(i, j int) bool {
		return wt.Data[i].Recovered > wt.Data[j].Recovered
	})
	wt.Sort = "Recoveries"
	wt.Construct()
}

// SortByActive sorts the countries by number of active cases
func (wt *WorldwideTable) SortByActive() {
	sort.SliceStable(wt.Data, func(i, j int) bool {
		return wt.Data[i].Active > wt.Data[j].Active
	})
	wt.Sort = "Active"
	wt.Construct()
}

// SortByCritical sorts the countries by number of critical cases
func (wt *WorldwideTable) SortByCritical() {
	sort.SliceStable(wt.Data, func(i, j int) bool {
		return wt.Data[i].Critical > wt.Data[j].Critical
	})
	wt.Sort = "Critical"
	wt.Construct()
}

// SortByMortality sorts the countries by mortality rate
func (wt *WorldwideTable) SortByMortality() {
	sort.SliceStable(wt.Data, func(i, j int) bool {
		return float64(wt.Data[i].Deaths)/float64(wt.Data[i].Cases) > float64(wt.Data[j].Deaths)/float64(wt.Data[j].Cases)
	})
	wt.Sort = "Mortality"
	wt.Construct()
}

func getCountryTotals() (*Countries, error) {
	url := "https://corona.lmao.ninja/countries"
	countries := &Countries{}

	return countries, fetch(url, countries)
}
