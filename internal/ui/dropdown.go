package ui

import (
	"github.com/derailed/tview"
	"github.com/gdamore/tcell/v2"
)

type DropDown struct {
	*tview.DropDown
	label      string
	options    []string
}

func NewDropDown(label string, options []string) *DropDown {
	d := DropDown{
		DropDown:   tview.NewDropDown(),
		label:      label,
		options:    options,
	}
	d.build()
	return &d
}

func (d *DropDown) SetSelectedFn(selectedFn DropdownSelectedFn) {
	d.SetSelectedFunc(func(text string, index int) {
		selectedFn(text, index)
	})
}

func (d *DropDown) build() {
	d.SetLabel(d.label)
	d.SetLabelColor(tcell.ColorOrange)
	d.SetOptions(d.options, func(text string, index int) {})
	d.SetCurrentOption(0)
	d.SetBorderPadding(0, 0, 0, 0)
	//profileDropdown.SetBorder(true)
}
