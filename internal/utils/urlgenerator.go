package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
)

const (
	templatePath = "internal/utils/templates/url.tpl"
)

func CreateMetricURL(params map[string]any) (string, error) {
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("cannot open template file: %v", err)
	}

	var buf bytes.Buffer

	if err = tpl.Execute(&buf, params); err != nil {
		return "", fmt.Errorf("execute template error: %s", err)
	}

	rb, err := io.ReadAll(&buf)
	if err != nil {
		return "", fmt.Errorf("cannot read buffer: %s", err)
	}

	url := string(rb)
	return url, nil
}
