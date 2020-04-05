package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// getGlobalTotals retrieves global stats for
// cases, deaths, recovered, time last updated,
// and active cases
func getGlobalTotals() (*Global, error) {
	url := "https://corona.lmao.ninja/all"
	global := &Global{}

	return global, fetch(url, global)
}

func globalStatsWidget() (*widgets.Paragraph, error) {
	p := message.NewPrinter(language.English)
	globalStats, err := getGlobalTotals()
	if err != nil {
		return nil, err
	}

	widget := widgets.NewParagraph()
	widget.Title = "Global statistics"
	widget.Text = p.Sprintf("[Infections](fg:blue): %d (%d today)\n\n", globalStats.Cases, globalStats.TodayCases)
	widget.Text += p.Sprintf("[Deaths](fg:red): %d (%d today)\n\n", globalStats.Deaths, globalStats.TodayDeaths)
	widget.Text += p.Sprintf("[Recoveries](fg:green): %d (%d remaining)\n\n", globalStats.Recovered, globalStats.Active)
	widget.Text += p.Sprintf("[Critical](fg:yellow): %d\n\n", globalStats.Critical)
	widget.SetRect(0, 0, 50, 10)
	widget.BorderStyle.Fg = ui.ColorYellow

	return widget, nil
}
