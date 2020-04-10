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

	countriesTable, err := affectedCountriesTable()
	if err != nil {
		log.Fatal(err)
	}

	gs := ui.NewRow(0.25, ui.NewCol(1.0, globalStats))
	ct := ui.NewRow(0.67, ui.NewCol(1.0, countriesTable))
	st := ui.NewRow(0.08, ui.NewCol(1.0, shortcuts()))

	grid.Set(gs, ct, st)
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
				countriesTable.ScrollDown()
			case "k", "<Up>":
				countriesTable.ScrollUp()
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
