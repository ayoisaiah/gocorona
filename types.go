package main

type Global struct {
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

type Countries []struct {
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
