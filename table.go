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
func (self *Table) FetchData(url string) error {
	err := fetch(url, &self.Data)
	if err != nil {
		return err
	}

	self.SortByCases()
	self.Sort = "Total Cases"

	return nil
}

// Construct calls the Construct method of the parent
func (self *Table) Construct() {
	self.parent.Construct()
}

// SortByCases sorts the data by number of total cases
func (self *Table) SortByCases() {
	sort.SliceStable(self.Data, func(i, j int) bool {
		return self.Data[i].Cases > self.Data[j].Cases
	})
	self.Sort = "Total Cases"
	self.Construct()
}

// SortByCasesToday sorts the data by number of cases today
func (self *Table) SortByCasesToday() {
	sort.SliceStable(self.Data, func(i, j int) bool {
		return self.Data[i].TodayCases > self.Data[j].TodayCases
	})
	self.Sort = "Cases (today)"
	self.Construct()
}

// SortByDeaths sorts the data by number of total deaths
func (self *Table) SortByDeaths() {
	sort.SliceStable(self.Data, func(i, j int) bool {
		return self.Data[i].Deaths > self.Data[j].Deaths
	})
	self.Sort = "Total Deaths"
	self.Construct()
}

// SortByDeathsToday sorts the data by number of deaths today
func (self *Table) SortByDeathsToday() {
	sort.SliceStable(self.Data, func(i, j int) bool {
		return self.Data[i].TodayDeaths > self.Data[j].TodayDeaths
	})
	self.Sort = "Deaths (today)"
	self.Construct()
}

// SortByActive sorts the data by number of active cases
func (self *Table) SortByActive() {
	sort.SliceStable(self.Data, func(i, j int) bool {
		return self.Data[i].Active > self.Data[j].Active
	})
	self.Sort = "Active"
	self.Construct()
}

// SortByRecoveries sorts the data by number of total recoveries
func (self *Table) SortByRecoveries() {
	sort.SliceStable(self.Data, func(i, j int) bool {
		return self.Data[i].Recovered > self.Data[j].Recovered
	})
	self.Sort = "Recoveries"
	self.Construct()
}

// SortByMortality sorts the data by mortality rate
func (self *Table) SortByMortality() {
	sort.SliceStable(self.Data, func(i, j int) bool {
		return float64(self.Data[i].Deaths)/float64(self.Data[i].Cases) > float64(self.Data[j].Deaths)/float64(self.Data[j].Cases)
	})
	self.Sort = "Mortality"
	self.Construct()
}
