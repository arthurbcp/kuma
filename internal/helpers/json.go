package helpers

import (
	"bytes"
	"encoding/json"
)

func PrettyJson(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func PrettyMarshal(data interface{}) (string, error) {
	j, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return PrettyJson(string(j)), nil
}
