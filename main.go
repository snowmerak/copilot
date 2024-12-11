package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"

	"github.com/snowmerak/copilot/prompt"
	"github.com/snowmerak/copilot/qwen"
	"github.com/snowmerak/copilot/template"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := qwen.New(qwen.Config{})
	if err != nil {
		panic(err)
	}

	filePath := ""
	if err := survey.AskOne(&survey.Input{
		Message: "Input Output File path or [[stdout]]",
	}, &filePath); err != nil {
		panic(err)
	}

	writer := (*bufio.Writer)(nil)

	switch filePath {
	case "[[stdout]]":
		writer = bufio.NewWriter(os.Stdout)
	default:
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer = bufio.NewWriter(file)
	}

	temple := ""
	if err := survey.AskOne(&survey.Select{
		Message: "Select Template",
		Options: template.List(),
	}, &temple); err != nil {
		panic(err)
	}

	query := ""
	if err := survey.AskOne(&survey.Multiline{
		Message: "Input Request",
	}, &query); err != nil {
		panic(err)
	}

	pmt := prompt.Make(template.Get(temple), query)

	log.Printf("Start generating code from template %s", temple)
	resp, err := cli.Generate(ctx, pmt)
	if err != nil {
		panic(err)
	}

	log.Printf("Code generated successfully")

	if _, err := writer.WriteString(resp); err != nil {
		panic(err)
	}

	if err := writer.Flush(); err != nil {
		panic(err)
	}

	log.Printf("File %s created", filePath)
}
