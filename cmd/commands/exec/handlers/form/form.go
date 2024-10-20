package execFormHandlers

import (
	"fmt"

	execBuilders "github.com/arthurbcp/kuma/v2/cmd/commands/exec/builders"
	"github.com/arthurbcp/kuma/v2/cmd/constants"
	"github.com/arthurbcp/kuma/v2/pkg/style"
	"github.com/charmbracelet/huh"
)

func HandleForm(formData map[string]interface{}, vars map[string]interface{}) error {
	data := vars["data"].(map[string]interface{})
	huhFields := []huh.Field{}
	title, err := execBuilders.BuildStringValue("title", formData, vars, false, constants.FormComponent)
	if err != nil {
		return err
	}
	description, err := execBuilders.BuildStringValue("description", formData, vars, false, constants.FormComponent)
	if err != nil {
		return err
	}
	accessibility, err := execBuilders.BuildBoolValue("accessibility", formData, vars, false, constants.FormComponent)
	if err != nil {
		return err
	}

	if _, ok := formData["fields"]; !ok {
		return fmt.Errorf("fields is required")
	}
	for _, field := range formData["fields"].([]interface{}) {
		fieldMap, ok := field.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid field map")
		}
		for key, value := range fieldMap {
			if value, ok := value.(map[string]interface{}); ok {
				switch key {
				case constants.SelectComponent:
					huhField, out, outValue, err := HandleSelect(value, vars)
					if err != nil {
						return fmt.Errorf("[field:%s] - %s", constants.SelectComponent, err.Error())
					}
					huhFields = append(huhFields, huhField)
					data[out] = outValue
				case constants.InputComponent:
					huhField, out, outValue, err := HandleInput(value, vars)
					if err != nil {
						return fmt.Errorf("[field:%s] - %s", constants.InputComponent, err.Error())
					}
					huhFields = append(huhFields, huhField)
					data[out] = outValue
				case constants.MultiSelectComponent:
					huhField, out, outValue, err := HandleMultiSelect(value, vars)
					if err != nil {
						return fmt.Errorf("[field:%s] - %s", constants.MultiSelectComponent, err.Error())
					}
					huhFields = append(huhFields, huhField)
					data[out] = outValue
				case constants.TextComponent:
					huhField, out, outValue, err := HandleText(value, vars)
					if err != nil {
						return fmt.Errorf("[field:%s] - %s", constants.TextComponent, err.Error())
					}
					huhFields = append(huhFields, huhField)
					data[out] = outValue
				case constants.ConfirmComponent:
					huhField, out, outValue, err := HandleConfirm(value, vars)
					if err != nil {
						return fmt.Errorf("[field:%s] - %s", constants.ConfirmComponent, err.Error())
					}
					huhFields = append(huhFields, huhField)
					data[out] = outValue
				default:
					return fmt.Errorf("invalid field type: %s", key)
				}
			} else {
				return fmt.Errorf("invalid field type: %s", key)
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
	err = form.Run()
	if err != nil {
		return fmt.Errorf("error running form: %s", err.Error())
	}

	return nil
}
