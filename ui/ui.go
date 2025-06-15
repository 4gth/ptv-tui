package ui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type RouteRow struct {
	Name    string
	Number  string
	RouteID int32
}

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "q":
			return m, tea.Quit

		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func StartNewTea(route []RouteRow) error {

	var rows []table.Row
	for _, r := range route {
		rows = append(rows, table.Row{r.Name, r.Number, fmt.Sprintf("%d", r.RouteID)})
	}

	colums := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Number", Width: 40},
		{Title: "RouteID", Width: 10},
	}

	t := table.New(
		table.WithColumns(colums),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
