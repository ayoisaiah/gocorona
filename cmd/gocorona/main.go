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

	errs.Go(func() error {
		return global.FetchData()
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

	global.Construct()

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
	countriesTable := ui.NewRow(0.56, ui.NewCol(1.0, countries.Widget))
	sortWidget := ui.NewRow(0.08, ui.NewCol(1.0, sortOptions.Widget))
	usaTable := ui.NewRow(0.84, ui.NewCol(1.0, usa.Widget))
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
			grid.Set(tabWidget, sortWidget, usaTable)
		case 2:
			grid.Set(tabWidget, infoWidget)
		case 3:
			grid.Set(tabWidget, creditsWidget)
		}
	}

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			currentTable.Widget.ScrollDown()
		case "k", "<Up>":
			currentTable.Widget.ScrollUp()
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
