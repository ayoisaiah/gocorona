package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Instructions struct {
	Widget *widgets.Paragraph
}

func (i *Instructions) Construct() {
	widget := widgets.NewParagraph()
	widget.Title = "👉 Navigation"
	widget.Text = "Press q to quit, Press h or l to switch tabs, Press k or j to scroll up or down and g or G to scroll to the top or bottom"
	widget.Border = true
	widget.BorderStyle.Fg = ui.ColorYellow
	widget.TitleStyle = ui.NewStyle(ui.ColorClear)
	widget.TextStyle = ui.NewStyle(ui.ColorClear)

	i.Widget = widget
}
