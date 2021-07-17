/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"html/template"
	"os"

	"github.com/rainforest-tools/config-tools/models"
	"github.com/rainforest-tools/config-tools/ui"
	"github.com/spf13/cobra"
)

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		job := CreateJob()
		CreateJobFile("job.sh", job)
	},
}

func init() {
	rootCmd.AddCommand(jobCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func CreateJob() models.Job {
	job := models.Job{}

	job.Name = ui.GetInput(models.PromptContent{
		Error: "",
		Label: "What's the name of job?",
	})

	job.Partition = ui.GetInput(models.PromptContent{
		Error: "",
		Label: "Which partition do you want to use?",
	})

	job.GPU = ui.GetInt(models.PromptContent{
		Error: "",
		Label: "How many gpus do you want to use?",
	})

	job.Modules = ui.GetMultiSelect(models.PromptContent{
		Error: "",
		Label: "Choose environment modules",
	}, []string{"gcc", "opt", "python"})

	job.Command = ui.GetInput(models.PromptContent{
		Error: "",
		Label: "Enter your command",
	})

	return job
}

func CreateJobFile(path string, job models.Job) {
	t, err := template.ParseFiles("./static/templates/job.tpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	err = t.Execute(f, job)
	if err != nil {
		fmt.Println(err)
	}
}
