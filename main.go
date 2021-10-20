package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

// Pipelines struct
type Pipelines struct {
	Pipelines []Pipeline `json:"pipelines"`
}

// Pipeline struct
type Pipeline struct {
	RepositoryName     string `json:"repositoryName"`
	Environment        string `json:"environment"`
	Stack              string `json:"stack"`
	PipelineName       string `json:"pipelineName"`
	Parameterized      bool   `json:"parameterized"`
	PipelineParameters string `json:"pipelineParameters"`
	QueueTimeout       int64  `json:"queueTimeout"`
	BuildTimeout       int64  `json:"buildTimeout"`
}

// main is the entry point
func main() {
	pipelineFile := "pipelines.json"
	pipelineTemplateFile := "pipeline_template.yml"
	pipelineRenderFile := "pipeline_render.yml"

	pipelineData, err := os.Open(pipelineFile)
	errorChecking(err)
	defer pipelineData.Close()
	fmt.Printf("Successfully Opened %s\n", pipelineFile)

	byteValue, _ := ioutil.ReadAll(pipelineData)

	var ps Pipelines

	json.Unmarshal(byteValue, &ps)

	for _, p := range ps.Pipelines {
		fmt.Printf("Repository Name: %s", p.RepositoryName)
		fmt.Printf("\nEnvironment: %s", p.Environment)
		fmt.Printf("\nStack: %s", p.Stack)
		fmt.Printf("\nPipeline Name: %s", p.PipelineName)
		fmt.Printf("\nParameterized: %t", p.Parameterized)
		fmt.Printf("\nPipeline Parameters: %s", p.PipelineParameters)
		fmt.Printf("\nQueue Timeout: %d", p.QueueTimeout)
		fmt.Printf("\nBuild Timeout: %d\n", p.BuildTimeout)

		fmt.Printf("Successfully Opened %s\n", pipelineTemplateFile)
		t, err := template.ParseFiles(pipelineTemplateFile)
		errorChecking(err)
		err = t.Execute(os.Stdout, p)
		errorChecking(err)
		f, err := os.OpenFile(pipelineRenderFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer f.Close()
		errorChecking(err)
		err = t.Execute(f, p)
		errorChecking(err)
		_, err = f.WriteString("\n")
		errorChecking(err)
	}

}

func errorChecking(err error) {
	if err != nil {
		panic(err)
	}
}
