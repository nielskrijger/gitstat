package internal

import (
	"fmt"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
	"os"
	"path/filepath"
)

type Project struct {
	Name     string    `json:"name"`
	Commits  []*Commit `json:"commits"`

	filepath string
}

func NewProject(path string) (*Project, error) {
	name, err := projectName(path)
	if err != nil {
		return nil, err
	}

	return &Project{
		Name: name,
		Commits: make([]*Commit, 0),
		filepath: path,
	}, nil
}

func (p *Project) ParseCommits() error {
	// Instantiate a new repository targeting the given path (the .git folder)
	fs := osfs.New(p.filepath)
	if _, err := fs.Stat(git.GitDirName); err == nil {
		fs, err = fs.Chroot(git.GitDirName)
		if err != nil {
			return err
		}
	}

	s := filesystem.NewStorageWithOptions(fs, cache.NewObjectLRUDefault(), filesystem.Options{KeepDescriptors: true})
	defer s.Close()

	r, err := git.Open(s, fs)
	if err != nil {
		return err
	}

	// ... retrieve the branch pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		return err
	}

	// ... retrieve the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return err
	}

	// ... iterate over the commits
	return cIter.ForEach(func(c *object.Commit) error {
		fmt.Print(".")
		commit, err := NewCommit(c)
		if err != nil {
			return err
		}
		p.Commits = append(p.Commits, commit)
		return nil
	})
}

func projectName(fp string) (string, error) {
	abs, err := filepath.Abs(fp)
	if err != nil {
		return "", err
	}
	for i := len(abs) - 1; i >= 0; i-- {
		if os.IsPathSeparator(abs[i]) {
			return abs[i+1:], nil
		}
	}
	return abs, nil
}
