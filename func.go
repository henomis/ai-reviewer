package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/waigani/diffparser"
)

func newStr(str string) *string {
	return &str
}

type Review struct {
	Language string  `json:"language" jsonschema:"description=Programming language"`
	Issues   []Issue `json:"issues" jsonschema:"description=List of issues"`
}

type Issue struct {
	File        string `json:"file" jsonschema:"description=File path"`
	Code        string `json:"code" jsonschema:"description=Code snippet of the issue"`
	Description string `json:"description" jsonschema:"description=Description of the issue in markdown format"`
	Suggestion  string `json:"suggestion" jsonschema:"description=Suggested fix for the issue in markdown format"`
}

func getInput(filepath string) (string, error) {
	if filepath != "" {
		input, err := os.ReadFile(filepath)
		if err != nil {
			return "", err
		}

		return string(input), nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	result := ""
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}
	return result, nil
}

func printMarkdownOutput(diff *diffparser.Diff, r *Review, verbose bool) {
	for _, file := range diff.Files {
		fmt.Println("# " + file.NewName + "\n")

		if verbose {
			for _, hunk := range file.Hunks {
				fmt.Printf("```%s\n", r.Language)
				for _, line := range hunk.NewRange.Lines {
					fmt.Printf("%d\t%s\n", line.Number, line.Content)
				}
				fmt.Printf("```\n\n---\n\n")
			}
		}

		fmt.Printf("\n\n")

		for i, issue := range r.Issues {
			if issue.File == file.NewName {
				fmt.Printf("### Issue #%d\n\n", i+1)
				fmt.Printf("```%s\n", r.Language)
				fmt.Println(issue.Code)
				fmt.Println("```")

				fmt.Printf("%s\n\n", issue.Description)

				fmt.Printf("### Fix\n\n")
				fmt.Printf("%s\n\n---\n\n", issue.Suggestion)
			}
		}
	}
}
