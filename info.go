package main

import "github.com/gizak/termui/v3/widgets"

// CoronavirusInfo represents the widget
// that provides info about the coronavirus pandemic
type CoronavirusInfo struct {
	Widget *widgets.Paragraph
}

// Construct creates a widget containing
// info about the coronavirus pandamic and how
// to protect oneself
func (ci *CoronavirusInfo) Construct() {
	widget := widgets.NewParagraph()

	widget.Title = "🤒 Learn about the Coronavirus Pandemic"
	widget.Text = `
[There is currently No Vaccine to prevent Coronavirus](fg:red)

[How it spreads](fg:white,bg:yellow,fg:bold)

The virus is thought to spread mainly from person-to-person contact through respiratory droplets produced when an infected person coughs or sneezes.

[Symptoms](fg:white,bg:yellow,fg:bold)

The main symptoms of coronavirus are:
👉 a high temperature – this means you feel hot to touch on your chest or back (you do not need to measure your temperature)
👉 a new, continuous cough – this means coughing a lot for more than an hour, or 3 or more coughing episodes in 24 hours (if you usually have a cough, it may be worse than usual)

To protect others, do not go to places like a GP surgery, pharmacy or hospital if you have these symptoms. Stay at home.

[Prevention](fg:white,bg:yellow,fg:bold)

Everyone must stay at home to help stop coronavirus (COVID-19) spreading. Wash your hands with soap and water often to reduce the risk of infection.

To stop the spread of coronavirus, you should:
👉 wash your hands with soap and water often – for at least 20 seconds
👉 wash your hands as soon as you get home (if you leave home for any reason)
👉 cover your mouth and nose with a tissue when you cough or sneeze
👉 put used tissues in the bin immediately and wash your hands
👉 not touch your face if your hands are not clean

Learn more at [https://www.nhs.uk/conditions/coronavirus-covid-19/](fg:blue)
	`

	ci.Widget = widget
}
