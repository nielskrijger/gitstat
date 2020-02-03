package internal

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
	project, err := NewProject(filepath)
	if err != nil {
		return err
	}

	err = project.ParseCommits()
	if err != nil {
		return err
	}

	p.Projects = append(p.Projects, project)

	return nil
}
