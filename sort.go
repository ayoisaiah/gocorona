package main

import (
	"fmt"

	"github.com/gizak/termui/v3/widgets"
)

// SortOptions represents the sort options widget
type SortOptions struct {
	Widget *widgets.Paragraph
}

func (self *SortOptions) Construct() {
	widget := widgets.NewParagraph()
	widget.Title = "📈 Sorting options"
	widget.Text = fmt.Sprintf("F1 (default): [Total Cases](fg:black,bg:green) | F2: [Cases Today](fg:black,bg:green) | F3: [Total Deaths](fg:black,bg:green) | F4: [Deaths Today](fg:black,bg:green) | F5: [Recoveries](fg:black,bg:green) | F6: [Active](fg:black,bg:green) | F7: [Critical](fg:black,bg:green) | F8: [Mortality](fg:black,bg:green)")

	self.Widget = widget
}
