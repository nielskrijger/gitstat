package internal

import (
	"errors"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

type Commits []*Commit

func (c Commits) ParseFileChanges() error {
	for _, commit := range c {
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

		changes, err := currentTree.Diff(toTree)
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
