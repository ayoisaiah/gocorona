package main

import (
	"fmt"

	"github.com/gizak/termui/v3/widgets"
)

func shortcuts() *widgets.Paragraph {
	widget := widgets.NewParagraph()
	widget.Title = "Keyboard shortcuts"
	widget.Text = fmt.Sprintf("F1: [Help](fg:black,bg:green) F2: [Sort](fg:black,bg:green)")

	return widget
}
