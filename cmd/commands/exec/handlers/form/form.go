package execFormHandlers

import (
	"fmt"
	"os"

	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
	"github.com/arthurbcp/kuma/v2/pkg/style"
	"github.com/charmbracelet/huh"
)

func HandleForm(formData map[string]interface{}, vars map[string]interface{}) {
	data := vars["data"].(map[string]interface{})
	huhFields := []huh.Field{}
	title, err := execBuilders.BuildStringValue("title", formData, vars, false, constants.FormComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	description, err := execBuilders.BuildStringValue("description", formData, vars, false, constants.FormComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}
	accessibility, err := execBuilders.BuildBoolValue("accessibility", formData, vars, false, constants.FormComponent)
	if err != nil {
		style.ErrorPrint(err.Error())
		os.Exit(1)
	}

	if _, ok := formData["fields"]; !ok {
		style.ErrorPrint("fields is required")
		os.Exit(1)
	}
	for _, field := range formData["fields"].([]interface{}) {
		fieldMap, ok := field.(map[string]interface{})
		if !ok {
			style.ErrorPrint("invalid field map")
			os.Exit(1)
		}
		for key, value := range fieldMap {
			if value, ok := value.(map[string]interface{}); ok {
				switch key {
				case "select":
					huhField, out, outValue := HandleSelect(value, vars)
					huhFields = append(huhFields, huhField)
					data[out] = outValue
				case "input":
					huhField, out, outValue := HandleInput(value, vars)
					huhFields = append(huhFields, huhField)
					data[out] = outValue
				case "multi-select":
					huhField, out, outValue := HandleMultiSelect(value, vars)
					huhFields = append(huhFields, huhField)
					data[out] = outValue
				case "text":
					huhField, out, outValue := HandleText(value, vars)
					huhFields = append(huhFields, huhField)
					data[out] = outValue
				case "confirm":
					huhField, out, outValue := HandleConfirm(value, vars)
					huhFields = append(huhFields, huhField)
					data[out] = outValue
				default:
					fmt.Println("invalid field type: " + key)
				}
			} else {
				fmt.Println("invalid field type: " + key)
				os.Exit(1)
			}
		}
	}
	form := huh.NewForm(
		huh.NewGroup(huhFields...).
			Title(title).
			Description(description),
	)
	form.WithTheme(style.KumaTheme())
	form.WithAccessible(accessibility)
	form.Run()
}
