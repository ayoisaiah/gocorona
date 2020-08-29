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
	widget.Text = "Press [q](fg:green) to quit, Press [h](fg:green) or [l](fg:green) to switch tabs, Press [k](fg:green) or [j](fg:green) to scroll up or down and [g](fg:green) or [G](fg:green) to scroll to the top or bottom"
	widget.Border = true
	widget.BorderStyle.Fg = ui.ColorYellow
	widget.TitleStyle = ui.NewStyle(ui.ColorClear)
	widget.TextStyle = ui.NewStyle(ui.ColorClear)

	i.Widget = widget
}
