package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type JsonToEnvOption struct {
	FieldSeparator string
}

func mapToEnv(prefix *string, data *map[string]interface{}, options *JsonToEnvOption) (*string, error) {
	var b strings.Builder
	for key, element := range *data {
		child_prefix := key
		if prefix != nil {
			child_prefix = fmt.Sprintf("%s%s%s", *prefix, options.FieldSeparator, key)
		}
		switch v := (element).(type) {
		case string:
			b.WriteString(fmt.Sprintf("%s: \"%s\"", child_prefix, v))
		case bool:
			b.WriteString(fmt.Sprintf("%s: %t", child_prefix, v))
		case float64:
			// JSON unmarshall all number to float
			// so we need to check for integrals here
			intVal := int(v)
			if v == float64(intVal) {
				b.WriteString(fmt.Sprintf("%s: %d", child_prefix, intVal))
			} else {
				b.WriteString(fmt.Sprintf("%s: %f", child_prefix, v))
			}
		case map[string]interface{}:
			content, err := mapToEnv(&child_prefix, &v, options)
			if err != nil {
				return nil, err
			}
			b.WriteString(*content)
		}
		b.WriteString("\n")
	}
	// Trim blank lines
	out := regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(strings.TrimSpace(b.String()), "\n")
	return &out, nil
}

func JsonToEnv(buffer *[]byte, options *JsonToEnvOption) (*string, error) {
	var jsonObj map[string]interface{}
	err := json.Unmarshal(*buffer, &jsonObj)
	if err != nil {
		return nil, err
	}
	return mapToEnv(nil, &jsonObj, options)
}
