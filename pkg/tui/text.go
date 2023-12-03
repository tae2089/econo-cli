package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// A model can be more or less any type of data. It holds all the data for a
// program, so often it's a struct. For this simple example, however, all
// we'll need is a simple integer.
type textOutputModel struct {
	text string
}

func InitTextOutputModel(text string) textOutputModel {
	return textOutputModel{text: text}
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m textOutputModel) Init() tea.Cmd {
	return nil
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m textOutputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

// View returns a string based on data in the model. That string which will be
// rendered to the terminal.
func (m textOutputModel) View() string {
	return fmt.Sprintf("%s", m.text)
}
