package execHandlers

import "github.com/charmbracelet/huh"

func HandleConfirm(input map[string]interface{}, vars map[string]interface{}) {
	huh.NewConfirm().
		Title("Are you sure?").
		Affirmative("Yes!").
		Negative("No.").
		Value(&confirm)
}
