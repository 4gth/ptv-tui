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

// DepaturesForStop represents the structure of the departure information for a stop.
type DepaturesForStop struct {
	StopName     string `col:"Stop Name"`
	Number       string `col:"Route Number"`
	Direction    int32  `col:"Direction"`
	EstDeparture string `col:"Estimated Departure"`
	SchDeparture string `col:"Scheduled Departure"`
}

// NewStyle is a function that creates a new style for the table.
var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

// model is the main model for the terminal user interface.
type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

// Update processes messages and updates the model.
// It handles key messages for quitting the application and updates the table model.
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

// columnsFromStruct generates table columns based on the struct fields.
// It uses struct tags to determine the column titles.
// If no tag is found, it uses the field name as the column title.
// Each column has a fixed width of 15 characters.
// TODO: Make the width dynamic based on the content.
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
			Width: 15,
		})
	}
	return cols
}

// RowsToTable converts a slice of structs into a slice of table rows.
// It iterates over each struct, extracts the field values, and formats them into rows.
// If the input is not a slice or if the elements are not structs, it returns an error.
func RowsToTable(d any) ([]table.Row, error) {
	val := reflect.ValueOf(d)
	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("data is not slice")
	}
	log.Printf("making rows from: %+v", val)
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
	log.Printf("rows: %+v : %+v", rows, d)
	return rows, nil
}

// BrewDepatureTea creates a terminal user interface to display departure information for a stop.
func BrewDepatureTea(Rows []api.Responses) error {

	parsedResponses := make([]DepaturesForStop, len(Rows))
	for i, r := range Rows {
		parsedResponses[i] = DepaturesForStop{
			StopName:     r.StopName,
			Number:       r.Number,
			Direction:    r.DirectionID,
			EstDeparture: r.EstDepTime,
			SchDeparture: r.SchDepTime,
		}
	}
	columns := columnsFromStruct(DepaturesForStop{})
	rows, err := RowsToTable(parsedResponses)
	if err != nil {
		log.Printf("error converting rows: %v", err)
	}

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

	_, err = p.Run()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
