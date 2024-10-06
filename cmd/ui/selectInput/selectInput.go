// Package multiInput provides functions that
// help define and draw a multi-input step
package selectInput

import (
	"fmt"

	"github.com/arthurbcp/kuma-cli/cmd/steps"
	"github.com/arthurbcp/kuma-cli/cmd/ui/textInput"
	"github.com/arthurbcp/kuma-cli/pkg/style"
	tea "github.com/charmbracelet/bubbletea"
)

// A Selection represents a choice made in a multiInput step
type Selection struct {
	Choice string
}

// Update changes the value of a Selection's Choice
func (s *Selection) Update(value string) {
	s.Choice = value
}

// A multiInput.model contains the data for the multiInput step.
//
// It has the required methods that make it a bubbletea.Model
type model struct {
	cursor    int
	choices   []steps.Item
	selected  map[int]struct{}
	choice    *Selection
	header    string
	other     *bool
	skippable *bool
	exit      *bool
}

func (m model) Init() tea.Cmd {
	return nil
}

// InitialSelectInputModel initializes a multiInput step with
// the given data
func InitialSelectInputModel(choices []steps.Item, selection *Selection, header string, other, skippable, exit bool) model {
	m := model{
		choices:   choices,
		selected:  make(map[int]struct{}),
		choice:    selection,
		header:    style.TitleStyle.Render(header),
		exit:      &exit,
		other:     &other,
		skippable: &skippable,
	}
	return m
}

// Update is called when "things happen", it checks for
// important keystrokes to signal when to quit, change selection,
// and confirm the selection.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			*m.exit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if len(m.selected) == 1 {
				m.selected = make(map[int]struct{})
			}
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "y":
			if len(m.selected) == 1 {
				for selectedKey := range m.selected {
					m.choice.Update(m.choices[selectedKey].Value)
					m.cursor = selectedKey
				}
				return m, tea.Quit
			}
		case "o":
			if *m.other {
				textValue := &textInput.Output{}
				p := tea.NewProgram(textInput.InitialTextInputModel(textValue, "", false))
				_, err := p.Run()
				if err != nil {
					style.ErrorPrint("error running program: " + err.Error())
					*m.exit = true
					return m, tea.Quit
				}
				m.choice.Update(textValue.Output)
				return m, tea.Quit
			}
		case "s":
			if *m.skippable {
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

// View is called to draw the multiInput step
func (m model) View() string {
	s := m.header + "\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = style.FocusedStyle.Render(">")
			choice.Label = style.SelectedItemStyle.Render(choice.Label)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = style.FocusedStyle.Render("x")
		}
		label := style.FocusedStyle.Render(choice.Label)

		s += fmt.Sprintf("%s [%s] %s\n\n", cursor, checked, label)
	}

	s += fmt.Sprintf("Press %s to confirm choice.\n", style.FocusedStyle.Render("y"))
	if *m.other {
		s += fmt.Sprintf("Press %s to text another option.\n", style.FocusedStyle.Render("o"))
	}
	if *m.skippable {
		s += fmt.Sprintf("Press %s to skip.\n", style.FocusedStyle.Render("s"))
	}
	s += fmt.Sprintf("Press %s to quit.\n", style.FocusedStyle.Render("q"))
	s += "\n"
	return s
}
