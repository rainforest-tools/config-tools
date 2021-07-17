package ui

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/manifoldco/promptui"
	"github.com/rainforest-tools/config-tools/models"
)

func GetSelect(pc models.PromptContent, options []string) string {
	if len(options) == 0 {
		fmt.Println("No Options")
		return ""
	}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label: pc.Label,
			Items: options,
		}

		index, result, err = prompt.Run()
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func GetMultiSelect(pc models.PromptContent, options []string) []string {
	var result []string

	prompt := &survey.MultiSelect{
		Message: pc.Label,
		Options: options,
	}
	survey.AskOne(prompt, &result)

	return result
}
