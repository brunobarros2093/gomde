package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	a := app.New()
	win := a.NewWindow("GOMDE - Markdown")
	edit, preview := cfg.makeUI()
	win.SetContent(container.NewHSplit(edit, preview))
	win.Resize(fyne.NewSize(400, 400))
	win.CenterOnScreen()
	win.ShowAndRun()
}

func (cfg *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	cfg.EditWidget = edit
	cfg.PreviewWidget = preview
	// the right panel gonna show the edited version
	edit.OnChanged = preview.ParseMarkdown
	return edit, preview
}
