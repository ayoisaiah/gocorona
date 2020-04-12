package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Instructions struct {
	Widget *widgets.Paragraph
}

func (self *Instructions) Construct() {
	widget := widgets.NewParagraph()
	widget.Title = "👉 Navigation"
	widget.Text = "Press q to quit, Press h or l to switch tabs, Press j or k to scroll up or down"
	widget.Border = true
	widget.BorderStyle.Fg = ui.ColorYellow

	self.Widget = widget
}
