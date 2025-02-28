package utilities

import (
	"html/template"
	"os"
)

func BuildTemplate(templatePath string, outputPath string, data any) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer output.Close()

	err = tmpl.Execute(output, data)

	return nil
}
