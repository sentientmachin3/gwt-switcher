package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"

	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
)

func InsideWorktree() bool {
	// parent folder has refs, HEAD and worktrees dirs/files
	entries, err := fs.ReadDir(os.DirFS(".."), ".")
	if err != nil {
		fmt.Printf("ERR: unable to detect if inside worktree, %v", err)
		os.Exit(1)
	}
	refsFound := false
	worktreeFound := false
	headFound := false
	for _, e := range entries {
		if e.Name() == "refs" {
			refsFound = true
		}
		if e.Name() == "HEAD" {
			headFound = true
		}
		if e.Name() == "worktrees" {
			worktreeFound = true
		}
	}
	return refsFound && worktreeFound && headFound
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

func FetchRefs(repo *git.Repository) {
	err := repo.Fetch(&git.FetchOptions{Force: true, Prune: true})
	if err != nil && err == git.NoErrAlreadyUpToDate {
		fmt.Println("Already up-to-date")
	} else if err != nil {
		fmt.Printf("ERR: unable to fetch, %v", err)
	}
}

func AddWorktree(ref *plumbing.Reference) {
	shortName := ref.Name().Short()
	trimmedName := strings.ReplaceAll(shortName, "/", "-")
	exec.Command("git", "worktree", "add", trimmedName, shortName)
}

