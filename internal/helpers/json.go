package helpers

import (
	"bytes"
	"encoding/json"
)

func (h *Helpers) PrettyJson(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func (h *Helpers) PrettyMarshal(data interface{}) (string, error) {
	j, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return h.PrettyJson(string(j)), nil
}
