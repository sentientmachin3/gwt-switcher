# gwt-switcher
Git worktree fuzzy switcher written in Golang to learn it.

This tool works for my needs to open a new worktree in a new tmux window or in the current one. 
Yes, I have all the repos in one folder.

## Usage
Flags:
- `-r` changes the repos root folder (default to `$HOME/repos`)
- `-s` if set opens the worktree in a new window (will be changed to `-w` probably)

If no args are set, prints the selected worktree path to the stdout so you can pipe it into `cd` and open it in the current window.

`gwt-switcher -r <repos-folder> [-s]`
