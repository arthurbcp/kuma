// Package multiSelect provides functions that
// help define and draw a multi-select step
package multiSelectInput

import (
	"fmt"

	"github.com/arthurbcp/kuma-cli/cmd/program"
	"github.com/arthurbcp/kuma-cli/cmd/steps"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	tea "github.com/charmbracelet/bubbletea"
)

// A Selection represents a choice made in a multiSelect step
type Selection struct {
	Choices map[string]bool
}

// Update changes the value of a Selection's Choice
func (s *Selection) Update(optionName string, value bool) {
	s.Choices[optionName] = value
}

// A multiSelect.model contains the data for the multiSelect step.
//
// It has the required methods that make it a bubbletea.Model
type model struct {
	cursor   int
	options  []steps.Item
	selected map[int]struct{}
	choices  *Selection
	header   string
	program  program.Program
}

func (m model) Init() tea.Cmd {
	return nil
}

// InitialModelMulti initializes a multiSelect step with
// the given data
func InitialMultiSelectInputModel(options []steps.Item, selection *Selection, header string, program *program.Program) model {
	return model{
		options:  options,
		selected: make(map[int]struct{}),
		choices:  selection,
		header:   style.TitleStyle.Render(header),
		program:  *program,
	}
}

// Update is called when "things happen", it checks for
// important keystrokes to signal when to quit, change selection,
// and confirm the selection.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.program.Exit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "y":
			for selectedKey := range m.selected {
				m.choices.Update(m.options[selectedKey].Value, true)
				m.cursor = selectedKey
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

// View is called to draw the multiSelect step
func (m model) View() string {
	s := m.header + "\n\n"

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = style.FocusedStyle.Render(">")
			option.Label = style.SelectedItemStyle.Render(option.Label)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = style.FocusedStyle.Render("*")
		}

		label := style.FocusedStyle.Render(option.Label)

		s += fmt.Sprintf("%s [%s] %s\n\n", cursor, checked, label)
	}

	s += fmt.Sprintf("Press %s to confirm choice.\n", style.FocusedStyle.Render("y"))
	s += fmt.Sprintf("Press %s to quit.\n", style.FocusedStyle.Render("q"))
	s += "\n"
	return s
}
