package internal

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"strings"
)

type Commit struct {
	Hash        string      `json:"hash"`
	Author      Signature   `json:"author"`
	Committer   Signature   `json:"committer"`
	Message     string      `json:"message"`
	FileChanges FileChanges `json:"files"`
	IsMerge     bool        `json:"isMerge"`

	originalCommit *object.Commit
}

func NewCommit(c *object.Commit) *Commit {
	return &Commit{
		Hash:        c.Hash.String(),
		Author:      NewSignature(c.Author),
		Committer:   NewSignature(c.Committer),
		FileChanges: make(FileChanges, 0),
		Message:     strings.TrimSpace(c.Message),
		IsMerge:     len(c.ParentHashes) > 1,

		originalCommit: c,
	}
}
