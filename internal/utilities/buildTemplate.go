package utilities

import (
	"html/template"
	"log"
	"os"
)

func BuildTemplate(templatePath string, outputPath string, data any) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Error parsing template %s: %v", templatePath, err)
	}

	output, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creating output file %s: %v", outputPath, err)
	}

	defer output.Close()

	err = tmpl.Execute(output, data)
	if err != nil {
		log.Fatalf("Error executing template %s: %v", templatePath, err)
	}
}
