package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/henomis/lingoose/llm/openai"
	"github.com/henomis/lingoose/thread"
	"github.com/waigani/diffparser"
)

func main() {
	// Define flags
	verbose := flag.Bool("verbose", false, "Enable verbose mode")
	temperature := flag.Float64("temp", 0.0, "OpenAI temperature")
	model := flag.String("model", "gpt-4o", "OpenAI model (gpt-3.5, gpt-4, gpt-4o)")
	maxTokens := flag.Int("maxTokens", 2000, "OpenAI max tokens")

	// Parse command-line flags
	flag.Parse()

	input, err := getInput(flag.Arg(0))
	if err != nil {
		fmt.Println("No input provided")
		return
	}

	var r *Review

	openaillm := openai.New().WithModel(openai.Model(*model)).WithTemperature(float32(*temperature)).WithMaxTokens(*maxTokens)
	openaillm.WithToolChoice(newStr("reviewCode"))
	err = openaillm.BindFunction(
		func(review Review) string {
			r = &review
			return "Review received"
		},
		"reviewCode",
		"use this function to submit a code review",
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	t := thread.New().AddMessage(
		thread.NewSystemMessage().AddContent(thread.NewTextContent(systemPrompt))).
		AddMessage(thread.NewUserMessage().AddContent(thread.NewTextContent("Please review the code snippet below and provide your feedback:\n\n" + string(input))))

	err = openaillm.Generate(context.Background(), t)
	if err != nil {
		fmt.Println(err)
		return
	}

	if t.LastMessage().Role != thread.RoleTool || r == nil {
		fmt.Println("No review received")
		return
	}

	diff, err := diffparser.Parse(string(input))
	if err != nil {
		fmt.Println(err)
		return
	}

	printMarkdownOutput(diff, r, *verbose)
}
