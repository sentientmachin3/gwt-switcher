package main

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	fzf "github.com/ktr0731/go-fuzzyfinder"
)

func main() {
	insideWorktree := InsideWorktree()
	isBare := false
	if insideWorktree {
		isBare = true
		os.Chdir("..")
	}
	repo, err := git.PlainOpen(".")
	if err == git.ErrRepositoryNotExists {
		fmt.Println("ERR: no .git directory detected")
		os.Exit(1)
	}
	FetchRefs(repo)
	refs := BranchNames(repo)
	selected := fuzzySelect(refs)

	if isBare {
		// this library does not well support worktrees...?
		AddWorktree(selected)
	} else {
		// checkout branch
		worktree, _ := repo.Worktree()
		err := worktree.Checkout(&git.CheckoutOptions{Branch: selected.Name()})
		if err != nil {
			fmt.Printf("ERR: unable to checkout worktree, %v", err)
		}

	}
}

func fuzzySelect(refs []*plumbing.Reference) *plumbing.Reference {
	idx, _ := fzf.Find(refs, func(i int) string {
		return refs[i].Name().String()
	})
	return refs[idx]
}
