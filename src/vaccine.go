package gocorona

import (
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// VaccineData is represents the response from disease.sh.
type VaccineData struct {
	Source          string `json:"source"`
	TotalCandidates string `json:"totalCandidates"`
	Phases          []struct {
		Phase      string `json:"phase"`
		Candidates string `json:"candidates"`
	} `json:"phases"`
	Data []struct {
		Candidate    string   `json:"candidate"`
		Sponsors     []string `json:"sponsors"`
		Details      string   `json:"details"`
		TrialPhase   string   `json:"trialPhase"`
		Institutions []string `json:"institutions"`
		Funding      []string `json:"funding"`
	} `json:"data"`
}

// Vaccine represents the vaccine info widget.
type Vaccine struct {
	VaccineData
	PhaseWidget      *widgets.Paragraph
	CandidatesWidget *widgets.List
}

// FetchData retrieves the latest vaccine information.
func (v *Vaccine) FetchData() error {
	url := "https://disease.sh/v3/covid-19/vaccine"
	return fetch(url, v)
}

func phaseWidget() *widgets.Paragraph {
	p := message.NewPrinter(language.English)

	widget := widgets.NewParagraph()
	widget.Title = "✅ Testing and approval process"
	widget.Text = p.Sprintf("[Preclinical testing](fg:cyan): Vaccine being tested on animals\n")
	widget.Text += p.Sprintf("[Phase 1 Safety Trials](fg:red): Vaccine being tested for safety and dosage\n")
	widget.Text += p.Sprintf("[Phase 2 Expanded Trials](fg:magenta): Vaccine in expanded safety trials\n")
	widget.Text += p.Sprintf("[Phase 3 Efficacy Trials](fg:blue): Vaccine in large-scale efficacy tests\n")
	widget.Text += p.Sprintf("[Limited Approval](fg:yellow): Vaccine approved for early or limited use\n")
	widget.Text += p.Sprintf("[Approved](fg:green): Vaccine approved for full use")
	widget.SetRect(0, 0, 50, 10)
	widget.BorderStyle.Fg = ui.ColorYellow
	widget.TitleStyle = ui.NewStyle(ui.ColorClear)
	widget.TextStyle = ui.NewStyle(ui.ColorClear)

	return widget
}

func candidatesWidget(v *Vaccine) *widgets.List {
	p := message.NewPrinter(language.English)

	widget := widgets.NewList()
	widget.Title = "🔥 Candidates (use j/k to scroll)"

	rows := []string{}

	for _, value := range v.Data {
		str := p.Sprintf(`/* [%s (%s)](fg:yellow)
========================================================================== */
		`, strings.ToUpper(value.Candidate), value.TrialPhase)
		if len(value.Sponsors) > 0 && value.Sponsors[0] != "" {
			str += p.Sprintf("Sponsors     => %s\n", strings.Join(value.Sponsors, ", "))
		}

		if len(value.Institutions) > 0 && value.Institutions[0] != "" {
			str += p.Sprintf("Institutions => %s\n", strings.Join(value.Institutions, ", "))
		}

		if len(value.Funding) > 0 && value.Funding[0] != "" {
			str += p.Sprintf("Funding      => %s\n", strings.Join(value.Funding, ", "))
		}

		str += "\n" + value.Details + "\n"
		rows = append(rows, str)
	}

	widget.Rows = rows

	widget.SelectedRowStyle.Fg = ui.ColorClear
	widget.TextStyle = ui.NewStyle(ui.ColorClear)
	widget.WrapText = true

	return widget
}

// Construct creates the vaccine widget using the VaccineData.
func (v *Vaccine) Construct() {
	v.PhaseWidget = phaseWidget()
	v.CandidatesWidget = candidatesWidget(v)
}
