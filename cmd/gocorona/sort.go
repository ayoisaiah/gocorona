package main

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// SortOptions represents the sort options widget
type SortOptions struct {
	Widget *widgets.Paragraph
}

func (so *SortOptions) Construct() {
	widget := widgets.NewParagraph()
	widget.Title = "📈 Sorting options"
	widget.Text = fmt.Sprintf("F1 (default): [ Cases ](fg:black,bg:green) | F2: [ Cases Today ](fg:black,bg:green) | F3: [ Deaths ](fg:black,bg:green) | F4: [ Deaths Today ](fg:black,bg:green) | F5: [ Recoveries ](fg:black,bg:green) | F6: [ Active ](fg:black,bg:green) | F7: [ Critical ](fg:black,bg:green) | F8: [ Mortality (IFR) ](fg:black,bg:green) | F9: [ Mortality (CFR) ](fg:black,bg:green)")
	widget.TextStyle = ui.NewStyle(ui.ColorClear)
	widget.TitleStyle = ui.NewStyle(ui.ColorClear)

	so.Widget = widget
}
