# How to test DRLM

This is a simple "how-to" of testing DRLM

## 0- Machine setup

```
Machine 1 -> (Linux) [Any Distro] -> drlmctl
Machine 2 -> (Linux) [Any Distro] -> DRLM Core
Machine 3 -> (Linux) [Any Distro] -> DRLM Agent 1
Machine 4 (OPTIONAL) -> (Linux) [Any Distro] -> DRLM Agent 2
```

## 1- DRLM Core Setup

- Download the `drlmctl` project on the `Machine 1`
- Modify the configuration file (`drlmctl.toml`):
    ```toml
    [core]
      ...
      tls = "false"
    ...
    ```
- Modify the `software/software.go` file. It needs to have commented out the `git checkout` part:

    ```go
    import (
        ...
        "gopkg.in/src-d/go-git.v4/plumbing"
    )

    ...

	_, err = git.PlainClone(d, false, &git.CloneOptions{
		URL:      s.GitService.CloneURL(s.User, s.Repo),
		Progress: stdOS.Stdout,
	})
	if err != nil {
		return "", fmt.Errorf("error clonning the repository: %v", err)
	}

	// if v == "" {
	// 	v, err = s.GitService.GetRelease(s.User, s.Repo, "latest")
	// 	if err != nil {
	// 		return "", fmt.Errorf("error getting the latest version: %v", err)
	// 	}
	// }

	// w, err := r.Worktree()
	// if err != nil {
	// 	return "", fmt.Errorf("error getting the repository worktree: %v", err)
	// }

	// if err = w.Checkout(&git.CheckoutOptions{
	// 	Branch: plumbing.NewTagReferenceName(v),
	// }); err != nil {
	// 	return "", fmt.Errorf("error checking out to the version %s: %v", v, err)
	// }
    ```

- Add the DRLM Core server: `drlmctl core add --host <Machine 2 IP> --user <SSH User> --password <SSH Password>`
- Install the DRLM Core binary: `drlmctl core install`
