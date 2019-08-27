####Exercise 4.11:
Build a tool that lets users cre ate, read, update, and delete GitHub issues from
the command line, invoking their preferred text editor when substantial text input is required.

GitHub Auth credentials:
```bash
export GIT_HUB_USERNAME=your_github_username
export GIT_HUB_PASSWORD=your_github_password
```

GitHub issues cli:
```bash
cd issues_crud
go build .
./issues_crud list
./issues_crud create {repo} {title} (then body in nano editor)
./issues_crud edit {repo} {title} (then body in nano editor)
./issues_crud lock {repo} {issue_number} {reason}
./issues_crud unlock {repo} {issue_number} {reason}
```





