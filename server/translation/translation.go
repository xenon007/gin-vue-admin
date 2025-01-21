package translation

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// NewTranslation initializes a Translation instance
func NewTranslation(path string, languages []string) (*Translation, error) {
	translations := make(map[string]map[string]string)
	for _, lang := range languages {
		filePath := fmt.Sprintf("%s/%s.json", path, lang)
		file, err := os.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to open translation file for %s: %w", lang, err)
		}
		defer file.Close()

		var messages map[string]string
		if err := json.NewDecoder(file).Decode(&messages); err != nil {
			return nil, fmt.Errorf("failed to decode translation file for %s: %w", lang, err)
		}
		translations[lang] = messages
	}
	return &Translation{Messages: translations}, nil
}

// Translate returns the translation for a given key and language
func (t *Translation) Translate(lang, key string, data map[string]string) (string, error) {
	if messages, ok := t.Messages[lang]; ok {
		if message, ok := messages[key]; ok {
			tmpl, err := template.New("translation").Parse(message)
			if err != nil {
				return "", fmt.Errorf("failed to parse template: %w", err)
			}
			var builder strings.Builder
			if err := tmpl.Execute(&builder, data); err != nil {
				return "", fmt.Errorf("failed to execute template: %w", err)
			}
			return builder.String(), nil
		}
		return "", errors.New("translation key not found")
	}
	return "", errors.New("language not supported")
}
