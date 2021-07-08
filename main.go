package main

import "log"
import ui "github.com/gizak/termui/v3"

import "github.com/gizak/termui/v3/widgets"

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("%v", err)
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Text = "Hello World!"
	p.SetRect(0, 0, 25, 5)

	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
