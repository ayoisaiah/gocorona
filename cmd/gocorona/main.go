package main

import (
	"context"
	"log"

	ui "github.com/gizak/termui/v3"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}

	defer ui.Close()

	ui.Theme.Default.Bg = ui.ColorBlack

	loading := &Loading{}
	loading.Construct()

	ui.Render(loading.Widget)

	tw, th := ui.TerminalDimensions()
	grid := ui.NewGrid()
	grid.SetRect(0, 0, tw, th)

	errs, _ := errgroup.WithContext(context.TODO())

	global := &Global{}
	countries := &Countries{}
	countries.parent = countries
	usa := &USA{}
	usa.parent = usa
	vaccine := &Vaccine{}

	errs.Go(func() error {
		return global.FetchData()
	})

	errs.Go(func() error {
		return vaccine.FetchData()
	})

	errs.Go(func() error {
		return countries.FetchData()
	})

	errs.Go(func() error {
		return usa.FetchData()
	})

	err := errs.Wait()
	if err != nil {
		log.Fatal(err)
	}

	global.Construct("🌐 Global statistics")

	usaData := countries.FilterByName("USA")
	usaAggregate := &Global{}
	usaAggregate.All = All{
		Cases:       usaData.Cases,
		TodayCases:  usaData.TodayCases,
		Deaths:      usaData.Deaths,
		TodayDeaths: usaData.TodayDeaths,
		Recovered:   usaData.Recovered,
		Active:      usaData.Active,
	}
	usaAggregate.Construct("📈 Cases overview")

	vaccine.Construct()

	tab := &Tab{}
	tab.Construct()
	tabpane := tab.Widget

	credits := &Credits{}
	credits.Construct()

	sortOptions := &SortOptions{}
	sortOptions.Construct()

	coronavirusInfo := &CoronavirusInfo{}
	coronavirusInfo.Construct()

	instructions := &Instructions{}
	instructions.Construct()

	tabWidget := ui.NewRow(0.08, ui.NewCol(1.0, tabpane))
	globalWidget := ui.NewRow(0.20, ui.NewCol(1.0, global.Widget))
	usaAggregateWidget := ui.NewRow(0.16, ui.NewCol(1.0, usaAggregate.Widget))
	countriesTable := ui.NewRow(0.56, ui.NewCol(1.0, countries.Widget))
	sortWidget := ui.NewRow(0.08, ui.NewCol(1.0, sortOptions.Widget))
	usaTable := ui.NewRow(0.68, ui.NewCol(1.0, usa.Widget))
	vaccinePhaseWidget := ui.NewRow(0.16, ui.NewCol(1.0, vaccine.PhaseWidget))
	vaccineListWidget := ui.NewRow(0.76, ui.NewCol(1.0, vaccine.CandidatesWidget))
	infoWidget := ui.NewRow(0.92, ui.NewCol(1.0, coronavirusInfo.Widget))
	creditsWidget := ui.NewRow(0.92, ui.NewCol(1.0, credits.Widget))
	instructionsWidget := ui.NewRow(0.08, ui.NewCol(1.0, instructions.Widget))

	currentTable := &countries.Table

	grid.Set(tabWidget, globalWidget, sortWidget, countriesTable, instructionsWidget)
	ui.Clear()
	ui.Render(grid)

	renderTab := func() {
		grid.Items = nil
		switch tabpane.ActiveTabIndex {
		case 0:
			currentTable = &countries.Table
			grid.Set(tabWidget, globalWidget, sortWidget, countriesTable, instructionsWidget)
		case 1:
			currentTable = &usa.Table
			grid.Set(tabWidget, usaAggregateWidget, sortWidget, usaTable)
		case 2:
			grid.Set(tabWidget, vaccinePhaseWidget, vaccineListWidget)
		case 3:
			grid.Set(tabWidget, infoWidget)
		case 4:
			grid.Set(tabWidget, creditsWidget)
		}
	}

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>", "<Escape>":
			return
		case "j", "<Down>":
			currentTable.Widget.ScrollDown()
			if tabpane.ActiveTabIndex == 2 {
				vaccine.CandidatesWidget.ScrollDown()
			}
		case "k", "<Up>":
			currentTable.Widget.ScrollUp()
			if tabpane.ActiveTabIndex == 2 {
				vaccine.CandidatesWidget.ScrollUp()
			}
		case "g", "<Home>":
			currentTable.Widget.ScrollTop()
		case "G", "<End>":
			currentTable.Widget.ScrollBottom()
		case "<PageUp>":
			currentTable.Widget.ScrollPageUp()
		case "<PageDown>":
			currentTable.Widget.ScrollPageDown()
		case "<F1>":
			currentTable.SortByCases()
		case "<F2>":
			currentTable.SortByCasesToday()
		case "<F3>":
			currentTable.SortByDeaths()
		case "<F4>":
			currentTable.SortByDeathsToday()
		case "<F5>":
			currentTable.SortByRecoveries()
		case "<F6>":
			currentTable.SortByActive()
		case "<F7>":
			countries.SortByCritical()
		case "<F8>":
			currentTable.SortByMortalityIFR()
		case "<F9>":
			currentTable.SortByMortalityCFR()
		case "<Resize>":
			tw, th = ui.TerminalDimensions()
			grid.SetRect(0, 0, tw, th)
		case "h":
			tabpane.FocusLeft()
			renderTab()
		case "l":
			tabpane.FocusRight()
			renderTab()
		}
		ui.Render(grid)
	}
}
