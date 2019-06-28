package internal

import (
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"strings"
)

type Commit struct {
	Hash             string        `json:"hash"`
	Author           Signature     `json:"author"`
	Committer        Signature     `json:"committer"`
	Message          string        `json:"message"`
	Files            []*FileChange `json:"files"`
	IsMerge          bool          `json:"isMerge"`
}

func NewCommit(c *object.Commit) (*Commit, error) {
	r := &Commit{
		Hash:             c.Hash.String(),
		Author:           NewSignature(c.Author),
		Committer:        NewSignature(c.Committer),
		Message:          strings.TrimSpace(c.Message),
		Files:            make([]*FileChange, 0),
		IsMerge:          len(c.ParentHashes) > 1,
	}

	// get file stats, these will be added later
	stats, err := c.Stats()
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve stats")
	}

	// loop through all files and store metadata in a "FileChange"
	// find stat for this file and add change
	for _, stat := range stats {
		fc := NewFileChange(stat.Name)

		fc.Additions = stat.Addition
		fc.Deletions = stat.Deletion

		r.Files = append(r.Files, fc)
	}

	return r, nil
}
