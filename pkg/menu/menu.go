package menu

import (
	"fmt"
	"io"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const menuWidth = 50

// return a paragraph object with given initial text.
func newParagraph(initText string, border bool, location int, wid int, ht int) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = initText
	p.Border = border
	p.SetRect(0, location, wid, location+ht)
	p.TextStyle.Fg = ui.ColorWhite
	return p
}

func processInput(input, warning *widgets.Paragraph, isValid func(string) (string, bool), uiEvents <-chan ui.Event) (string, error) {
	// keep tracking all input from user
	for {
		e := <-uiEvents
		if e.Type != ui.KeyboardEvent {
			continue
		}
		switch e.ID {
		case "<C-d>":
			return "", io.EOF
		case "<Enter>":
			warningWords, ok := isValid(input.Text)
			if ok {
				return input.Text, nil
			}
			input.Text = ""
			warning.Text = warningWords
			ui.Render(input)
			ui.Render(warning)
		case "<Backspace>":
			if len(input.Text) == 0 {
				continue
			}
			input.Text = input.Text[:len(input.Text)-1]
			ui.Render(input)
		default:
			if warning.Text != "" {
				warning.Text = ""
				ui.Render(warning)
			}
			input.Text += e.ID
			ui.Render(input)
		}
	}
}

// present an input box to user and return the user's input.
// GetInput will check validation of input using isValid function.
func GetInput(introwords string, location int, wid int, ht int, isValid func(string) (string, bool)) (string, error) {
	intro := newParagraph(introwords, false, location, len(introwords)+4, 3)
	location += 2
	input := newParagraph("", true, location, wid, ht+2)
	location += ht + 2
	warning := newParagraph("", false, location, wid, 3)

	ui.Render(intro)
	ui.Render(input)
	ui.Render(warning)

	uiEvents := ui.PollEvents()
	return processInput(input, warning, isValid, uiEvents)
}

// create a new ui window and display an input box.
func NewCustomInputWindow(introwords string, wid int, ht int, isValid func(string) (string, bool)) (string, error) {
	uiEvents := ui.PollEvents()
	return internalNewInputWindow(introwords, wid, ht, isValid, uiEvents)
}

func internalNewInputWindow(introwords string, wid int, ht int, isValid func(string) (string, bool), uiEvents <-chan ui.Event) (string, error) {
	if err := ui.Init(); err != nil {
		return "", fmt.Errorf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	location := 0
	intro := newParagraph(introwords, false, location, len(introwords)+4, 3)
	location += 2
	input := newParagraph("", true, location, wid, ht+2)
	location += ht + 2
	warning := newParagraph("", false, location, wid, 3)

	ui.Render(intro)
	ui.Render(input)
	ui.Render(warning)

	return processInput(input, warning, isValid, uiEvents)
}
