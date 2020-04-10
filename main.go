package main

import (
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	tw, th := ui.TerminalDimensions()

	grid := ui.NewGrid()
	grid.SetRect(0, 0, tw, th)

	globalStats, err := globalStatsWidget()
	if err != nil {
		log.Fatal(err)
	}

	worldwideTable := &WorldwideTable{}
	err = worldwideTable.Construct()
	if err != nil {
		log.Fatal(err)
	}

	gs := ui.NewRow(0.25, ui.NewCol(1.0, globalStats))
	ct := ui.NewRow(0.67, ui.NewCol(1.0, worldwideTable.Widget))
	st := ui.NewRow(0.08, ui.NewCol(1.0, shortcuts()))

	grid.Set(gs, st, ct)
	ui.Render(grid)

	ticker := time.Tick(time.Second / time.Duration(60))
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				worldwideTable.Widget.ScrollDown()
			case "k", "<Up>":
				worldwideTable.Widget.ScrollUp()
			case "<F1>":
				worldwideTable.SortByCases()
			case "<F2>":
				worldwideTable.SortByCasesToday()
			case "<F3>":
				worldwideTable.SortByDeaths()
			case "<F4>":
				worldwideTable.SortByDeathsToday()
			case "<F5>":
				worldwideTable.SortByRecoveries()
			case "<F6>":
				worldwideTable.SortByActive()
			case "<F7>":
				worldwideTable.SortByCritical()
			case "<F8>":
				worldwideTable.SortByMortality()
			case "<Resize>":
				tw, th = ui.TerminalDimensions()
				grid.SetRect(0, 0, tw, th)
				ui.Render(grid)
			}
		case <-ticker:
			ui.Render(grid)
		}
	}
}
