package internal

import (
	"gopkg.in/src-d/go-git.v4/plumbing/format/diff"
	"strings"
)

type FileChange struct {
	Name         string  `json:"filepath"`
	IsBinary     bool    `json:"isBinary"`
	Additions    int     `json:"additions"`
	Deletions    int     `json:"deletions"`
	RawAdditions int     `json:"rawAdditions"`
	RawDeletions int     `json:"rawDeletions"`
	RenameFrom   string  `json:"renameOf,omitempty"`
	RenameTo     string  `json:"renameTo,omitempty"`
	Similarity   float32 `json:"similarity,omitempty"`
}

// Based heavily on:
// https://github.com/src-d/go-git/blob/d6c4b113c17a011530e93f179b7ac27eb3f17b9b/plumbing/object/patch.go
func NewFileChange(name string, fp diff.FilePatch) *FileChange {
	fc := &FileChange{Name: name, IsBinary: fp.IsBinary()}

	for _, chunk := range fp.Chunks() {
		s := chunk.Content()
		if len(s) == 0 {
			continue
		}

		switch chunk.Type() {
		case diff.Add:
			fc.Additions += strings.Count(s, "\n")
			if s[len(s)-1] != '\n' {
				fc.Additions++
			}
		case diff.Delete:
			fc.Deletions += strings.Count(s, "\n")
			if s[len(s)-1] != '\n' {
				fc.Deletions++
			}
		}
	}

	// Additions & Deletions are mutable, Raw* are not after this
	fc.RawAdditions = fc.Additions
	fc.RawDeletions = fc.Deletions

	return fc
}
