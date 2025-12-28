package main

import (
	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
)

func InsideWorktree() bool {
	return true
}

func BranchNames(repo *git.Repository) []*plumbing.Reference {
	iter, _ := repo.Branches()
	branches := make([]*plumbing.Reference, 0)
	iter.ForEach(func(ref *plumbing.Reference) error {
		branches = append(branches, ref)
		return nil
	})
	return branches
}
