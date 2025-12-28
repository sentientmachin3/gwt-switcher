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
	basePath := "."
	if insideWorktree {
		basePath = ".."
	}
	repo, err := git.PlainOpen(basePath)
	if err == git.ErrRepositoryNotExists {
		fmt.Println("ERR: no .git directory detected")
		os.Exit(1)
	}
	refs := BranchNames(repo)
	fuzzySelect(refs)
}

func fuzzySelect(refs []*plumbing.Reference) *plumbing.Reference {
	idx, _ := fzf.Find(refs, func(i int) string {
		return refs[i].Name().String()
	})
	return refs[idx]
}
