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
- Modify the `software/software.go` file. It needs to have commented out the `git` dependencies and the `Compile` function as follows:

    ```go
    import (
        ...
        // "gopkg.in/src-d/go-git.v4"
        // "gopkg.in/src-d/go-git.v4/plumbing"
    )

    ...

	// Compile compiles a software piece
	func (s *software) Compile(o os.OS, arch os.Arch, v string) (string, error) {
        d := filepath.Join("../", s.Repo)
        // d, err := afero.TempDir(fs.FS, "", "drlmctl-compile-")
        // if err != nil {
        //      return "", fmt.Errorf("error creating a temporary directory: %v", err)
        // }

        // r, err := git.PlainClone(d, false, &git.CloneOptions{
        //      URL:      s.GitService.CloneURL(s.User, s.Repo),
        //      Progress: stdOS.Stdout,
        // })
        // if err != nil {
        //      return "", fmt.Errorf("error clonning the repository: %v", err)
        // }

        // if v == "" {
        //      v, err = s.GitService.GetRelease(s.User, s.Repo, "latest")
        //      if err != nil {
        //              return "", fmt.Errorf("error getting the latest version: %v", err)
        //      }
        // }

        // w, err := r.Worktree()
        // if err != nil {
        //      return "", fmt.Errorf("error getting the repository worktree: %v", err)
        // }

        // if err = w.Checkout(&git.CheckoutOptions{
        //      Branch: plumbing.NewTagReferenceName(v),
        // }); err != nil {
        //      return "", fmt.Errorf("error checking out to the version %s: %v", v, err)
		// }

		...
    ```

- Add the DRLM Core server: `drlmctl core add --host <Machine 2 IP> --user <SSH User> --password <SSH Password>`
- Install the DRLM Core binary: `drlmctl core install`
- Create the DRLM Database: `CREATE DATABASE drlm3 DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;`
- Create the DRLM Core configuration (in the DRLM Core server): `/home/drlm/core.toml`:

	```toml
	[security]
	tokens_secret = "secretsecretsecretsecretsecretsecret"

	[grpc]
	tls = "false"

	[db]
	host = "localhost"

	[log]
	file = "drlm-core.log"
	```
- Start the DRLM Core server: `su - drlm`, `drlm-core`

# 2- DRLM Agent Setup
- Add the DRLM Agent: `drlmctl agent add --host <Machine 3 IP> --user <SSH User> --password <SSH Password>`
- Install the DRLM Agent `drlmctl agent install --host <Machine 3 IP>`
- Start the DRLM Agent: `su - drlm`, `drlm-agent`
