package ui

import (
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"
)

func Select(label string, items []string) (int, string, error) {
	templates := promptui.SelectTemplates{
		Active:   `➜ {{ . | cyan | bold }}`,
		Inactive: `   {{ . | cyan }}`,
	}
	prompt := promptui.Select{
		Label:     label,
		Items:     items,
		Templates: &templates,
	}

	return prompt.Run()
}

func Validate(label, key string) (string, error) {

	validate := getValidationFunction(key)
	templates := &promptui.PromptTemplates{
		Prompt:  `{{ . }} `,
		Valid:   `{{ . | green }} `,
		Invalid: `{{ . | red }} `,
		// Success:         `{{ "✔" | bold | green }} `,
		ValidationError: `{{ . | yellow }}`,
	}

	prompt := promptui.Prompt{
		Label:     label,
		Validate:  validate,
		Templates: templates,
	}

	return prompt.Run()
}

func getValidationFunction(key string) func(input string) error {

	switch key {

	case "email":
		return func(e string) error {
			var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
			if !emailRegex.MatchString(e) {
				return fmt.Errorf("Invalid %s", key)
			}
			return nil
		}

	default:
		return func(input string) error {
			if len(input) < 3 {
				return fmt.Errorf("%s can't be less than 3 characters", key)
			}
			return nil
		}

	}
}
