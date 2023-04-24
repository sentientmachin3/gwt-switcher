package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	fzf "github.com/ktr0731/go-fuzzyfinder"
)

func main() {
	var user_home, _ = os.UserHomeDir()
	var sessionizeFlag bool
	var rootFolderFlag string

	var help = flag.Bool("h", false, "Show this help")
	flag.BoolVar(&sessionizeFlag, "s", false, "Open the worktree in a new tmux session")
	flag.StringVar(&rootFolderFlag, "r", user_home+"/repos", "Root folder to list worktree from")
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	var repo_names = repo_names(rootFolderFlag)
	if repo_names == nil {
		fmt.Println("No repos found.")
		os.Exit(0)
	}

	var names = worktree_names(rootFolderFlag, &repo_names)
	var worktree_name_list = as_worktree_list(&names)
	var selected_wt = ff_select(&worktree_name_list)

	var full_path = rootFolderFlag + "/" + selected_wt
	if !sessionizeFlag {
		fmt.Printf(full_path)
		os.Exit(0)
	} else {
		cmd := exec.Command("tmux", "neww", "-n", selected_wt, "-c", full_path)
		cmd.Run()
	}
}

func repo_names(rootFolder string) []string {
	if _, err := os.Stat(rootFolder); !os.IsNotExist(err) {
		repo_dirs, _ := ioutil.ReadDir(rootFolder)
		var repos []string
		for _, repo := range repo_dirs {
			repos = append(repos, repo.Name())
		}
		return repos
	}
	return nil
}

func worktree_names(rootFolder string, repo_names *[]string) map[string][]string {
	var worktree_names = make(map[string][]string)
	for _, repo := range *repo_names {

		var full_path string = rootFolder + "/" + repo + "/worktrees"
		if _, err := os.Stat(full_path); !os.IsNotExist(err) {
			worktree_names[repo] = make([]string, 0)
			worktrees, _ := ioutil.ReadDir(full_path)
			for _, wt := range worktrees {
				worktree_names[repo] = append(worktree_names[repo], wt.Name())
			}
		}
	}
	return worktree_names
}

func ff_select(paths *[]string) string {
	idx, _ := fzf.Find(*paths, func(i int) string {
		return (*paths)[i]
	})
	return (*paths)[idx]
}

func as_worktree_list(wt_map *map[string][]string) []string {
	wt_list := make([]string, 0)
	for repo, worktrees := range *wt_map {
		for _, worktree := range worktrees {
			wt_list = append(wt_list, repo+"/"+worktree)
		}
	}
	return wt_list
}
