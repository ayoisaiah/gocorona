package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	globalStats, err := globalStatsWidget()
	if err != nil {
		log.Fatal(err)
	}

	countriesTable, err := affectedCountriesTable()
	if err != nil {
		log.Fatal(err)
	}

	ui.Render(globalStats, countriesTable)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}

}
