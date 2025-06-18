package ui

import (
	"fmt"
	"log"
	"ptv-tui/api"
	"reflect"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

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

// finds column name tags and constructs column names
func columnsFromStruct(s any) []table.Column {
	t := reflect.TypeOf(s)

	var cols []table.Column

	for i := range t.NumField() {
		field := t.Field(i)

		colName := field.Tag.Get("col")
		if colName == "" {
			colName = field.Name
		}

		cols = append(cols, table.Column{
			Title: colName,
			Width: 20,
		})
	}
	return cols
}

// finds which columns have valid data to print for each row
func RowsToTable(d any) ([]table.Row, error) {
	val := reflect.ValueOf(d)
	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("data is not slice")
	}

	var rows []table.Row
	for i := range val.Len() {
		item := val.Index(i)
		if item.Kind() == reflect.Ptr {
			item = item.Elem()
		}

		if item.Kind() != reflect.Struct {
			return nil, fmt.Errorf("slice is not a struct")
		}

		var row table.Row
		for j := range item.NumField() {
			field := item.Field(j)
			row = append(row, fmt.Sprintf("%v", field.Interface()))
		}
		rows = append(rows, row)
	}
	log.Print(rows, d)
	return rows, nil
}

func StartNewTea(Rows *[]api.Responses) error {

	columns := columnsFromStruct(api.Responses{})
	rows, _ := RowsToTable(Rows)

	t := table.New(
		table.WithColumns(columns),
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
