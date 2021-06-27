package gocorona

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Tab represents the tab widget.
type Tab struct {
	Widget *widgets.TabPane
}

// Construct creates the tab widget for
// switching between views.
func (t *Tab) Construct() {
	widget := widgets.NewTabPane("🌎 Worldwide", "🇺  USA", "💉 Vaccine tracker", "😷 Protect yourself", "👌 Credits")
	widget.Border = true
	widget.InactiveTabStyle = ui.NewStyle(ui.ColorClear)

	t.Widget = widget
}
