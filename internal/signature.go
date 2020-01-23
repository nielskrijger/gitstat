package internal

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"time"
)

type Signature struct {
	Name  string    `json:"name"`
	Email string    `json:"-"` // don't include email, it's redundant PI
	When  time.Time `json:"time"`
}

func NewSignature(s object.Signature) Signature {
	return Signature{
		Name:  s.Name,
		Email: s.Email,
		When:  s.When,
	}
}
