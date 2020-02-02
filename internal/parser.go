package internal

import (
	"fmt"
	"time"
)

type Parser struct {
	Version  string     `json:"version"`
	Projects []*Project `json:"projects"`
}

func NewParser() *Parser {
	return &Parser{
		Version:  "1.0.0",
		Projects: make([]*Project, 0),
	}
}

func (p *Parser) ParseProject(filepath string) error {

	fmt.Printf("processing %q\n", filepath)
	project, err := NewProject(filepath)
	if err != nil {
		return err
	}

	start := time.Now()
	fmt.Print("parsing commits...")
	err = project.ParseCommits()
	if err != nil {
		return err
	}
	fmt.Printf(" done (%v)\n", time.Since(start).Round(time.Millisecond))

	p.Projects = append(p.Projects, project)

	return nil
}
