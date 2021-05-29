package internal

import (
	"errors"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Commits []*Commit

func (c Commits) ParseFileChanges() error {
	bar := NewProgressBar("Comparing file changes", int64(len(c)))

	for _, commit := range c {
		bar.Increment()
		currentTree, err := commit.originalCommit.Tree()
		if err != nil {
			return err
		}

		toTree := &object.Tree{}
		if commit.originalCommit.NumParents() > 0 {
			// Only compare with first parent, same as go-git's patch.Stats()
			firstParent, err := commit.originalCommit.Parents().Next()
			if err != nil {
				return err
			}
			if firstParent == nil {
				return errors.New("unable to find parent")
			}

			parentCommit := c.Find(firstParent.Hash.String())

			toTree, err = parentCommit.originalCommit.Tree()
			if err != nil {
				return err
			}
		}

		changes, err := toTree.Diff(currentTree)
		if err != nil {
			return err
		}

		fcs, err := NewFileChanges(changes)
		if err != nil {
			return err
		}

		commit.FileChanges = fcs
	}
	return nil
}

func (c Commits) Find(hash string) *Commit {
	for _, commit := range c {
		if commit.Hash == hash {
			return commit
		}
	}
	return nil
}
