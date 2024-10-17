package execHandlers

import (
	"fmt"
)

func HandleForm(fields []interface{}, vars map[string]interface{}) {
	for _, field := range fields {
		fieldMap := field.(map[string]interface{})
		for key, value := range fieldMap {
			if key == "select" {
				HandleSelect(value.(map[string]interface{}), vars)
			} else if key == "input" {
				HandleInput(value.(map[string]interface{}), vars)
			} else {
				fmt.Println("invalid field type: " + key)
			}
		}
	}
}
