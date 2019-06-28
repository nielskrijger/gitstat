package internal

type FileChange struct {
	Name      string `json:"filepath"`
	IsBinary  bool   `json:"isBinary"`
	Additions int    `json:"additions"`
	Deletions int    `json:"deletions"`
}

func NewFileChange(path string) *FileChange {
	return &FileChange{
		Name: path,
	}
}
