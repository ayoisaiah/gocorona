package main

import "github.com/gizak/termui/v3/widgets"

// Tab represents the tab widget
type Tab struct {
	Widget *widgets.TabPane
}

// Construct creates the tab widget for
// switching between views
func (t *Tab) Construct() {
	widget := widgets.NewTabPane("🌎 Global", " 🇺  USA", "😷 Protect Yourself", "👌 Credits")
	widget.Border = true

	t.Widget = widget
}
