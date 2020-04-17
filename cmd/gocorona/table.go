package main

import (
	"sort"

	"github.com/gizak/termui/v3/widgets"
)

type Table struct {
	Data []struct {
		State       string `json:"state"`
		Country     string `json:"country"`
		CountryInfo struct {
			ID   int     `json:"_id"`
			Iso2 string  `json:"iso2"`
			Iso3 string  `json:"iso3"`
			Lat  float64 `json:"lat"`
			Long float64 `json:"long"`
			Flag string  `json:"flag"`
		} `json:"countryInfo"`
		Updated             int64   `json:"updated"`
		Cases               int     `json:"cases"`
		TodayCases          int     `json:"todayCases"`
		Deaths              int     `json:"deaths"`
		TodayDeaths         int     `json:"todayDeaths"`
		Recovered           int     `json:"recovered"`
		Active              int     `json:"active"`
		Critical            int     `json:"critical"`
		CasesPerOneMillion  float64 `json:"casesPerOneMillion"`
		DeathsPerOneMillion float64 `json:"deathsPerOneMillion"`
		Tests               int     `json:"tests"`
		TestsPerOneMillion  float64 `json:"testsPerOneMillion"`
	}
	Widget *widgets.Table
	Sort   string
	parent Parent
}

type Parent interface {
	Construct()
}

// FetchData retrieves the latest data from the `url`
// and sorts it by total cases
func (t *Table) FetchData(url string) error {
	err := fetch(url, &t.Data)
	if err != nil {
		return err
	}

	t.SortByCases()
	return nil
}

// Construct calls the Construct method of the parent
func (t *Table) Construct() {
	t.parent.Construct()
}

// SortByCases sorts the data by number of total cases
func (t *Table) SortByCases() {
	sort.SliceStable(t.Data, func(i, j int) bool {
		return t.Data[i].Cases > t.Data[j].Cases
	})
	t.Sort = "Total Cases"
	t.Construct()
}

// SortByCasesToday sorts the data by number of cases today
func (t *Table) SortByCasesToday() {
	sort.SliceStable(t.Data, func(i, j int) bool {
		return t.Data[i].TodayCases > t.Data[j].TodayCases
	})
	t.Sort = "Cases (today)"
	t.Construct()
}

// SortByDeaths sorts the data by number of total deaths
func (t *Table) SortByDeaths() {
	sort.SliceStable(t.Data, func(i, j int) bool {
		return t.Data[i].Deaths > t.Data[j].Deaths
	})
	t.Sort = "Total Deaths"
	t.Construct()
}

// SortByDeathsToday sorts the data by number of deaths today
func (t *Table) SortByDeathsToday() {
	sort.SliceStable(t.Data, func(i, j int) bool {
		return t.Data[i].TodayDeaths > t.Data[j].TodayDeaths
	})
	t.Sort = "Deaths (today)"
	t.Construct()
}

// SortByActive sorts the data by number of active cases
func (t *Table) SortByActive() {
	sort.SliceStable(t.Data, func(i, j int) bool {
		return t.Data[i].Active > t.Data[j].Active
	})
	t.Sort = "Active"
	t.Construct()
}

// SortByRecoveries sorts the data by number of total recoveries
func (t *Table) SortByRecoveries() {
	sort.SliceStable(t.Data, func(i, j int) bool {
		return t.Data[i].Recovered > t.Data[j].Recovered
	})
	t.Sort = "Recoveries"
	t.Construct()
}

// SortByMortality sorts the data by mortality rate
func (t *Table) SortByMortality() {
	sort.SliceStable(t.Data, func(i, j int) bool {
		return float64(t.Data[i].Deaths)/float64(t.Data[i].Cases) > float64(t.Data[j].Deaths)/float64(t.Data[j].Cases)
	})
	t.Sort = "Mortality"
	t.Construct()
}
