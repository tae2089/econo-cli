package tui

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tableModel struct {
	table table.Model
}

func InitalTableModel(columns []table.Column, rows []table.Row) tableModel {
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(len(rows)),
		table.WithFocused(false),
	)
	// fmt.Println(t.Focused())
	s := table.DefaultStyles()
	s.Selected = lipgloss.NewStyle().UnsetBold()
	t.SetStyles(s)
	return tableModel{table: t}
}

func (m tableModel) Init() tea.Cmd {

	return nil
}

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m tableModel) View() string {
	var baseStyle = lipgloss.NewStyle()

	return baseStyle.Render(m.table.View()) + "\n"
}
