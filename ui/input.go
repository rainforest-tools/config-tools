package ui

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/rainforest-tools/config-tools/models"
)

func GetInput(pc models.PromptContent, validates ...func(string) error) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.Error)
		}
		for _, v := range validates {
			v(input)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func GetInt(pc models.PromptContent, validates ...func(string) error) int {
	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New(pc.Error)
		}
		return nil
	}
	str := GetInput(pc, append(validates, validate)...)
	result, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
