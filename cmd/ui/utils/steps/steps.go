// Package steps provides utility for creating
// each step of the CLI
package steps

import "strings"

// A StepSchema contains the data that is used
// for an individual step of the CLI
type StepSchema struct {
	StepName string // The name of a given step
	Options  []Item // The slice of each option for a given step
	Headers  string // The title displayed at the top of a given step
	Field    string
}

// Steps contains a slice of steps
type Steps struct {
	Steps map[string]StepSchema
}

// An Item contains the data for each option
// in a StepSchema.Options
type Item struct {
	Label, Value, Description, Tags string
}

func NewItem(label, value, description string, tags []string) Item {
	if description != "" {
		description = "\n\t\t" + description
	}
	tagsStr := ""
	if len(tags) > 0 {
		tagsStr = "\n\t\ttags: " + strings.Join(tags, ", ")
	}
	return Item{
		Label:       label,
		Value:       value,
		Description: description,
		Tags:        tagsStr,
	}
}
