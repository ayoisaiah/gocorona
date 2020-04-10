package main

import (
	"fmt"

	"github.com/gizak/termui/v3/widgets"
)

func shortcuts() *widgets.Paragraph {
	widget := widgets.NewParagraph()
	widget.Title = "Sorting options"
	widget.Text = fmt.Sprintf("F1 (default): [Cases](fg:black,bg:green) | F2: [Cases Today](fg:black,bg:green) | F3: [Deaths](fg:black,bg:green) | F4: [Deaths Today](fg:black,bg:green) | F5: [Recoveries](fg:black,bg:green) | F6: [Active](fg:black,bg:green) | F7: [Critical](fg:black,bg:green) | F8: [Mortality](fg:black,bg:green)")

	return widget
}
