package ui

import (
	"github.com/fatih/color"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

/*Draw function to draw the rectangle list of commits */
func Draw(items []string) int {
	if err := ui.Init(); err != nil {
		color.HiRed(err.Error())
	}
	defer ui.Close()
	var row int

	l := widgets.NewList()
	l.Title = "Commits (Press enter to select the commit)"
	l.Rows = items
	l.TextStyle = ui.NewStyle(ui.ColorBlue)
	l.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)
	l.WrapText = false
	l.SetRect(0, 0, 100, 10)

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			row = 0
			return row
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		case "<Enter>":
			row = l.SelectedRow + 1
			return row
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}
		ui.Render(l)
	}
}
